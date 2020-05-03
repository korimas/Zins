package service

import (
	"encoding/base64"
	set "github.com/deckarep/golang-set"
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/api/jsfmt"
	cons "github.com/zpdev/zins/common/constance"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/common/utils"
	"github.com/zpdev/zins/model"
	"time"
)

var UserService = &userService{}

type userService struct {
}

func (sec *userService) CreateUser(db *gorm.DB, user *model.User) *errutils.ZinError {
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
	user.CreatedTime = time.Now().Unix()
	user.Status = cons.ACTIVE
	if err := db.Create(user).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}
	return nil
}

//
//func (sec *userService) GetUsers(db *gorm.DB) ([]*model.User, *errutils.ZinError) {
//	users := make([]*model.User, 0)
//	columns := []string{"ID", "Username", "Email", "Nickname", "Description", "Status", "Role", "created_time"}
//	if err := db.Select(columns).Find(&users).Error; err != nil {
//		return nil, errutils.DBOperationsFailed(err.Error())
//	} else {
//		return users, nil
//	}
//}

func (sec *userService) GetUsers(db *gorm.DB, query *jsfmt.Query) ([]*model.User, int, *errutils.ZinError) {
	users := make([]*model.User, 0)
	total := 0
	Zerr := &errutils.ZinError{}
	columns := []string{"ID", "Username", "Email", "Nickname", "Description", "Status", "Role", "created_time"}
	if total, Zerr = query.Find(db.Select(columns), &model.User{}, &users); Zerr != nil {
		return nil, 0, Zerr
	}

	return users, total, nil
}

func (sec *userService) GetUser(db *gorm.DB, username string) (*model.User, *errutils.ZinError) {
	columns := []string{"ID", "Username", "Email", "Nickname", "Description", "Status", "Role", "created_time"}
	var user model.User
	if db.Select(columns).Where("Username = ?", username).First(&user).RecordNotFound() {
		return nil, errutils.SpecifiedUserNotFound(username)
	}
	return &user, nil
}

func (sec *userService) DeleteUsers(db *gorm.DB, users []model.User) *errutils.ZinError {
	needDeleteUsers := set.NewSet()
	for i := 0; i < len(users); i++ {
		needDeleteUsers.Add(users[i].Username)
	}
	if err := db.Where("Username in (?)", needDeleteUsers.ToSlice()).Delete(&model.Token{}).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}

	if err := db.Where("Username in (?)", needDeleteUsers.ToSlice()).Delete(&model.User{}).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}
	return nil
}

func (sec *userService) DeleteUser(db *gorm.DB, username string) *errutils.ZinError {
	if err := db.Where("Username = ?", username).Delete(&model.Token{}).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}

	if err := db.Where("Username = ?", username).Delete(&model.User{}).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}
	return nil
}
