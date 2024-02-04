package dao

import (
	"gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{DB: db}
}

// CheckUserName 检查用户名是否存在
func CheckUserName(username string) (user *model.User, exist bool, err error) {
	var count int64
	err = DB.Model(&model.User{}).Where("name = ?", username).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

// CheckUserPassword 检查用户密码是否存在
func CheckUserPassword(username, password string) (user *model.User, exist bool, err error) {
	var count int64
	err = DB.Model(&model.User{}).Where("name = ?", username).Find(&user).Count(&count).Error
	if password != user.Password {
		return nil, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func CreateUser(user *model.User) error {
	return DB.Model(&model.User{}).Create(&user).Error
}

// GetUserById 根据ID获取用户
func GetUserById(id uint) *model.User {
	var user model.User
	DB.Model(&model.User{}).Where("id = ?", id).Find(&user)
	return &user
}

// UpdateUser 更新用户
func UpdateUser(u *model.User) error {
	return DB.Model(&model.User{}).Where("id = ?", u.ID).Updates(&u).Error
}
