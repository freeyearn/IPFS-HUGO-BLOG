package crontab

import (
	"IPFS-Blog-Hugo/internal/services"
	"IPFS-Blog-Hugo/utils/message"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"sync"
)

var (
	crontab     *cron.Cron
	crontabOnce sync.Once
)

// GetCrontab get crontab object.
func GetCrontab() *cron.Cron {
	return crontab
}

// InitCrontab init crontab object
func InitCrontab() {
	crontabOnce.Do(func() {
		crontab = cron.New(cron.WithSeconds())
		// register functions
		_, err := crontab.AddFunc(viper.GetString("crontab.spec"), services.CompileAndUpload)
		if err != nil {
			message.PrintErr(err)
		}
	})
}
