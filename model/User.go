package model

import "gorm.io/gorm"

type User struct {
	// gorm.Model 提供了以下四个字段: ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	// 0-admin 1-reader
	Role int `gorm:"type:int" json:"role"`
}
