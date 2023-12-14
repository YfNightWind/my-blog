package model

import (
	"errors"
	"github.com/YfNightWind/my-blog/utils/errormsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// IsCategoryExist 查询分类是否存在
func IsCategoryExist(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ? ", name).Find(&category) // SELECT * FROM category LIMIT 1;
	if category.ID > 0 {
		return errormsg.ERROR_CATEGORYNAME_USED // 3001
	}

	return errormsg.SUCCESS // 200
}

// CreateCategory 添加分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}

	return errormsg.SUCCESS // 200
}

// GetCategory 查询单个分类
func GetCategory(id int) (Category, int) {
	var category Category
	err := db.Where("id = ? ", id).Find(&category).Error

	if err != nil {
		return category, errormsg.ERROR_CATEGORY_NOT_EXIST
	}

	return category, errormsg.SUCCESS
}

// GetCategoryList 查询分类列表
func GetCategoryList(pageSize int, pageNum int) ([]Category, int64) {
	var categoryList []Category
	var total int64
	// 分页
	// gorm中"Cancel offset condition with -1"
	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}

	err := db.Limit(pageSize).Offset(offSet).Find(&categoryList).Count(&total).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0
	}
	return categoryList, total
}

// EditCategory 编辑分类信息
func EditCategory(id int, data *Category) int {
	var categoryMap = make(map[string]interface{})
	categoryMap["name"] = data.Name

	// 更新
	err := db.Model(&Category{}).Where("id = ?", id).Updates(categoryMap).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	err := db.Where("id = ? ", id).Delete(&Category{}).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}
