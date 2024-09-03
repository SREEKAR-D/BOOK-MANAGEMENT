package service

import (
	"errors"
	"time"

	"golang2/model"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(user UserDTO) error
	Login(user UserDTO) (string, error)
}

type userService struct {
	repo model.UserRepository
}

func NewUserService(db *model.DB) UserService {
	return &userService{repo: model.NewUserRepository(db)}
}

func (s *userService) SignUp(user UserDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//above hashedPassword is of type byte[] we need to change it to string down the way to use it
	if err != nil {
		return err
	}
	return s.repo.AddUser(model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	})
}

func (s *userService) Login(user UserDTO) (string, error) {
	storedUser, err := s.repo.GetUserByUsername(user.Username)
	if err != nil {
		return "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	token, err := generateJWT(storedUser.ID, user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateJWT(userID uuid.UUID, user UserDTO) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"userName": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 5).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

type UserDTO struct {
	Username string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
