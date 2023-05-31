package engine

import (
	"IPFS-Blog-Hugo/apis/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGinEngine() (*gin.Engine, error) {
	gin.SetMode(viper.GetString("system.Mode"))
	engine := gin.Default()

	// 初始化Session
	// 将session存在redis
	//v := vipers.GetDatabaseViper()
	//store, err := redis.NewStore(10, "tcp", v.GetString("redis.addr"), v.GetString("redis.password"), []byte(viper.GetString("system.Secret")))
	//if err != nil {
	//	return nil, err
	//}
	// 将session存在cookie
	store := cookie.NewStore([]byte(viper.GetString("system.Secret")))

	store.Options(sessions.Options{
		MaxAge: viper.GetInt("system.SessionExpireTime"),
	})
	engine.Use(sessions.Sessions("mySession", store))

	routers.InitRouter(engine)

	return engine, nil
}
