package drivers

import (
	"gorm.io/gorm"

	_userDomain "github.com/superosystem/bantumanten-backend/src/businesses/users"
	_userDB "github.com/superosystem/bantumanten-backend/src/drivers/mysql/users"
)

func NewUserRepository(conn *gorm.DB) _userDomain.Repository {
	return _userDB.NewMySQLRepository(conn)
}
