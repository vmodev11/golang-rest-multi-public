package models

import (
	"errors"

	"github.com/golang/internal/pkg/utils"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (obj *User) IsEmailExist() (bool, error) {
	var tmpObj User
	err := db.Where("email = ?", obj.Email).First(&tmpObj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tmpObj.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (obj *User) IsExist(id uint) (bool, error) {
	err := db.Where("id = ?", id).Find(&obj).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if obj.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (obj *User) Login() (*User, error) {
	var tmpObj User
	err := db.Where("email = ?", obj.Email).First(&tmpObj).Error
	if err != nil {
		return nil, err
	}
	db.Model(&tmpObj).Update(obj)
	//response
	resObj, err := tmpObj.FindOne(tmpObj.ID)
	if err != nil {
		return nil, err
	}
	return resObj, nil
}

func (obj *User) Register() (*User, error) {

	if err := db.Create(&obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (obj *User) GetEmail(email string) (*User, error) {
	err := db.
		Where("email = ? AND active = 1", email).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (obj *User) FindOne(id uint) (*User, error) {
	err := db.Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return obj, err
}

func (obj *User) Add() (*User, error) {
	if err := db.Create(&obj).Error; err != nil {
		return nil, err
	}
	obj, err = obj.FindOne(obj.ID)
	if err != nil {
		return nil, err
	}
	return obj, err
}

func (obj *User) Update(id uint) (*User, error) {
	var tmpObj User
	if err := db.Where("id = ?", id).Find(&tmpObj).Error; err != nil {
		return nil, err
	}
	db.Model(&tmpObj).Update(obj)
	obj, err = obj.FindOne(obj.ID)
	if err != nil {
		return nil, err
	}
	return obj, err
}

func (obj *User) Delete(id uint) (*User, error) {
	resObj, err := obj.FindOne(id)
	if err != nil {
		return nil, err
	}
	if resObj.ID > 0 {

		if err := db.Where("id = ?", resObj.ID).Delete(&resObj).Error; err != nil {
			return nil, err
		}
		return resObj, nil
	}
	return nil, errors.New("User does not exist")
}

func (obj *User) GetIn(ids []uint) ([]*User, error) {
	var list []*User
	err := db.Where("id in (?)", ids).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

func (obj *User) UpdatePassword(id uint) (bool, error) {
	var tmpObj User
	err := db.Where("id = ?", id).First(&tmpObj).Error
	if err != nil {
		return false, err
	}
	db.Model(&tmpObj).Update(obj)
	return true, nil
}
func (obj *User) MigrateData() {
	passHash, _ := utils.HashPassword("123456")
	obj = &User{Model: gorm.Model{ID: 1}, Email: "admin@vmodev.com", Password: passHash, Name: "Admin"}
	if err := db.Unscoped().FirstOrCreate(obj).Error; err != nil {
		log.Error("Migrate Data User Error ", err)
	}
}
