package database

import (
	"IPFS-Blog-Hugo/internal/models"
	"IPFS-Blog-Hugo/utils/message"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	mysqlClient *gorm.DB
)

func GetMysqlClient() *gorm.DB {
	return mysqlClient
}

func InitMysqlClient(url string) (err error) {
	mysqlClient, err = gorm.Open("mysql", url)
	if err != nil {
		return
	}
	err = InitTables()
	if err != nil {
		return
	}
	return
}

func InitTables() (err error) {
	message.Println("start init tables")
	if mysqlClient == nil {
		err = errors.New("client is not initialized")
	}
	mysqlClient.AutoMigrate(&models.User{}, &models.Article{}, &models.Category{})
	message.Println("init tables finished")
	return
}
