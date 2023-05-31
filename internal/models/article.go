package models

import (
	"IPFS-Blog-Hugo/utils/structs"
	"time"
)

type Article struct {
	AutoId      int64     `gorm:"column:auto_id;primary_key;AUTO_INCREMENT" json:"auto_id" form:"auto_id"`
	ArticleId   string    `gorm:"column:article_id;type:varchar(22);" json:"article_id" form:"article_id"`
	Path        string    `gorm:"column:path;type:varchar(255);" json:"path" form:"path"`
	AuthorId    string    `gorm:"column:author_id;type:varchar(22);" json:"author_id" form:"author_id"`
	Title       string    `gorm:"column:title;type:varchar(255);" json:"title" form:"title"`
	Description string    `gorm:"column:description;type:text;" json:"description" form:"description"`
	Tags        string    `gorm:"column:tags;type:text;" json:"tags" form:"tags"`
	Category    string    `gorm:"column:category;type:text;" json:"category" form:"category"`
	Keyword     string    `gorm:"column:keyword;type:text;" json:"keyword" form:"keyword"`
	Next        string    `gorm:"column:next;type:varchar(255);" json:"next" form:"next"`
	Prev        string    `gorm:"column:prev;type:varchar(255);" json:"prev" form:"prev"`
	Status      int       `gorm:"column:status;type:tinyint(2);" json:"status" form:"status"`
	IsDeleted   int       `gorm:"column:is_deleted;type:tinyint(1);" json:"is_deleted" form:"is_deleted"`
	PublishTime time.Time `gorm:"column:publish_time;type:DATETIME;not null;default:CURRENT_TIMESTAMP;" json:"publish_time,omitempty" form:"publish_time"`
	UpdateTime  time.Time `gorm:"column:update_time;type:DATETIME;not null;default:CURRENT_TIMESTAMP on update current_timestamp;" json:"update_time,omitempty" form:"update_time"`
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) GetAutoId() int64 {
	return a.AutoId
}

func (a *Article) SetAutoId(id int64) {
	a.AutoId = id
}

func (a *Article) GetArticleId() string {
	return a.ArticleId
}

func (a *Article) SetArticleId(id string) {
	a.ArticleId = id
}

func (a *Article) GetPath() string {
	return a.Path
}

func (a *Article) SetPath(path string) {
	a.Path = path
}

func (a *Article) GetAuthorId() string {
	return a.AuthorId
}

func (a *Article) SetAuthorId(id string) {
	a.AuthorId = id
}

func (a *Article) GetTitle() string {
	return a.Title
}

func (a *Article) SetTitle(title string) {
	a.Title = title
}

func (a *Article) GetDescription() string {
	return a.Description
}

func (a *Article) SetDescription(description string) {
	a.Description = description
}

func (a *Article) GetTags() string {
	return a.Tags
}

func (a *Article) SetTags(tags string) {
	a.Tags = tags
}

func (a *Article) GetCategory() string {
	return a.Category
}

func (a *Article) SetCategory(category string) {
	a.Category = category
}

func (a *Article) GetKeyword() string {
	return a.Keyword
}

func (a *Article) SetKeyword(keyword string) {
	a.Keyword = keyword
}

func (a *Article) GetNext() string {
	return a.Next
}

func (a *Article) SetNext(next string) {
	a.Next = next
}

func (a *Article) GetPrev() string {
	return a.Prev
}

func (a *Article) SetPrev(prev string) {
	a.Prev = prev
}

func (a *Article) GetStatus() int {
	return a.Status
}

func (a *Article) SetStatus(status int) {
	a.Status = status
}

func (a *Article) GetIsDeleted() int {
	return a.IsDeleted
}

func (a *Article) SetIsDeleted(isDeleted int) {
	a.IsDeleted = isDeleted
}

func (a *Article) GetPublishTime() time.Time {
	return a.PublishTime
}

func (a *Article) SetPublishTime(publishTime time.Time) {
	a.PublishTime = publishTime
}

func (a *Article) GetUpdateTime() time.Time {
	return a.UpdateTime
}

func (a *Article) SetUpdateTime(updateTime time.Time) {
	a.UpdateTime = updateTime
}

func (a *Article) GetModelMap() (map[string]interface{}, error) {
	return structs.StructToMap(a, "json")
}

func (a *Article) Assign(in interface{}) {
	structs.StructAssign(a, in, "json")
}
