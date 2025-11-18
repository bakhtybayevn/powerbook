package http

import (
	"net/http"

	userHandlers "github.com/bakhtybayevn/powerbook/internal/adapters/http/handlers"
	"github.com/bakhtybayevn/powerbook/internal/adapters/http/middleware"
	jwtToken "github.com/bakhtybayevn/powerbook/internal/adapters/http/token"
	memrepo "github.com/bakhtybayevn/powerbook/internal/adapters/postgres"
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

	// === USE CASES ===
	registerUserHandler := appUser.NewRegisterUserHandler(userRepo)
	loginUserHandler := appUser.NewLoginUserHandler(userRepo, tokenService)

	// === API VERSIONING (/api/v1) ===
	v1 := s.router.Group("/api/v1")

	// ---- Public endpoints ----
	v1.POST("/users/register", userHandlers.RegisterUser(registerUserHandler))
	v1.POST("/users/login", userHandlers.LoginUser(loginUserHandler))

	// ---- Protected endpoints ----
	auth := v1.Group("/")
	auth.Use(middleware.AuthMiddleware(tokenService))
	auth.GET("/users/me", userHandlers.GetMe(userRepo))
}
