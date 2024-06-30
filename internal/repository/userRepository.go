package repository

import (
	"errors"
	"github.com/programmerolajide/go-ecommerce/internal/domain"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	updateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf(" Error occurred while creating user %v", err)
		return domain.User{}, errors.New("Failed to create user")
	}

	return user, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {

	return domain.User{}, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {

	return domain.User{}, nil
}

func (r userRepository) updateUser(id uint, user domain.User) (domain.User, error) {

	return domain.User{}, nil
}
