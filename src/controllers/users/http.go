package users

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/app/helper"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"github.com/superosystem/bantumanten-backend/src/controllers/users/request"
	"github.com/superosystem/bantumanten-backend/src/controllers/users/response"
	"net/http"
)

type Controller struct {
	userUseCase users.UseCase
	jwtConfig   *config.JWTConfig
}

func NewUserController(uc users.UseCase, jwt *config.JWTConfig) *Controller {
	return &Controller{
		userUseCase: uc,
		jwtConfig:   jwt,
	}
}

func (c *Controller) HelloMessage(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Welcome to Bantu Manten API.")
}

func (c *Controller) SignUp(ctx echo.Context) error {
	input := request.User{}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	if input.Password != input.ConfirmationPassword {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrPasswordNotMatch.Error()))
	}

	if err := input.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(err.Error()))
	}

	err := c.userUseCase.SignUp(input.ToDomainSignUp())
	if err != nil {
		if errors.Is(err, constant.ErrPasswordLengthInvalid) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrPasswordLengthInvalid.Error()))
		} else if errors.Is(err, constant.ErrEmailAlreadyExist) {
			return ctx.JSON(http.StatusBadRequest,
				helper.MessageErrorResponse(constant.ErrEmailAlreadyExist.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError,
				helper.MessageErrorResponse(constant.ErrInternalServerError.Error()))
		}
	}

	return ctx.JSON(http.StatusCreated,
		helper.MessageSuccessResponse("user has been created"))
}

func (c *Controller) SignIn(ctx echo.Context) error {
	input := request.SignIn{}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.BadRequestResponse(constant.ErrInvalidRequest.Error()))
	}

	if err := input.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.BadRequestResponse(err.Error()))
	}

	data, err := c.userUseCase.SignIn(input.ToDomainSignIn())
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

	return ctx.JSON(http.StatusOK, data)
}

func (c *Controller) FindAll(ctx echo.Context) error {
	var data []response.User

	usersData := c.userUseCase.GetAll()

	for _, user := range *usersData {
		data = append(data, response.FromDomain(user))
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c *Controller) FindByID(ctx echo.Context) error {
	var ID string = ctx.Param("id")

	rec, err := c.userUseCase.GetByID(ID)
	if err != nil {
		if errors.Is(err, constant.ErrUserNotFound) {
			return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrUserNotFound.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
		}
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*rec))
}

func (c *Controller) FindByEmail(ctx echo.Context) error {
	var email string = ctx.Param("email")

	rec, err := c.userUseCase.GetByEmail(email)
	if err != nil {
		if errors.Is(err, constant.ErrUserNotFound) {
			return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrUserNotFound.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
		}
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*rec))
}

func (c *Controller) FindProfile(ctx echo.Context) error {
	token, _ := c.jwtConfig.ExtractToken(ctx)

	rec, err := c.userUseCase.GetByID(token.UserId)
	if err != nil {
		if errors.Is(err, constant.ErrUserNotFound) {
			return ctx.JSON(http.StatusNotFound, helper.MessageErrorResponse(constant.ErrUserNotFound.Error()))
		} else {
			return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
		}
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*rec))
}

func (c *Controller) UpdateProfile(ctx echo.Context) error {
	token, _ := c.jwtConfig.ExtractToken(ctx)

	input := request.UserUpdate{}
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			helper.MessageErrorResponse(constant.ErrInvalidRequest.Error()))
	}

	updatedUser, err := c.userUseCase.Update(token.UserId, input.ToDomainUserUpdate())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.MessageErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.FromDomain(*updatedUser))
}
