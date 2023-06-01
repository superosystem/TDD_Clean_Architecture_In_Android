package orders

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/app/helper"
	"github.com/superosystem/bantumanten-backend/src/businesses/orders"
	"github.com/superosystem/bantumanten-backend/src/businesses/transactions"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	"github.com/superosystem/bantumanten-backend/src/controllers/orders/request"
	"net/http"
)

type Controller struct {
	orderUseCase       orders.UseCase
	userUseCase        users.UseCase
	vendorUseCase      vendors.UseCase
	transactionUseCase transactions.UseCase
	jwtConfig          *config.JWTConfig
}

func NewOrderController(
	ou orders.UseCase,
	uu users.UseCase,
	vu vendors.UseCase,
	tr transactions.UseCase,
	jwt *config.JWTConfig,
) *Controller {
	return &Controller{
		orderUseCase:       ou,
		userUseCase:        uu,
		vendorUseCase:      vu,
		transactionUseCase: tr,
		jwtConfig:          jwt,
	}
}

func (c *Controller) Create(ctx echo.Context) error {
	input := request.Order{}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}
	if err := input.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(err.Error()))
	}

	// Find Credential User
	token, _ := c.jwtConfig.ExtractToken(ctx)
	user, err := c.userUseCase.GetByID(token.UserId)
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrRecordNotFound.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error()))
	}
	input.UserID = user.ID

	// Set up Vendor Price
	err2, done := setUpVendorPrice(ctx, err, c, input)
	if done {
		return err2
	}

	// Create Order and Transaction
	order, err := c.orderUseCase.Create(input.ToDomain())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error()))
	}

	var transaction transactions.Domain
	transaction.ReferenceNumber = order.ReferenceNumber
	transaction.Status = constant.TRANSACTION_PENDING
	transaction.Cred = 0
	transaction.Debt = order.TotalAmount
	transaction.Total = order.TotalAmount
	transaction.UserID = order.UserID
	transaction.OrderID = order.ID
	_, err = c.transactionUseCase.Create(&transaction)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error()))
	}
	return ctx.JSON(http.StatusCreated,
		helper.MessageSuccessResponse("order has been created "+order.ReferenceNumber))
}

func setUpVendorPrice(ctx echo.Context, err error, c *Controller, input request.Order) (error, bool) {
	venue, err := c.vendorUseCase.GetByName(input.Venue)
	if err != err {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrRecordNotFound.Error())), true
		}
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error())), true
	}
	input.VenuePrice = venue.Price

	decoration, err := c.vendorUseCase.GetByName(input.Decoration)
	if err != err {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrRecordNotFound.Error())), true
		}
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error())), true
	}
	input.DecorationPrice = decoration.Price

	catering, err := c.vendorUseCase.GetByName(input.Catering)
	if err != err {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrRecordNotFound.Error())), true
		}
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error())), true
	}
	input.CateringPrice = catering.Price

	mua, err := c.vendorUseCase.GetByName(input.Mua)
	if err != err {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrRecordNotFound.Error())), true
		}
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error())), true
	}
	input.MuaPrice = mua.Price

	documentary, err := c.vendorUseCase.GetByName(input.Documentary)
	if err != err {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrRecordNotFound.Error())), true
		}
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error())), true
	}
	input.DocumentaryPrice = documentary.Price
	return nil, false
}
