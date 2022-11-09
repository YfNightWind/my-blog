package model

import "gorm.io/gorm"

type Article struct {
	// Cid CategoryId和分类的id对应
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid   int    `gorm:"type:int;not null" json:"cid"`
	// Desc Description
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}
