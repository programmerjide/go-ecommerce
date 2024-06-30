package service

import (
	"fmt"
	"github.com/programmerolajide/go-ecommerce/internal/domain"
	"github.com/programmerolajide/go-ecommerce/internal/dto"
	"github.com/programmerolajide/go-ecommerce/internal/repository"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func (us UserService) Signup(createUserRequestDTO dto.UserSignupRequestDTO) (string, error) {
	log.Printf("User signup with email: %v", createUserRequestDTO.Email)

	user, err := us.Repo.CreateUser(domain.User{
		Email:    createUserRequestDTO.Email,
		Password: createUserRequestDTO.Password,
		Phone:    createUserRequestDTO.Phone,
	})

	// generate token

	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)

	return userInfo, err
}

func (us UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform db  operation
	return nil, nil
}

func (us UserService) Login(input any) (string, error) {

	return "", nil
}

func (us UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (us UserService) VerifyCode(id uint, code int) error {

	return nil
}

func (us UserService) CreateProfile(id uint, input any) error {

	return nil
}

func (us UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (us UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (us UserService) BecomeSeller(id uint, input any) (string, error) {

	return "", nil
}

func (us UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (us UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (us UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (us UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (us UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {

	return nil, nil
}
