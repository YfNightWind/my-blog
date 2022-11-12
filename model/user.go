package model

import (
	"gorm.io/gorm"
	"my-blog/utils/errormsg"
)

type User struct {
	// gorm.Model 提供了以下四个字段: ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required"`
	// 0-admin 1-reader
	Role int `gorm:"type:int" json:"role" binding:"required"`
}

//
// 对数据库的操作👇
//

// IsUserExist 查询用户是否存在
func IsUserExist(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user) // SELECT * FROM user LIMIT 1;
	if user.ID > 0 {
		return errormsg.ERROR_USERNAME_USED // 1001
	}

	return errormsg.SUCCESS // 200
}

// AddUser 添加用户(注册)
func AddUser(data *User) int {
	err := db.Create(data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}

	return errormsg.SUCCESS // 200
}
