package routes

import (
	"github.com/labstack/echo/v4"
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
}

func (cfg *Config) Start() {
}

func userRoutes(cfg *Config) {
	// USER DI
	userRepository := _driver.NewUserRepository(cfg.MySQLCONN)
	userUseCase := _userUseCase.NewUserUseCase(userRepository)
	userController := _userController.NewUserController(userUseCase)

	// ROUTES
	v1 := cfg.Echo.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.POST("/register", userController.SignUp)
	auth.POST("/login", userController.SignUp)

}
