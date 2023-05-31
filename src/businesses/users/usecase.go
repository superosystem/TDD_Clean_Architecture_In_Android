package users

import (
	"github.com/superosystem/bantumanten-backend/src/app/config"
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type UserUseCase struct {
	userRepository Repository
	jwtConfig      *config.JWTConfig
}

func NewUserUseCase(ur Repository, jwt *config.JWTConfig) UseCase {
	return &UserUseCase{
		userRepository: ur,
		jwtConfig:      jwt,
	}
}

func (u UserUseCase) SignUp(domain *Domain) error {
	var err error

	// CHECK EMAIL IS_EXIST
	user, _ := u.userRepository.GetByEmail(domain.Email)
	if user != nil {
		return constant.ErrEmailAlreadyExist
	}
	// HASHED PASSWORD
	if len(domain.Password) < 6 {
		return constant.ErrPasswordLengthInvalid
	}
	hashedPasswd, _ := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.DefaultCost)

	// SET DEFAULT USER
	domain.Password = string(hashedPasswd)
	domain.Photo = ""
	domain.Roles = constant.ROLE_USER

	err = u.userRepository.Create(domain)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUseCase) SignIn(domain *SignInDomain) (interface{}, error) {
	var user *Domain

	// Check Email
	user, err := u.userRepository.GetByEmail(domain.Email)
	if err != nil {
		return nil, constant.ErrUserNotFound
	}
	// Check Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(domain.Password))
	if err != nil {
		return nil, constant.ErrWrongPassword
	}
	// Generate Token
	var token string
	exp := time.Now().Add(6 * time.Hour)
	token, err = u.jwtConfig.GenerateToken(strconv.Itoa(int(user.ID)), user.Roles, exp)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"token": token,
	}

	return data, nil
}

func (u UserUseCase) Update(ID string, domain *Domain) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) Delete(ID string) bool {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetAll() []Domain {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetByID(ID string) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) GetByEmail(email string) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}
