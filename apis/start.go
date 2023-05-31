package apis

import (
	"IPFS-Blog-Hugo/internal/engine"
	"IPFS-Blog-Hugo/internal/settings"
	"IPFS-Blog-Hugo/utils"
	"IPFS-Blog-Hugo/utils/message"
	"encoding/gob"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func StartHttp() {
	defer utils.GetWaitGroup().Done()
	gob.Register(time.Time{})
	var err error

	//初始化数据库（mysql、redis）
	message.Println("数据库初始化")
	err = settings.InitDatabase()
	if err != nil {
		message.PrintErr(err)
		return
	}
	message.Println("数据库初始化完成")

	//初始化gin引擎
	engine, err := engine.InitGinEngine()
	if err != nil {
		message.PrintErr(err)
		return
	}

	//开始运行
	err = engine.Run(fmt.Sprintf("%s:%s", viper.GetString("system.SysIP"), viper.GetString("system.SysPort")))
	if err != nil {
		message.PrintErr(err)
		return
	}
}
