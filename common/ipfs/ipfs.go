package ipfs

import (
	"bytes"
	"context"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"os"
	"sync"
)

// IpfsNode is the node on a certain machine with ipfs-daemon running. Check the ipfs-daemon on your machine or you
// could set the ipfs url with a remote ipfs-daemon url.
// daemonUrl is the ipfs-daemon url, for example http://localhost:5001
type IpfsNode struct {
	daemonUrl string
	sh        *shell.Shell
}

var ipfsNode *IpfsNode
var once = sync.Once{}

// GetIpfs get ipfs with singleton function.
func GetIpfs(url string) *IpfsNode {
	once.Do(func() {
		sh := shell.NewShell(url)
		ipfsNode = &IpfsNode{
			daemonUrl: url,
			sh:        sh,
		}
	})
	return ipfsNode
}

// UploadStr upload a single string to ipfs with cid given back.
func (i *IpfsNode) UploadStr(str string) (hash string, err error) {
	hash, err = i.sh.Add(bytes.NewBufferString(str), shell.Pin(true))
	if err != nil {
		return
	}
	return
}

// UploadFile upload a file to ipfs with cid given back.
func (i *IpfsNode) UploadFile(path string) (hash string, err error) {
	file, err := os.Open(path)

	defer file.Close()

	hash, err = i.sh.Add(file, shell.Pin(true))
	if err != nil {
		return
	}
	return
}

// UploadDir upload a dictionary to ipfs with cid given back.
func (i *IpfsNode) UploadDir(path string) (hash string, err error) {
	hash, err = i.sh.AddDir(path, shell.Pin(true))
	if err != nil {
		return
	}
	return
}

// UnPinIPFS delete data from ipfs.
func (i *IpfsNode) UnPinIPFS(hash string) (err error) {
	err = i.sh.Unpin(hash)
	if err != nil {
		return
	}

	err = i.sh.Request("repo/gc", hash).
		Option("recursive", true).
		Exec(context.Background(), nil)
	if err != nil {
		return
	}

	return nil
}

// CatIPFS get data from ipfs.
func (i *IpfsNode) CatIPFS(hash string) (string, error) {
	read, err := i.sh.Cat(hash)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)

	return string(body), nil
}

// Status returns whether ipfs daemon is start or not.
func (i *IpfsNode) Status() bool {
	return i.sh.IsUp()
}

// Publish uses IPNS to publish a content with a given ipns url.
func (i *IpfsNode) Publish(name string, cid string) error {
	err := i.sh.Publish(name, cid)
	if err != nil {
		return err
	}
	return nil
}
