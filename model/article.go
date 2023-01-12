package model

import (
	"gorm.io/gorm"
	"my-blog/utils/errormsg"
)

type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignKey:Cid"` // Cid CategoryId和分类的id对应
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int;not null" json:"cid"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Img         string   `gorm:"type:varchar(100)" json:"img"`
}

// =============
// 对数据库的操作👇
// =============

// IsArticleExist 查询分类是否存在
func IsArticleExist(name string) (code int) {
	var article Article
	db.Select("id").Where("name = ? ", name).Find(&article) // SELECT * FROM article LIMIT 1;
	if article.ID > 0 {
		return errormsg.ERROR_CATEGORYNAME_USED // 3001
	}

	return errormsg.SUCCESS // 200
}

// GetCategoryArticleList 查询分类下所有文章
func GetCategoryArticleList(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var categoryArticle []Article
	var total int64

	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}
	err := db.Preload("Category").Limit(pageSize).Offset(offSet).Where("cid = ?", id).Find(&categoryArticle).Count(&total).Error

	if err != nil {
		return nil, errormsg.ERROR_CATEGORY_NOT_EXIST, 0
	}

	return categoryArticle, errormsg.SUCCESS, total
}

// GetArticleInfo 查询文章信息
func GetArticleInfo(id int) (*Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error

	if err != nil {
		return nil, errormsg.ERROR_ARTICLE_NOT_EXIST
	}

	return &article, errormsg.SUCCESS
}

// CreateArticle 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR // 500
	}

	return errormsg.SUCCESS // 200
}

// GetArticleList 查询文章列表
func GetArticleList(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	// 分页
	// gorm中"Cancel offset condition with -1"
	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}

	if title == "" {
		err = db.Order("updated_at DESC").Preload("Category").Limit(pageSize).Offset(offSet).Find(&articleList).Count(&total).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errormsg.ERROR, 0
		}

		return articleList, errormsg.SUCCESS, total
	} else {
		// 模糊查询
		err = db.Order("updated_at DESC").Preload("Category").Where("title LIKE ? ", title+"%").Limit(pageSize).Offset(offSet).Find(&articleList).Count(&total).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errormsg.ERROR, 0
		}

		return articleList, errormsg.SUCCESS, total
	}

}

// EditArticle 编辑文章信息
func EditArticle(id int, data *Article) int {
	var articleMap = make(map[string]interface{})
	articleMap["title"] = data.Title
	articleMap["cid"] = data.Cid
	articleMap["description"] = data.Description
	articleMap["content"] = data.Content
	articleMap["img"] = data.Img

	// 更新
	err := db.Model(&Article{}).Where("id = ?", id).Updates(articleMap).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// DeleteArticle 删除分类
func DeleteArticle(id int) int {
	err = db.Where("id = ? ", id).Delete(&Article{}).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}
