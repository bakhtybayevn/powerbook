package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
	leaderboardHandler := handlers.NewLeaderboardHandler(redisLB)

	// === USE CASES ===
	registerUserHandler := appUser.NewRegisterUserHandler(userRepo)
	loginUserHandler := appUser.NewLoginUserHandler(userRepo, tokenService)
	logReadingHandler := appReading.NewLogReadingHandler(userRepo, readingRepo, competitionRepo, redisLB)
	createCompetitionHandler := appCompetition.NewCreateCompetitionHandler(competitionRepo)
	joinCompetitionHandler := appCompetition.NewJoinCompetitionHandler(competitionRepo)
	closeCompetitionHandler := appCompetition.NewCloseCompetitionHandler(competitionRepo)
	listAllCompetitionsHandler := appCompetition.NewListAllCompetitionsHandler(competitionRepo, userRepo)
	listMyCompetitionsHandler := appCompetition.NewListMyCompetitionsHandler(competitionRepo, userRepo)

	// === API VERSIONING (/api/v1) ===
	v1 := s.router.Group("/api/v1")

	// ---- Public endpoints ----
	v1.POST("/users/register", handlers.RegisterUser(registerUserHandler))
	v1.POST("/users/login", handlers.LoginUser(loginUserHandler))

	// ---- Protected endpoints ----
	auth := v1.Group("/")
	auth.Use(middleware.AuthMiddleware(tokenService))
	auth.GET("/users/me", handlers.GetMe(userRepo))
	auth.POST("/reading/log", handlers.LogReading(logReadingHandler))
	auth.POST("/competitions/create", handlers.CreateCompetition(createCompetitionHandler))
	auth.POST("/competitions/:id/join", handlers.JoinCompetition(joinCompetitionHandler))
	auth.POST("/competitions/:id/close", handlers.CloseCompetition(closeCompetitionHandler))
	auth.GET("/competitions/:id/leaderboard", lbHealth, leaderboardHandler.GetLeaderboard)
	auth.GET("/competitions/:id/rank/:userID", lbHealth, leaderboardHandler.GetRank)
	auth.GET("/competitions/:id/rank/me", lbHealth, leaderboardHandler.GetRankMe)
	auth.GET("/competitions", handlers.ListAllCompetitions(listAllCompetitionsHandler))
	auth.GET("/competitions/my", handlers.ListMyCompetitions(listMyCompetitionsHandler))
}
