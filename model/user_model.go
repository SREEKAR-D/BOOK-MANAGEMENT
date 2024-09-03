package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Username string    `json:"userName" gorm:"unique;not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
}

type UserRepository interface {
	AddUser(user User) error
	GetUserByUsername(username string) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *DB) UserRepository {
	return &userRepository{db: db.GormDB}
}

func (u *userRepository) AddUser(user User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) GetUserByUsername(username string) (User, error) {
	var user User
	err := u.db.Where("userName = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

//`gorm:"type:uuid;default:uuid_generate_v4()"`
