package model

import "gorm.io/gorm"

// LogOperation 日志字段
// 请求时间 CreatedAt
// 状态码 Code
// 请求IP Ip
// 客户端信息 UserAgent
// 请求大小 Size
// 请求方法 Method
// 请求路径 Type
type LogOperation struct {
	gorm.Model
	Type      uint   `gorm:"type:smallint(5); not null" json:"type"`
	Ip        string `gorm:"type:varchar(100); not null" json:"ip"`
	UserAgent string `gorm:"type:varchar(100); not null" json:"user_agent"`
	Size      uint   `gorm:"type:int(100); not null" json:"size"`
	Method    string `gorm:"type:varchar(20); not null" json:"method"`
	Code      uint   `gorm:"type:int(5); not null" json:"code"`
	HostName  string `gorm:"type:varchar(100); not null" json:"host_name"`
}

// CreateLogs 创建日志
//func CreateLogs(log *LogOperation) int {
//	db.Create(&log).Error
//}
//
//// GetLogs 从数据库中获取日志
//func GetLogs() []LogOperation {
//
//}
