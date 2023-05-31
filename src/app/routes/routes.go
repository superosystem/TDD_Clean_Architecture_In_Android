package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/middleware"
	"gorm.io/gorm"

	_driver "github.com/superosystem/bantumanten-backend/src/drivers"

	_userUseCase "github.com/superosystem/bantumanten-backend/src/businesses/users"
	_userController "github.com/superosystem/bantumanten-backend/src/controllers/users"
)

type Config struct {
	// ECHO TOP LEVEL INSTANCE
	Echo *echo.Echo
	// MYSQL CONNECTION
	MySQLCONN *gorm.DB
	// JWT
	JwtConfig *config.JWTConfig
}

func (cfg *Config) Start() {
	authMiddleware := middleware.NewAuthMiddleware(cfg.JwtConfig)
	registerUserRoute(cfg, authMiddleware)
}

func registerUserRoute(cfg *Config, authMiddleware *middleware.AuthMiddleware) {
	// USER DI
	userRepository := _driver.NewUserRepository(cfg.MySQLCONN)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, cfg.JwtConfig)
	userController := _userController.NewUserController(userUseCase, cfg.JwtConfig)

	// ROUTES
	v1 := cfg.Echo.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.POST("/register", userController.SignUp)
	auth.POST("/login", userController.SignIn)

	user := v1.Group("/users", authMiddleware.IsAuthenticated())
	user.GET("", userController.FindAll, authMiddleware.IsAdminRole)
	user.GET("/:id", userController.FindByID, authMiddleware.IsAdminRole)
	user.GET("/:email", userController.FindByEmail, authMiddleware.IsAdminRole)
	user.PUT("", userController.Update)

}
