package init

import (
	"IPFS-Blog-Hugo/utils"
	"IPFS-Blog-Hugo/utils/loggers"
	"IPFS-Blog-Hugo/utils/message"
	"IPFS-Blog-Hugo/utils/security"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

func Init() *sync.WaitGroup {
	// 消息初始化
	waitGroup := MessageInit()

	// 配置读取初始化
	err := ViperInit()
	if err != nil {
		message.PrintErr(err)
		message.Exit()
	}

	// 日志初始化
	LoggerInit(viper.GetString("log.type"))

	// 安全模块初始化
	err = SecurityInit(viper.GetInt("system.RSABit"))
	if err != nil {
		message.PrintErr(err)
		message.Exit()
	}

	return waitGroup
}

// MessageInit message init
func MessageInit() *sync.WaitGroup {
	waitGroup := utils.GetWaitGroup()
	waitGroup.Add(1)
	go message.InitMsg()
	return waitGroup
}

// ViperInit viper object init
func ViperInit() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs") // 添加搜索路径
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: ", err)
		return
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file:", e.Name, "Op: ", e.Op)
	})

	return
}

// LoggerInit Log object init
func LoggerInit(logType string) {
	loggers.InitLogger(logType)
}

// SecurityInit  read key from config file
func SecurityInit(bit int) error {
	security.InitRSAHelper(viper.GetString("system.RSAPublic"), viper.GetString("system.RSAPrivate"))
	if !utils.PathExists(viper.GetString("system.RSAPublic")) || !utils.PathExists(viper.GetString("system.RSAPrivate")) {
		err := security.RSAHelper.GenerateRSAKey(bit)
		if err != nil {
			fmt.Println("Generate rsa pem err:", err)
			return err
		}
	}
	return nil
}
