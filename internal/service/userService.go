package service

import (
	"errors"
	"github.com/programmerolajide/go-ecommerce/internal/domain"
	"github.com/programmerolajide/go-ecommerce/internal/dto"
	"github.com/programmerolajide/go-ecommerce/internal/helper"
	"github.com/programmerolajide/go-ecommerce/internal/repository"
	"log"
	"time"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (us UserService) Signup(createUserRequestDTO dto.UserSignupRequestDto) (dto.SignupResponseData, error) {
	log.Printf("User signup with email: %v", createUserRequestDTO.Email)

	hashPassword, err := us.Auth.CreateHashedPassword(createUserRequestDTO.Password)
	if err != nil {
		return dto.SignupResponseData{}, err
	}

	user, err := us.Repo.CreateUser(domain.User{
		Email:    createUserRequestDTO.Email,
		Password: hashPassword,
		Phone:    createUserRequestDTO.Phone,
	})
	if err != nil {
		return dto.SignupResponseData{}, err
	}

	signResponseData := dto.SignupResponseData{
		Id:    user.ID,
		Email: user.Email,
	}

	return signResponseData, nil
}

func (us UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform db  operation
	user, err := us.Repo.FindUser(email)
	return &user, err
}

func (us UserService) Login(email string, password string) (string, error) {

	user, err := us.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with email")
	}

	err = us.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	// generate token
	return us.Auth.GenerateAccessToken(user.ID, user.Email, user.UserType)
}

func (us UserService) isUserVerified(id uint) bool {
	currentUser, err := us.Repo.FindUserById(id)
	return err == nil && currentUser.Verified
}

func (us UserService) GetVerificationCode(e domain.User) (int, error) {

	if us.isUserVerified(e.ID) {
		return 0, errors.New("user is already verified")
	}

	code, err := us.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	_, err = us.Repo.UpdateUser(e.ID, domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	})

	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	// send SMS

	return code, nil
}

func (us UserService) VerifyCode(id uint, code int) error {

	if us.isUserVerified(id) {
		return errors.New("user is already verified")
	}

	user, err := us.Repo.FindUserById(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code has expired")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = us.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("unable to verify user")
	}

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
