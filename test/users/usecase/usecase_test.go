package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"github.com/superosystem/bantumanten-backend/src/businesses/users"
	"github.com/superosystem/bantumanten-backend/src/businesses/users/mocks"
	"testing"
	"time"
)

var (
	userDomain  users.Domain
	userRepo    mocks.Repository
	userUseCase users.UseCase
)

func TestMain(m *testing.M) {
	userRepo = mocks.Repository{Mock: mock.Mock{}}
	userUseCase = users.NewUserUseCase(&userRepo)

	userDomain = users.Domain{
		ID:        1,
		FullName:  "John Doe",
		Email:     "johndoe@exampl.com",
		Password:  "hashedpassword",
		Photo:     "",
		Roles:     "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.Run()
}

func TestSignUp(t *testing.T) {
	t.Run("SIGNUP | SUCCESS", func(t *testing.T) {
		userRepo.Mock.On("GetByEmail", userDomain.Email).Return(nil, nil)
		userRepo.Mock.On("Create", mock.Anything).Return(nil)

		err := userUseCase.SignUp(&userDomain)

		assert.NoError(t, err)
	})

	t.Run("SIGN UP | FAIL EMAIL EXIST", func(t *testing.T) {
		userRepo.Mock.On("GetByEmail", userDomain.Email).Return(userDomain)

		err := userUseCase.SignUp(&userDomain)

		assert.Error(t, constant.ErrEmailAlreadyExist, err)
	})
}
