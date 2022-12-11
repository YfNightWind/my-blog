package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"my-blog/utils"
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

// =============
// 对数据库的操作👇
// =============

// IsUserExist 查询用户是否存在
func IsUserExist(name string) (code int) {
	var user User
	db.Select("id").Where("username = ? ", name).Find(&user) // SELECT * FROM user LIMIT 1;
	if user.ID > 0 {
		return errormsg.ERROR_USERNAME_USED // 1001
	}

	return errormsg.SUCCESS // 200
}

// AddUser 添加用户(注册)
func AddUser(data *User) int {
	// 写进数据库之前，需要将密码进行加密
	// 通过钩子函数来实现
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}

	return errormsg.SUCCESS // 200
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	err = db.Where("id = ? ", id).Delete(&User{}).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// GetUserList 查询用户列表
func GetUserList(pageSize int, pageNum int) []User {
	var userList []User
	// 分页
	// gorm中"Cancel offset condition with -1"
	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}

	err = db.Limit(pageSize).Offset(offSet).Find(&userList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return userList
}

// EditUser 编辑用户
func EditUser(id int) {

}

// ScryptPassword 密码加密
func ScryptPassword(password string) string {
	const keyLen = 10

	key, err := scrypt.Key([]byte(password), utils.Salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	FinalPassword := base64.StdEncoding.EncodeToString(key)

	return FinalPassword
}

// BeforeSave 开始事务前，由GORM处理
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	// 密码加密
	u.Password = ScryptPassword(u.Password)
	return
}
