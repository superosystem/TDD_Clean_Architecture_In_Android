package transactions

import (
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/businesses/transactions"
)

type Controller struct {
	transactionUseCase transactions.UseCase
}

func NewTransactionController(tu transactions.UseCase) *Controller {
	return &Controller{
		transactionUseCase: tu,
	}
}

func (c *Controller) FindAll(ctx echo.Context) error {
	panic("please implement")
}
