package settings

import (
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/utils/message"
)

func InitDatabase() (err error) {
	v := GetDatabaseViper()
	err = database.InitMysqlClient(v.GetString("mysql.dsn"))
	if err != nil {
		message.PrintErr("mysql初始化出错:", err)
		return
	}
	return
}
