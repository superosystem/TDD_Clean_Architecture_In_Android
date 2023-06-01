package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/middleware"
	"gorm.io/gorm"

	_driver "github.com/superosystem/bantumanten-backend/src/drivers"

	_user "github.com/superosystem/bantumanten-backend/src/businesses/users"
	_userController "github.com/superosystem/bantumanten-backend/src/controllers/users"

	_vendor "github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	_vendorController "github.com/superosystem/bantumanten-backend/src/controllers/vendors"
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
	addUserRoutes(cfg, authMiddleware)
	addVendorRoutes(cfg, authMiddleware)
}

// USER ROUTES
func addUserRoutes(cfg *Config, authMiddleware *middleware.AuthMiddleware) {
	// USER DI
	userRepository := _driver.NewUserRepository(cfg.MySQLCONN)
	userUseCase := _user.NewUserUseCase(userRepository, cfg.JwtConfig)
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

// VENDOR ROUTES
func addVendorRoutes(cfg *Config, authMiddleware *middleware.AuthMiddleware) {
	// VENDOR DI
	vendorRepository := _driver.NewVendorRepository(cfg.MySQLCONN)
	vendorUseCase := _vendor.NewVendorUseCase(vendorRepository, cfg.JwtConfig)
	vendorController := _vendorController.NewVendorController(vendorUseCase, cfg.JwtConfig)

	// ROUTES
	v1 := cfg.Echo.Group("/api/v1")
	vendor := v1.Group("/vendors", authMiddleware.IsAuthenticated())
	vendor.POST("", vendorController.Create, authMiddleware.IsAdminRole)
	vendor.PUT("", vendorController.Update, authMiddleware.IsAdminRole)
	vendor.GET("", vendorController.FindAll, authMiddleware.IsAdminRole)
	vendor.GET("/:id", vendorController.FindByID)
	vendor.GET("/:type", vendorController.FindByType)

}
