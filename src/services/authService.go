package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wavekanit/book-store-backend/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret []byte
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	var users []models.Users
	result := s.db.Where("username = ?", username).Find(&users)
	if result.Error != nil {
		return "", result.Error
	}

	if len(users) == 0 || bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password)) != nil {
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Register(user *models.Users) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
