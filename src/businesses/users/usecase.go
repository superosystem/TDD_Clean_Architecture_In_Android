package users

import (
	"github.com/superosystem/bantumanten-backend/src/app/constant"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository Repository
}

func NewUserUseCase(ur Repository) UseCase {
	return &UserUseCase{
		userRepository: ur,
	}
}

func (u UserUseCase) SignUp(domain *Domain) error {
	var err error

	// CHECK EMAIL IS_EXIST
	user, _ := u.userRepository.GetByEmail(domain.Email)
	if user != nil {
		return err
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

func (u UserUseCase) SignIn(domain *SignInDomain) (*Domain, error) {
	//TODO implement me
	panic("implement me")
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
