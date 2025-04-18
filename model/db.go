package model

import (
	"fmt"
	"time"

	"github.com/YfNightWind/my-blog/utils"
	"github.com/YfNightWind/my-blog/utils/mylog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitDb 用于连接配置数据库
var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
		// 打印SQL语句
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		mylog.Error("连接数据库出错")
	}

	// 自动迁移
	err = db.AutoMigrate(&User{}, &Category{}, &Article{}, &Profile{}, &Comment{})
	if err != nil {
		mylog.Error("自动迁移出错")
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。不要超过Gin的超时时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
