package services

import (
	"todos/core"
	database "todos/db"

	"github.com/phuslu/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) GetUserByID(id int) (*core.User, error) {
	db := database.GetDB()
	user := core.User{}
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (s *UserService) GetUserByEmail(email string) (*core.User, error) {
	db := database.GetDB()
	user := core.User{}
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (s *UserService) CreateUser(email, rawPassword string) error {
	db := database.GetDB()
	password, _ := HashPassword(rawPassword)
	log.Info().Str("email", email).Str("password", "******").Msg("CreateUser")
	user := &core.User{Email: email, Password: password}
	return db.Create(user).Error
}

func (s *UserService) CheckPassword(user *core.User, password string) bool {
	return CheckPasswordHash(password, user.Password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
