package model

import (
	"gorm.io/gorm"
	"myBlog/utils/errors"
)

type Article struct {
	Category Category `gorm:"foreignKey:ID;references:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Cid     uint   `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(200)" json:"img"`
}

// SelectArt 分页查询文章
func SelectArt(pageNum int, pageSize int) ([]Article, int64) {
	var articles []Article
	var count int64
	db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Count(&count)
	return articles, count
}

// InsertArt 插入文章
func InsertArt(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// DelArtById 根据id删除文章
func DelArtById(id uint) int {
	rowsAffected := db.Delete(&Article{}, id).RowsAffected
	if rowsAffected == 0 {
		return errors.ErrorArtIdNotExits
	}
	return errors.SUCCESS
}

// UpdateArtById 根据id更新文章
func UpdateArtById(id uint, data *Article) int {
	var article = make(map[string]interface{})
	article["title"] = data.Title
	article["cid"] = data.Cid
	article["desc"] = data.Desc
	article["img"] = data.Img
	article["Content"] = data.Content
	err := db.Model(&Article{}).Where("id=?", id).Updates(article).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// SelectArtById 根据Id查找文章
func SelectArtById(id uint, article *Article) int {
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return errors.ErrorArtIdNotExits
	}
	return errors.SUCCESS
}

// SelectArtByCate 根据Cid查找文章
func SelectArtByCate(cid uint, pageNum int, pageSize int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Where("cid = ?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		return articles, errors.ErrorArtIdNotExits
	}
	return articles, errors.SUCCESS
}
