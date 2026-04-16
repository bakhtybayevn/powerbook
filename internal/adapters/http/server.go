package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/handlers"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	jwtToken "github.com/bakhtybayevn/powerbook/internal/adapters/http/token"
	postgres "github.com/bakhtybayevn/powerbook/internal/adapters/postgres"
	"github.com/bakhtybayevn/powerbook/internal/adapters/redis"
	appCompetition "github.com/bakhtybayevn/powerbook/internal/application/competition"
	appReading "github.com/bakhtybayevn/powerbook/internal/application/reading"
	appUser "github.com/bakhtybayevn/powerbook/internal/application/user"
	"github.com/bakhtybayevn/powerbook/internal/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
}

func NewServer(router *gin.Engine, cfg *config.Config) *Server {
	return &Server{
		router: router,
		cfg:    cfg,
	}
}

func (s *Server) RegisterRoutes() {
	// === GLOBAL MIDDLEWARES ===
	s.router.Use(middleware.ErrorMiddleware())
	s.router.Use(middleware.CORSMiddleware())

	// === SWAGGER ===
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// === HEALTH CHECK ===
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	db, err := sql.Open("postgres", s.cfg.PostgresDSN())
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	redisAddr := fmt.Sprintf("%s:%d", s.cfg.Redis.Host, s.cfg.Redis.Port)

	userRepo := postgres.NewPostgresUserRepo(db)
	tokenService := jwtToken.NewJWTService(s.cfg.JWT.Secret)
	readingRepo := postgres.NewPostgresReadingRepo(db)
	competitionRepo := postgres.NewPostgresCompetitionRepo(db)
	redisLB := redis.NewRedisLeaderboard(redisAddr, s.cfg.Redis.Password, s.cfg.Redis.UseTLS)
	lbHealth := middleware.RedisHealth(redisLB)

	// === HANDLERS ===
	leaderboardHandler := handlers.NewLeaderboardHandler(redisLB, userRepo)

	// === USE CASES ===
	registerUserHandler := appUser.NewRegisterUserHandler(userRepo)
	loginUserHandler := appUser.NewLoginUserHandler(userRepo, tokenService)
	logReadingHandler := appReading.NewLogReadingHandler(userRepo, readingRepo, competitionRepo, redisLB)
	createCompetitionHandler := appCompetition.NewCreateCompetitionHandler(competitionRepo)
	joinCompetitionHandler := appCompetition.NewJoinCompetitionHandler(competitionRepo)
	closeCompetitionHandler := appCompetition.NewCloseCompetitionHandler(competitionRepo, userRepo)
	listAllCompetitionsHandler := appCompetition.NewListAllCompetitionsHandler(competitionRepo, userRepo)
	listMyCompetitionsHandler := appCompetition.NewListMyCompetitionsHandler(competitionRepo, userRepo)

	// === API VERSIONING (/api/v1) ===
	v1 := s.router.Group("/api/v1")

	// ---- Public endpoints ----
	v1.POST("/users/register", handlers.RegisterUser(registerUserHandler))
	v1.POST("/users/login", handlers.LoginUser(loginUserHandler))
	v1.GET("/users/:id", handlers.GetUserProfile(userRepo))
	v1.GET("/competitions", handlers.ListAllCompetitions(listAllCompetitionsHandler))
	v1.GET("/competitions/:id", handlers.GetCompetition(competitionRepo, userRepo))
	v1.GET("/competitions/:id/leaderboard", lbHealth, leaderboardHandler.GetLeaderboard)
	v1.GET("/competitions/:id/rank/:userID", lbHealth, leaderboardHandler.GetRank)
	v1.GET("/competitions/:id/gifts", handlers.GetGiftExchanges(competitionRepo, userRepo))

	// ---- Protected endpoints ----
	auth := v1.Group("/")
	auth.Use(middleware.AuthMiddleware(tokenService))
	auth.GET("/users/me", handlers.GetMe(userRepo))
	auth.PUT("/users/me/profile", handlers.UpdateProfile(userRepo))
	auth.POST("/reading/log", handlers.LogReading(logReadingHandler))
	auth.GET("/reading/history", handlers.ReadingHistory(readingRepo))
	auth.POST("/competitions/create", handlers.CreateCompetition(createCompetitionHandler))
	auth.POST("/competitions/:id/join", handlers.JoinCompetition(joinCompetitionHandler))
	auth.POST("/competitions/:id/close", handlers.CloseCompetition(closeCompetitionHandler))
	auth.GET("/competitions/:id/rank/me", lbHealth, leaderboardHandler.GetRankMe)
	auth.GET("/competitions/my", handlers.ListMyCompetitions(listMyCompetitionsHandler))
	auth.POST("/gifts/:giftId/confirm", handlers.ConfirmGift(competitionRepo))

	// === AUTO-CLOSE SCHEDULER ===
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			comps, err := competitionRepo.GetAll()
			if err != nil {
				log.Printf("[AutoClose] failed to get competitions: %v", err)
				continue
			}
			now := time.Now().UTC()
			for _, c := range comps {
				if c.Status == "open" && now.After(c.EndDate) {
					log.Printf("[AutoClose] closing competition %s (%s)", c.ID, c.Name)
					closeCompetitionHandler.Handle(appCompetition.CloseCompetitionCommand{
						CompetitionID: c.ID,
					})
				}
			}
		}
	}()
}
