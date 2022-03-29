package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Cid     int    `json:"cid"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
