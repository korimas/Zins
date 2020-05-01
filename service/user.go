package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"time"
)

var UserService = &userService{}

type userService struct {
}

func (service *userService) CreateUser(db *gorm.DB, user *model.User) *errutils.ZinError {
	var checkUser model.User
	if !db.Where("Email = ?", user.Email).First(&checkUser).RecordNotFound() {
		return errutils.EmailAlreadyExit(user.Email)
	}

	if !db.Where("Username = ?", user.Username).First(&checkUser).RecordNotFound() {
		return errutils.UserAlreadyExit(user.Username)
	}

	user.CreatedAt = time.Now().Unix()
	user.Status = "Active"
	if err := db.Create(user).Error; err != nil {
		print(err.Error())
		return errutils.DBOperationsFailed()
	}
	return nil
}

func (service *userService) GetUsers(db *gorm.DB) ([]*model.User, *errutils.ZinError) {
	users := make([]*model.User, 0)
	if db.Find(&users).RecordNotFound() {
		return nil, errutils.DBOperationsFailed()
	} else {
		return users, nil
	}
}

func (service *userService) GetUser(db *gorm.DB, username string) (*model.User, *errutils.ZinError) {
	var user model.User
	if db.Where("Username = ?", username).First(&user).RecordNotFound() {
		return nil, errutils.UserNotFound(username)
	}
	return &user, nil
}
