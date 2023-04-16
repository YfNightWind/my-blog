package model

import (
	"gorm.io/gorm"
	"my-blog/utils/errormsg"
)

type Comment struct {
	gorm.Model
	UserId       uint   `json:"user_id"`
	ArticleId    uint   `json:"article_id"`
	ArticleTitle string `json:"article_title"`
	UserName     string `json:"username"`
	Content      string `gorm:"varchar(500);not null" json:"content"`
	Status       int    `json:"status"`
}

// NewComment 新增评论
func NewComment(comment *Comment) int {
	err := db.Create(&comment).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// GetComment 查询单个评论
func GetComment(id int) (Comment, int) {
	var comment Comment

	err := db.Where("id = ?", id).Error
	if err != nil {
		return comment, errormsg.ERROR
	}

	return comment, errormsg.SUCCESS
}

// GetCommentList 获取评论列表(后台)
func GetCommentList(pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64

	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}

	err := db.
		Model(&commentList).
		Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status, comment.created_at, comment.deleted_at").
		Limit(pageSize).
		Offset(offSet).
		Order("Created_At DESC").Error
	if err != nil {
		return commentList, 0, errormsg.ERROR
	}

	return commentList, total, errormsg.SUCCESS
}

// GetCommentNumber 获取评论数量
func GetCommentNumber(id int) int64 {
	var comment Comment
	var total int64

	db.Find(&comment).Where("article_id = ? ", id).Where("status = 1").Count(&total)

	return total
}

// DeleteComment 删除评论
func DeleteComment(id int) int {
	var comment Comment

	err := db.Where("id = ? ", id).Delete(&comment).Error
	if err != nil {
		return errormsg.ERROR
	}

	return errormsg.SUCCESS
}

// ArticleGetCommentList 展示文章底下的评论
func ArticleGetCommentList(id int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	offSet := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSet = -1
	}

	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = 1").Count(&total)

	err := db.Model(&Comment{}).
		Limit(pageSize).
		Offset(offSet).
		Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status, comment.created_at, comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").
		Where("article_id = ?", id).
		Where("status = 1").
		Scan(&commentList).Error

	if err != nil {
		return commentList, 0, errormsg.ERROR
	}

	return commentList, total, errormsg.SUCCESS
}

// PassTheComment 通过评论
