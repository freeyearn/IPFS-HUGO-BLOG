package cmd

import (
	"IPFS-Blog-Hugo/apis"
	"IPFS-Blog-Hugo/common/ipfs"
	"IPFS-Blog-Hugo/internal/crontab"
	"IPFS-Blog-Hugo/utils"
	"IPFS-Blog-Hugo/utils/errHelper"
	"IPFS-Blog-Hugo/utils/message"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "project_layout",
	Short: "IPFS-Blog-Hugo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Start doing things.开始做事情
		message.Println("Start Server")
		utils.GetWaitGroup().Add(1)

		// check ipfs node
		ipfsNode := ipfs.GetIpfs(viper.GetString("ipfs.Url"))
		if !ipfsNode.Status() {
			message.PrintErr("Error: ipfs node is not running! please have a check")
			message.Exit()
		}

		go apis.StartHttp()

		// 定时任务模块
		crontab.InitCrontab()
		cron := crontab.GetCrontab()
		cron.Start()
		message.Println("定时任务模块加载完成")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	errHelper.ErrExit(rootCmd.Execute())
}

func init() {
	//rootCmd.PersistentFlags().StringP("Port", "P", "8000", "配置文件名(注意-C为大写)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	//port, err := rootCmd.Flags().GetString("Port")
	//errHelper.ErrExit(err)
}
