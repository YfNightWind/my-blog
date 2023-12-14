package model

import "github.com/YfNightWind/my-blog/utils/errormsg"

type Profile struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(20)" json:"name"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Github      string `gorm:"type:varchar(200)" json:"github"`
	Email       string `gorm:"type:varchar(200)" json:"email"`
	Img         string `gorm:"type:varchar(200)" json:"img"`
	Avatar      string `gorm:"type:varchar(200)" json:"avatar"`
	Bili        string `gorm:"type:varchar(200)" json:"bili"`
	IcpRecord   string `gorm:"type:varchar(200)" json:"icp_record"`
}

// GetProfile 获取个人信息
func GetProfile(id int) (Profile, int) {
	var profile Profile

	err := db.Where("ID = ? ", id).Find(&profile).Error

	if err != nil {
		return profile, errormsg.ERROR
	}

	return profile, errormsg.SUCCESS
}

// UpdateProfile 更新个人信息
func UpdateProfile(id int, data *Profile) int {
	var profile Profile

	err := db.Model(&profile).Where("ID = ? ", id).Updates(&data).Error

	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}
