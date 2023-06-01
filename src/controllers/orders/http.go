package orders

import (
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/app/helper"
	"github.com/superosystem/bantumanten-backend/src/businesses/orders"
	"github.com/superosystem/bantumanten-backend/src/controllers/orders/request"
	"net/http"
	"strconv"
)

type Controller struct {
	orderUseCase orders.UseCase
	jwtConfig    *config.JWTConfig
}

func NewOrderController(ou orders.UseCase, jwt *config.JWTConfig) *Controller {
	return &Controller{
		orderUseCase: ou,
		jwtConfig:    jwt,
	}
}

func (c *Controller) Create(ctx echo.Context) error {
	token, _ := c.jwtConfig.ExtractToken(ctx)
	id, _ := strconv.Atoi(token.UserId)

	input := request.Order{}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}
	if err := input.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(err.Error()))
	}

	input.UserId = uint(id)
	err := c.orderUseCase.Create(input.ToDomain())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error()))
	}

	return ctx.JSON(http.StatusCreated,
		helper.MessageSuccessResponse("order has been created"))
}
