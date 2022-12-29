package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/RomanVas30/store/internal/entities"
	"time"

	"github.com/RomanVas30/store/internal/storage"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(userCred entities.UserCred) (string, error)
	ParseToken(token string) (int, string, error)
	ChangePassword(changePass entities.ChangePassword) error
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserRole string `json:"role"`
}

type AuthService struct {
	repo storage.Authorization
}

func NewAuthService(repo storage.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entities.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(userCred entities.UserCred) (string, error) {
	user, err := s.repo.GetUser(userCred.Username, generatePasswordHash(userCred.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.UserRole, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ChangePassword(changePass entities.ChangePassword) error {
	changePass.OldPassword = generatePasswordHash(changePass.OldPassword)
	changePass.NewPassword = generatePasswordHash(changePass.NewPassword)

	err := s.repo.ChangePassword(changePass)
	if err != nil {
		return err
	}

	return nil
}
