package model

import (
	"errors"
	"gorm.io/gorm"
	"myBlog/utils/errorUtils"
)

type Profile struct {
	gorm.Model
	Uid              uint   `gorm:"type:int;not null;"  json:"uid"`
	Author           string `gorm:"type:varchar(20);not null;default:匿名"  json:"author"`
	Motto            string `gorm:"type:varchar(100);default:这家伙很懒,没有写"  json:"motto"`
	SelfIntroduction string `gorm:"type:varchar(200);default:这家伙很懒,没有写"  json:"selfIntroduction"`
	QQ               string `gorm:"type:varchar(15);"  json:"qq"`
	WeChat           string `gorm:"type:varchar(15);"  json:"weChat"`
	Github           string `gorm:"type:varchar(60);"  json:"github"`
	Email            string `gorm:"type:varchar(30);"  json:"email"`
	BackGroundImg    string `gorm:"type:varchar(60);"  json:"backGroundImg"`
	Avatar           string `gorm:"type:varchar(60);"  json:"avatar"`
}

// InsertProfile 新建Profile
func InsertProfile(uid uint) {
	var profile Profile
	profile.Uid = uid
	db.Create(&profile)
}

// FindProfileByUid 根据Uid查找是否有Profile
func FindProfileByUid(uid uint) bool {
	var profile Profile
	db.Select("id").Where("uid = ?", uid).First(&profile)
	if profile.ID > 0 {
		return true
	}
	return false
}

// SelectProfileByUid 根据Uid查找个人信息
func SelectProfileByUid(uid uint, profile *Profile) int {
	// 先检查管理员是否存在
	err := db.First(&User{}, uid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorUtils.ErrorAdminNotExist
	}
	// 查找管理员Profile
	rows := db.Where("uid = ?", uid).Find(&profile).RowsAffected
	if rows <= 0 {
		// 初始化
		InsertProfile(uid)
		db.Where("uid = ?", uid).Find(&profile)
	}

	return errorUtils.SUCCESS
}

// UpdateProfile 更新个人信息
func UpdateProfile(uid uint, data *Profile) int {
	if !FindProfileByUid(uid) {
		return errorUtils.ErrorAdminNotExist
	}
	var profile = make(map[string]interface{})
	profile["author"] = data.Author
	profile["motto"] = data.Motto
	profile["self_introduction"] = data.SelfIntroduction
	profile["qq"] = data.QQ
	profile["we_chat"] = data.WeChat
	profile["github"] = data.Github
	profile["email"] = data.Email
	profile["back_ground_img"] = data.BackGroundImg
	profile["avatar"] = data.Avatar
	err := db.Model(&Profile{}).Where("uid=?", uid).Updates(profile).Error
	if err != nil {
		return errorUtils.ERROR
	}
	return errorUtils.SUCCESS
}
