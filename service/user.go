package service

import (
	"encoding/base64"
	mapset "github.com/deckarep/golang-set"
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/common/utils"
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

	encryptedPass, err := utils.DerivePassphrase(user.Password, 32)
	if err != nil {
		return errutils.PasswordEncryptError()
	}
	user.Password = base64.StdEncoding.EncodeToString(encryptedPass)
	//user.Password = utils.B2str(encryptedPass)
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
	columns := []string{
		"Username", "Email", "Nickname",
		"Avatar", "home_page", "Description",
		"Status", "Roles", "created_at"}
	if db.Select(columns).Find(&users).RecordNotFound() {
		return nil, errutils.DBOperationsFailed()
	} else {
		return users, nil
	}
}

func (service *userService) GetUser(db *gorm.DB, username string) (*model.User, *errutils.ZinError) {
	columns := []string{
		"Username", "Email", "Nickname",
		"Avatar", "home_page", "Description",
		"Status", "Roles", "created_at"}
	var user model.User
	if db.Select(columns).Where("Username = ?", username).First(&user).RecordNotFound() {
		return nil, errutils.UserNotFound(username)
	}
	return &user, nil
}

func (service *userService) DeleteUsers(db *gorm.DB, users []model.User) *errutils.ZinError {
	needDeleteUsers := mapset.NewSet()
	print(needDeleteUsers)
	for i := 0; i < len(users); i++ {
		needDeleteUsers.Add(users[i].Username)
	}
	if err := db.Where("Username in (?)", needDeleteUsers.ToSlice()).Delete(&model.User{}).Error; err != nil {
		return errutils.DBOperationsFailed()
	}
	return nil
}

func (service *userService) DeleteUser(db *gorm.DB, username string) *errutils.ZinError {
	if err := db.Where("Username = ?", username).Delete(&model.User{}).Error; err != nil {
		return errutils.DBOperationsFailed()
	}
	return nil
}
