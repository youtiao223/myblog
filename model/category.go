package model

import (
	"myBlog/utils/errorUtils"
)

type Category struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CateName string `gorm:"type:varchar(20)" json:"name"`
}

// SelectCateByName 根据分类名查询用户
// true : 用户存在
// false : 用户不存在
func SelectCateByName(name string) bool {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {
		return true
	}
	return false
}

// SelectCate 分页查询分类
func SelectCate(pageNum int, pageSize int) ([]Category, int64) {
	var cates []Category
	var count int64
	db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&count)
	return cates, count
}

// SelectCateById 根据Id查找分类
func SelectCateById(id uint, category *Category) int {
	err := db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return errorUtils.ErrorArtIdNotExits
	}
	return errorUtils.SUCCESS
}

// InsertCate 插入Cate
func InsertCate(cate *Category) int {
	// 检查分类名是否存在
	isExit := SelectCateByName(cate.CateName)
	if isExit {
		return errorUtils.ErrorCateExits
	}
	err := db.Create(&cate).Error
	if err != nil {
		return errorUtils.ERROR
	}
	return errorUtils.SUCCESS
}

// DelCateById 根据id删除Cart
func DelCateById(id uint) int {
	rowsAffected := db.Delete(&Category{}, id).RowsAffected
	if rowsAffected == 0 {
		return errorUtils.ErrorCateIdNotExits
	}
	return errorUtils.SUCCESS
}

// UpdateCateById 根据id更新Cart
func UpdateCateById(id uint, data *Category) int {
	isExit := SelectCateByName(data.CateName)
	if isExit {
		return errorUtils.ErrorCateExits
	}
	var newUser = make(map[string]interface{})
	newUser["name"] = data.CateName
	err := db.Model(&Category{}).Where("id=?", id).Updates(newUser).Error
	if err != nil {
		return errorUtils.ERROR
	}
	return errorUtils.SUCCESS
}
