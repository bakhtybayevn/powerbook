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

// RegisterRoutes sets up routes including swagger and a test endpoint
func (s *Server) RegisterRoutes() {
	// Swagger UI
	// Swagger JSON будет доступен по адресу: /swagger/doc.json
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// in-memory repo
	userRepo := memrepo.NewInMemoryUserRepo()
	tokenService := jwtToken.NewJWTService("supersecret")

	loginUserHandler := appUser.NewLoginUserHandler(userRepo, tokenService)
	registerUserHandler := appUser.NewRegisterUserHandler(userRepo)

	authMiddleware := middleware.AuthMiddleware(tokenService)
	auth := s.router.Group("/")
	auth.Use(authMiddleware)
	auth.GET("/users/me", userHandlers.GetMe(userRepo))

	// User endpoints
	s.router.POST("/users/register", userHandlers.RegisterUser(registerUserHandler))
	s.router.POST("/users/login", userHandlers.LoginUser(loginUserHandler))
}
