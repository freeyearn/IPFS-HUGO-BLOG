package services

import (
	"IPFS-Blog-Hugo/common/hugo"
	"IPFS-Blog-Hugo/common/ipfs"
	"IPFS-Blog-Hugo/utils/message"
	"github.com/spf13/viper"
)

// CompileAndUpload compile the blogs with hugo and then publish it with ipfs.
func CompileAndUpload() {
	dir, err := hugo.GetDir(viper.GetString("blog.Dir"))
	if err != nil {
		message.PrintErr(err)
		return
	}
	_, err = hugo.Build(dir)
	if err != nil {
		message.PrintErr(err)
	}
	ipfsNode := ipfs.GetIpfs(viper.GetString("ipfs.Url"))
	hash, err := ipfsNode.UploadDir(hugo.GetPublicDirPath(dir))
	if err != nil {
		message.PrintErr(hash, err)
	}
	message.Println("Blog published with CID:", hash)

	// publish with ipns
	go func() {
		err := ipfsNode.Publish("", hash)
		if err != nil {
			message.PrintErr("Error: IPNS publish failed:", err)
		}
		message.Println("IPNS published success")
	}()
}
