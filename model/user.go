package model

import (
	"gorm.io/gorm"
	"myBlog/utils/errors"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	// 用户权限 0:普通用户 1:管理员用户
	Role int `gorm:"type:int;not null;default:0" json:"role"`
}

// SelectUserByName 根据用户名查询用户
// true : 用户存在
// false : 用户不存在
func SelectUserByName(user *User) bool {
	rowsAffected := db.First(&user, "username = ?", user.Username).RowsAffected
	if rowsAffected >= 1 {
		return true
	}
	return false
}

// SelectUsers 分页查询用户信息
func SelectUsers(pageNum int, pageSize int) []User {
	var users []User
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	return users
}

// InsertUser 插入用户
func InsertUser(user *User) int {
	// 检查用户名是否存在
	isExit := SelectUserByName(user)
	if isExit {
		return errors.ErrorUserExits
	}
	err := db.Create(&user).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// DelUser 删除用户
func DelUser(user *User) int {
	// 检查用户名是否存在
	isExit := SelectUserByName(user)
	if isExit {
		return errors.ErrorUserExits
	}
	err := db.Delete(&user).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// GetUserByNamePwd 根据用户名和密码查询用户
func GetUserByNamePwd() {

}

// GetUserInfo 获取登录用户信息 todo
func GetUserInfo() {

}
