package main

import (
	"IPFS-Blog-Hugo/cmd"
	initialization "IPFS-Blog-Hugo/init"
	"IPFS-Blog-Hugo/utils/message"
)

func main() {
	waitGroup := initialization.Init()

	// 执行cmd
	cmd.Execute()

	// 退出系统
	waitGroup.Wait()
	message.Exit()
}
