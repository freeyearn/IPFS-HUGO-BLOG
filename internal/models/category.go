package models

import (
	"IPFS-Blog-Hugo/utils/structs"
	"time"
)

type Category struct {
	AutoId       int64     `gorm:"column:auto_id;primary_key;AUTO_INCREMENT" json:"auto_id" form:"auto_id"`
	CategoryId   string    `gorm:"column:category_id;type:varchar(22);" json:"category_id" form:"category_id"`
	UserId       string    `gorm:"column:user_id;type:varchar(22);" json:"user_id" form:"user_id"`
	CategoryName string    `gorm:"column:category_name;type:varchar(255);" json:"category_name" form:"category_name"`
	IsDeleted    int       `gorm:"column:is_deleted;type:tinyint(1);" json:"is_deleted" form:"is_deleted"`
	CreateTime   time.Time `gorm:"column:create_time;type:DATETIME;not null;default:CURRENT_TIMESTAMP;" json:"create_time" form:"create_time"`
}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) GetAutoId() int64 {
	return c.AutoId
}

func (c *Category) SetAutoId(autoId int64) {
	c.AutoId = autoId
}

func (c *Category) GetCategoryId() string {
	return c.CategoryId
}

func (c *Category) SetCategoryId(categoryId string) {
	c.CategoryId = categoryId
}

func (c *Category) GetUserId() string {
	return c.UserId
}

func (c *Category) SetUserId(userId string) {
	c.UserId = userId
}

func (c *Category) GetCategoryName() string {
	return c.CategoryName
}

func (c *Category) SetCategoryName(CategoryName string) {
	c.CategoryName = CategoryName
}

func (c *Category) GetIsDeleted() int {
	return c.IsDeleted
}

func (c *Category) SetIsDeleted(isDeleted int) {
	c.IsDeleted = isDeleted
}

func (c *Category) GetCreateTime() time.Time {
	return c.CreateTime
}

func (c *Category) SetCreateTime(createTime time.Time) {
	c.CreateTime = createTime
}

func (c *Category) GetModelMap() (map[string]interface{}, error) {
	return structs.StructToMap(c, "json")
}

func (c *Category) Assign(in interface{}) {
	structs.StructAssign(c, in, "json")
}
