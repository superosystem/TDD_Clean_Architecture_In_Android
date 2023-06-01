package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/middleware"

	_order "github.com/superosystem/bantumanten-backend/src/businesses/orders"
	_transaction "github.com/superosystem/bantumanten-backend/src/businesses/transactions"
	_user "github.com/superosystem/bantumanten-backend/src/businesses/users"
	_vendor "github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	_orderController "github.com/superosystem/bantumanten-backend/src/controllers/orders"
	_transactionController "github.com/superosystem/bantumanten-backend/src/controllers/transactions"
	_userController "github.com/superosystem/bantumanten-backend/src/controllers/users"
	_vendorController "github.com/superosystem/bantumanten-backend/src/controllers/vendors"

	"gorm.io/gorm"

	_driver "github.com/superosystem/bantumanten-backend/src/drivers"
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

	// DEPENDENCY INJECTION
	userRepository := _driver.NewUserRepository(cfg.MySQLCONN)
	vendorRepository := _driver.NewVendorRepository(cfg.MySQLCONN)
	orderRepository := _driver.NewOrderRepository(cfg.MySQLCONN)
	transactionRepository := _driver.NewTransactionRepository(cfg.MySQLCONN)

	userUseCase := _user.NewUserUseCase(userRepository, cfg.JwtConfig)
	vendorUseCase := _vendor.NewVendorUseCase(vendorRepository)
	orderUseCase := _order.NewOrderUseCase(orderRepository, vendorRepository, userRepository)
	transactionUseCase := _transaction.NewTransactionUseCase(transactionRepository)

	userController := _userController.NewUserController(userUseCase, cfg.JwtConfig)
	vendorController := _vendorController.NewVendorController(vendorUseCase, cfg.JwtConfig)
	orderController := _orderController.NewOrderController(orderUseCase, userUseCase, vendorUseCase, transactionUseCase, cfg.JwtConfig)
	transactionController := _transactionController.NewTransactionController(transactionUseCase)

	// ROUTES V1
	v1 := cfg.Echo.Group("/http/v1")
	v1.GET("", userController.HelloMessage)
	// AUTH ROUTES
	auth := v1.Group("")
	auth.POST("/register", userController.SignUp)
	auth.POST("/login", userController.SignIn)
	// USER ROUTES
	user := v1.Group("/users", authMiddleware.IsAuthenticated())
	user.GET("", userController.FindAll, authMiddleware.IsAdminRole)
	user.GET("/:id", userController.FindByID, authMiddleware.IsAdminRole)
	user.GET("/:email", userController.FindByEmail, authMiddleware.IsAdminRole)
	user.GET("/profile", userController.FindProfile)
	user.PUT("/profile", userController.UpdateProfile)
	// VENDOR ROUTES
	vendor := v1.Group("/vendors", authMiddleware.IsAuthenticated())
	vendor.POST("", vendorController.Create, authMiddleware.IsAdminRole)
	vendor.PUT("", vendorController.Update, authMiddleware.IsAdminRole)
	vendor.GET("", vendorController.FindAll, authMiddleware.IsAdminRole)
	vendor.GET("/:type", vendorController.FindByType)
	vendor.GET("/:id", vendorController.FindByID)
	vendor.GET("/:name", vendorController.FindByName)
	vendor.DELETE("/:id", vendorController.Delete, authMiddleware.IsAdminRole)
	// ORDER ROUTES
	order := v1.Group("/orders", authMiddleware.IsAuthenticated())
	order.POST("", orderController.Create, authMiddleware.IsUserRole)
	// TRANSACTION ROUTES
	transaction := v1.Group("/transactions", authMiddleware.IsAuthenticated())
	transaction.GET("", transactionController.FindAll, authMiddleware.IsAdminRole)
	// PAYMENT ROUTES
	// SCHEDULE ROUTES
}
