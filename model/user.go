package model

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/YfNightWind/my-blog/utils"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model 提供了以下四个字段: ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=15" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	// 0无权限，1为管理员
	Role int `gorm:"type:int;default:2" json:"role" validate:"required,lte=2" label:"角色码"`
	// 是否有权限查看工具页面 有权限为1，无权限为0
	IsAccessTools int `gorm:"type:int8;default:0" json:"is_access_tools" validate:"required,lte=2" label:"工具页面权限"`
}

// BeforeSave 开始事务前，由GORM处理
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	// 密码加密
	u.Password = ScryptPassword(u.Password)
	return
}

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

// DeleteUser 删除用户(soft delete)
func DeleteUser(id int) int {
	err := db.Where("id = ? ", id).Delete(&User{}).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// GetUser 查询单个用户
func GetUser(id int) (User, int) {
	var user User

	err := db.Where("id = ? ", id).Find(&user).Error
	if err != nil {
		return user, errormsg.ERROR
	}
	return user, errormsg.SUCCESS
}

// GetUserList 查询用户列表
func GetUserList(username string, pageSize int, pageNum int) ([]User, int64) {
	var userList []User
	var total int64
	// 分页
	// gorm中"Cancel offset condition with -1"
	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}

	// 查询所有用户
	if username == "" {
		//err = db.Limit(pageSize).Offset(offSet).Find(&userList).Count(&total).Error
		err := db.Select("id, username, role, created_at").Limit(pageSize).Offset(offSet).Find(&userList).Error
		db.Model(&userList).Count(&total)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0
		}

		return userList, total
	} else {
		// 模糊查询
		//err = db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset(offSet).Find(&userList).Count(&total).Error
		//

		err := db.Select("id, username, role, created_at ").Where("username LIKE ?", username+"%").
			Limit(pageSize).Offset(offSet).Find(&userList).Error
		db.Model(&userList).Where("username LIKE ?", username+"%").Count(&total)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0
		}
		return userList, total
	}

}

// UpdateUser 编辑用户信息时，防止不能修改用户名
func UpdateUser(id int, username string) int {
	var user User
	db.Select("id, username").Where("username = ? ", username).Find(&user)

	if user.ID == uint(id) {
		return errormsg.SUCCESS
	}

	if user.ID > 0 {
		return errormsg.ERROR_USERNAME_USED
	}

	return errormsg.SUCCESS
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var userMap = make(map[string]interface{})
	userMap["username"] = data.Username
	userMap["role"] = data.Role

	// 更新
	err := db.Model(&User{}).Where("id = ?", id).Updates(userMap).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {

	err := db.Select("password").Where("id = ? ", id).Updates(&data).Error

	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// VerifyAdminLogin 后台登录验证
func VerifyAdminLogin(username, password string) int {
	var user User

	db.Where("username = ?", username).Find(&user)

	// 用户名错误或不存在
	if user.ID == 0 {
		return errormsg.ERROR_USER_NOT_EXIST
	}

	//密码错误
	if ScryptPassword(password) != user.Password {
		return errormsg.ERROR_PASSWORD_WRONG
	}

	// 无权限登录后台
	if user.Role != 1 {
		return errormsg.ERROR_NO_PERMISSION
	}

	return errormsg.SUCCESS
}

// VerifyFrontLogin 前台登录验证
func VerifyFrontLogin(username, password string) (User, int) {
	var user User

	db.Where("username = ? ", username).Find(&user)

	// 用户名错误或不存在
	if user.ID == 0 {
		return user, errormsg.ERROR_USER_NOT_EXIST
	}

	//密码错误
	if ScryptPassword(password) != user.Password {
		return user, errormsg.ERROR_PASSWORD_WRONG
	}

	return user, errormsg.SUCCESS
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

// AdjustToolsPageAccess 调整用户是否可以查看工具类页面 TODO
func (u *User) AdjustToolsPageAccess(result int8) error {
	if u.ID == 0 {
		return errors.New("用户ID不得为空")
	}
	err := db.Model(u).Where("id = ?", u.ID).Update("is_access_tools", result).Error

	return err
}
