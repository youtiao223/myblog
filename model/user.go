package model

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"io"
	"myBlog/middleware"
	"myBlog/utils/errors"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(35);not null" json:"password"`
	// 用户权限 0:普通用户 1:管理员用户
	Role int `gorm:"type:int;not null;default:0" json:"role"`
}

// SelectUserByName 根据用户名查询用户
// true : 用户存在
// false : 用户不存在
func SelectUserByName(username string) bool {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
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
	isExit := SelectUserByName(user.Username)
	if isExit {
		return errors.ErrorUserNameExits
	}
	user.Password = Encrypt(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// SelectUserById 通过id查找用户
// true : 用户存在
// false : 用户不存在
func SelectUserById(id uint) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

// DelUserById DelUser 删除用户
func DelUserById(id uint) int {
	rowsAffected := db.Delete(&User{}, id).RowsAffected
	if rowsAffected == 0 {
		return errors.ErrorUserIdNotExits
	}
	return errors.SUCCESS
}

// UpdateUserById UpdateUser 更新用户信息
// 修改用户名要检查用户名是否重名
// 密码修改独立，不在这里修改
func UpdateUserById(id uint, data *User) int {
	isExit := SelectUserById(id)
	if !isExit {
		return errors.ErrorUserIdNotExits
	}
	isExit = SelectUserByName(data.Username)
	if isExit {
		return errors.ErrorUserNameExits
	}
	var newUser = make(map[string]interface{})
	newUser["username"] = data.Username
	newUser["role"] = data.Role
	err := db.Model(&User{}).Where("id=?", id).Updates(newUser).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// CheckLogin 登录验证
func CheckLogin(loginUser *User) (string, int) {
	var user User
	username := loginUser.Username
	password := loginUser.Password
	pwdEncrypted := Encrypt(password)
	db.Select("id").Where("username = ? and password = ?", username, pwdEncrypted).First(&user)

	if user.ID == 0 {
		return "", errors.ErrorNameOrPwd
	}
	// 登录成功后生成token
	token, code := middleware.GenToken(username)
	if code == errors.ERROR {
		return "", errors.ERROR
	}
	return token, errors.SUCCESS
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
