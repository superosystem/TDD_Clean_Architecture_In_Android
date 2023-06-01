package drivers

import (
	"gorm.io/gorm"

	_userDomain "github.com/superosystem/bantumanten-backend/src/businesses/users"
	_vendorDomain "github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	_userDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
	_vendorDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/vendors"
)

func NewUserRepository(conn *gorm.DB) _userDomain.Repository {
	return _userDB.NewMySQLRepository(conn)
}

func NewVendorRepository(conn *gorm.DB) _vendorDomain.Repository {
	return _vendorDB.NewMySQLRepository(conn)
}
