package users

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/app/helper"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"github.com/superosystem/bantumanten-backend/src/controllers/users/request"
	"net/http"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(uc users.UseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (c *UserController) HelloMessage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Welcome to Bantu Manten API.")
}

func (c *UserController) SignUp(ctx echo.Context) error {
	input := request.User{}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.BadRequestResponse(constant.ErrInvalidRequest.Error()))
	}

	if input.Password != input.ConfirmationPassword {
		return ctx.JSON(http.StatusBadRequest,
			helper.BadRequestResponse(constant.ErrPasswordNotMatch.Error()))
	}

	if err := input.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.BadRequestResponse(err.Error()))
	}

	err := c.userUseCase.SignUp(input.ToDomainRegister())
	if err != nil {
		if errors.Is(err, constant.ErrPasswordLengthInvalid) {
			return ctx.JSON(http.StatusBadRequest,
				helper.BadRequestResponse(constant.ErrPasswordLengthInvalid.Error()))
		} else if errors.Is(err, constant.ErrEmailAlreadyExist) {
			return ctx.JSON(http.StatusBadRequest,
				helper.BadRequestResponse(constant.ErrEmailAlreadyExist.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError,
				helper.BadRequestResponse(constant.ErrInternalServerError.Error()))
		}
	}

	return ctx.JSON(http.StatusOK,
		helper.SuccessCreatedResponse("user has been created", nil))
}

//
//func (c *UserController) SignIn(ctx echo.Context) error {
//	input := request.SignIn{}
//
//	if err := ctx.Bind(&input); err != nil {
//		return ctx.JSON(http.StatusBadRequest,
//			helper.BadRequestResponse(constant.ErrInvalidRequest.Error()))
//	}
//
//	if err := input.Validate(); err != nil {
//		return ctx.JSON(http.StatusBadRequest,
//			helper.BadRequestResponse(err.Error()))
//	}
//
//	token := c.userUseCase.SignIn(input.ToDomainLogin())
//
//	newUser := c.userUseCase.SignUp(input.ToDomainRegister())
//	if newUser.ID == 0 {
//		return ctx.JSON(http.StatusBadRequest,
//			helper.BadRequestResponse(constant.ErrEmailAlreadyExist.Error()))
//	}
//
//	return ctx.JSON(http.StatusOK,
//		helper.SuccessCreatedResponse("user has been created", response.FromDomain(newUser)))
//}
