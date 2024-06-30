package repository

import (
	"errors"
	"github.com/programmerolajide/go-ecommerce/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf(" Error occurred while finding user %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, id).Error

	if err != nil {
		log.Printf(" Error occurred while finding user %v", err)
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) updateUser(id uint, user domain.User) (domain.User, error) {

	var updatedUser domain.User

	err := r.db.Model(&updatedUser).Clauses(clause.Returning{}).Where("id=?", id).Updates(user).Error

	if err != nil {
		log.Printf(" Error occurred while finding user %v", err)
		return domain.User{}, errors.New("updating user failed, try again")
	}

	return updatedUser, nil
}
