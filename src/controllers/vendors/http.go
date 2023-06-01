package vendors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/app/helper"
	"github.com/superosystem/bantumanten-backend/src/businesses/vendors"
	"github.com/superosystem/bantumanten-backend/src/controllers/vendors/request"
	"github.com/superosystem/bantumanten-backend/src/controllers/vendors/response"
	"net/http"
)

type Controller struct {
	vendorUseCase vendors.UseCase
	jwtConfig     *config.JWTConfig
}

func NewVendorController(vc vendors.UseCase, jwt *config.JWTConfig) *Controller {
	return &Controller{
		vendorUseCase: vc,
		jwtConfig:     jwt,
	}
}

func (c *Controller) Create(ctx echo.Context) error {
	input := request.Vendor{}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	if err := input.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(err.Error()))
	}

	var currVendorType bool
	switch input.VendorType {
	case constant.VENUE:
		currVendorType = true
	case constant.DECORATION:
		currVendorType = true
	case constant.CATERING:
		currVendorType = true
	case constant.DOCUMENTARY:
		currVendorType = true
	default:
		currVendorType = false
	}

	if currVendorType == false {
		return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrRecordNotFound.Error()))
	}

	err := c.vendorUseCase.Create(input.ToDomain())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			helper.MessageErrorResponse(constant.ErrInternalServerError.Error()))
	}

	return ctx.JSON(http.StatusCreated,
		helper.MessageSuccessResponse("vendor has been created"))
}

func (c *Controller) Update(ctx echo.Context) error {
	var ID string = ctx.Param("id")
	if err := ctx.Bind(&ID); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	input := request.Vendor{}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	rec, err := c.vendorUseCase.Update(ID, *input.ToDomain())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*rec))
}

func (c *Controller) Delete(ctx echo.Context) error {
	var ID string = ctx.Param("id")
	if err := ctx.Bind(&ID); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	deleted := c.vendorUseCase.Delete(ID)
	if deleted == false {
		return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrRecordNotFound.Error()))
	}

	return ctx.JSON(http.StatusOK, helper.MessageErrorResponse("Success delete vendor"))
}

func (c *Controller) FindAll(ctx echo.Context) error {
	var data []response.Vendor

	vendorsData := c.vendorUseCase.GetAll()

	for _, vendor := range *vendorsData {
		data = append(data, response.FromDomain(vendor))
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *Controller) FindByID(ctx echo.Context) error {
	var ID string = ctx.Param("id")
	if err := ctx.Bind(&ID); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	rec, err := c.vendorUseCase.GetByID(ID)
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrRecordNotFound.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
		}
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*rec))
}

func (c *Controller) FindByName(ctx echo.Context) error {
	var name = ctx.Param("name")
	if err := ctx.Bind(&name); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	rec, err := c.vendorUseCase.GetByName(name)
	if err != nil {
		if errors.Is(err, constant.ErrRecordNotFound) {
			return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrRecordNotFound.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
		}
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*rec))
}

func (c *Controller) FindByType(ctx echo.Context) error {
	var vendorType string = ctx.Param("type")
	var data []response.Vendor

	var currVendorType bool
	switch vendorType {
	case constant.VENUE:
		currVendorType = true
	case constant.DECORATION:
		currVendorType = true
	case constant.CATERING:
		currVendorType = true
	case constant.DOCUMENTARY:
		currVendorType = true
	default:
		currVendorType = false
	}

	if currVendorType == false {
		return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrRecordNotFound.Error()))
	}

	vendorsData := c.vendorUseCase.GetByType(vendorType)

	for _, vendor := range *vendorsData {
		data = append(data, response.FromDomain(vendor))
	}

	return ctx.JSON(http.StatusOK, data)
}
