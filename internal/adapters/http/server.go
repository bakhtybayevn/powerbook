package http

import (
	"net/http"

	"github.com/bakhtybayevn/powerbook/internal/adapters/http/handlers"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	jwtToken "github.com/bakhtybayevn/powerbook/internal/adapters/http/token"
	memrepo "github.com/bakhtybayevn/powerbook/internal/adapters/postgres"
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

	// === DEPENDENCIES (temporary in-memory) ===
	userRepo := memrepo.NewInMemoryUserRepo()
	tokenService := jwtToken.NewJWTService("supersecret")
	readingRepo := memrepo.NewInMemoryReadingRepo()
	competitionRepo := memrepo.NewInMemoryCompetitionRepo()

	// === USE CASES ===
	registerUserHandler := appUser.NewRegisterUserHandler(userRepo)
	loginUserHandler := appUser.NewLoginUserHandler(userRepo, tokenService)
	logReadingHandler := appReading.NewLogReadingHandler(userRepo, readingRepo, competitionRepo)
	createCompetitionHandler := appCompetition.NewCreateCompetitionHandler(competitionRepo)
	joinCompetitionHandler := appCompetition.NewJoinCompetitionHandler(competitionRepo)
	closeCompetitionHandler := appCompetition.NewCloseCompetitionHandler(competitionRepo)

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
}
