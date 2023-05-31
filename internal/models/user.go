package models

import (
	"IPFS-Blog-Hugo/utils/structs"
	"time"
)

type User struct {
	AutoId     int64     `gorm:"column:auto_id;primary_key;AUTO_INCREMENT" json:"auto_id" form:"auto_id"`
	UserId     string    `gorm:"column:user_id;type:varchar(22);" json:"user_id" form:"user_id"`
	Wallet     string    `gorm:"column:wallet;type:varchar(255);" json:"wallet" form:"wallet"`
	Account    string    `gorm:"column:account;type:varchar(500);" json:"account" form:"account"`
	Password   string    `gorm:"column:password;type:varchar(500);" json:"password" form:"password"`
	Name       string    `gorm:"column:name;type:varchar(255);" json:"name" form:"name"`
	Profile    string    `gorm:"column:profile;type:varchar(500);" json:"profile" form:"profile"`
	Status     int       `gorm:"column:status;type:tinyint(2);" json:"status" form:"status"`
	IsDeleted  int       `gorm:"column:is_deleted;type:tinyint(1);" json:"is_deleted" form:"is_deleted"`
	CreateTime time.Time `gorm:"column:create_time;type:DATETIME;not null;default:CURRENT_TIMESTAMP;" json:"create_time" form:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:DATETIME;not null;default:CURRENT_TIMESTAMP on update current_timestamp;" json:"update_time" form:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetUserId() string {
	return u.UserId
}

func (u *User) SetUserId(userId string) {
	u.UserId = userId
}

func (u *User) GetWallet() string {
	return u.Wallet
}

func (u *User) SetWallet(wallet string) {
	u.Wallet = wallet
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetAccount() string {
	return u.Account
}

func (u *User) SetAccount(account string) {
	u.Account = account
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetStatus() int {
	return u.Status
}

func (u *User) SetStatus(status int) {
	u.Status = status
}

func (u *User) GetCreateTime() time.Time {
	return u.CreateTime
}

func (u *User) SetCreateTime(createTime time.Time) {
	u.CreateTime = createTime
}

func (u *User) GetUpdateTime() time.Time {
	return u.UpdateTime
}

func (u *User) SetUpdateTime(updateTime time.Time) {
	u.UpdateTime = updateTime
}

func (u *User) GetIsDeleted() int {
	return u.IsDeleted
}

func (u *User) SetIsDeleted(isDeleted int) {
	u.IsDeleted = isDeleted
}

func (u *User) GetModelMap() (map[string]interface{}, error) {
	return structs.StructToMap(u, "json")
}

func (u *User) Assign(in interface{}) {
	structs.StructAssign(u, in, "json")
}
