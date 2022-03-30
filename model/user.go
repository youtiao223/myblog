package model

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"io"
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

// InsertUser 插入用户 //todo 密码加密
func InsertUser(user *User) int {
	// 检查用户名是否存在
	isExit := SelectUserByName(user)
	if isExit {
		return errors.ErrorUserExits
	}
	user.Password = Encrypt(user.Password)
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

// Encrypt md5加盐加密
func Encrypt(data string) string {
	h := md5.New()
	_, _ = io.WriteString(h, data)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	// 指定两个salt
	salt1 := "&*o5"
	salt2 := "^&*()"

	// salt1+用户名+salt2+MD5拼接
	_, _ = io.WriteString(h, salt1)
	_, _ = io.WriteString(h, "abc")
	_, _ = io.WriteString(h, salt2)
	_, _ = io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	return last
}

// GetUserByNamePwd 根据用户名和密码查询用户
func GetUserByNamePwd() {

}

// GetUserInfo 获取登录用户信息 todo
func GetUserInfo() {

}
