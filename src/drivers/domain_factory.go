package drivers

import (
	"gorm.io/gorm"

	_orderDomain "github.com/superosystem/bantumanten-backend/src/businesses/orders"
	_transactionDomain "github.com/superosystem/bantumanten-backend/src/businesses/transactions"
	_userDomain "github.com/superosystem/bantumanten-backend/src/businesses/users"
	_vendorDomain "github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	_orderDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/orders"
	_transactionDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/transactions"
	_userDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
	_vendorDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/vendors"
)

func NewUserRepository(conn *gorm.DB) _userDomain.Repository {
	return _userDB.NewMySQLRepository(conn)
}

func NewVendorRepository(conn *gorm.DB) _vendorDomain.Repository {
	return _vendorDB.NewMySQLRepository(conn)
}

func NewOrderRepository(conn *gorm.DB) _orderDomain.Repository {
	return _orderDB.NewMySQLRepository(conn)
}

func NewTransactionRepository(conn *gorm.DB) _transactionDomain.Repository {
	return _transactionDB.NewMySQLRepository(conn)
}
