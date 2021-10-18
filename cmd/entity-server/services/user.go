package services

import (
	"errors"
	"time"

	"github.com/golang/cmd/entity-server/models"
	"github.com/golang/internal/pkg/utils"
)

type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
}
type UserRes struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func toUserRes(user *models.User) *UserRes {
	var userRes = &UserRes{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
	return userRes
}

func (obj *User) Login() (*UserRes, error) {
	model := models.User{
		Email: obj.Email,
	}
	user, err := model.Login()
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(obj.Password, user.Password) {
		return nil, errors.New("Login failed")
	}
	return toUserRes(user), nil
}

func (obj *User) Register() (bool, error) {
	passHash, hashErr := utils.HashPassword(obj.Password)
	if hashErr != nil {
		return false, hashErr
	}
	model := models.User{
		Email:    obj.Email,
		Password: passHash,
		Name:     obj.Name,
	}
	isExists, _ := model.IsEmailExist()
	if isExists {
		return false, errors.New("Email is not available!")
	}
	if _, err := model.Register(); err == nil {
		return true, err
	} else {
		return false, err
	}
}
