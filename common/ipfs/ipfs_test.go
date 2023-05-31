package ipfs

import (
	"os"
	"testing"
)

var ipfs *IpfsNode

func setup() {
	ipfs = GetIpfs("localhost:5001")
}

func TestIpfs_Status(t *testing.T) {
	if !ipfs.Status() {
		t.Error("ipfs not up")
	} else {
		t.Log("ipfs up")
	}
}

func TestIpfs_UploadStr(t *testing.T) {
	hash, err := ipfs.UploadStr("hello")
	if err != nil {
		t.Error(err)
	}
	t.Log("CID:", hash)
	if hash != "QmWfVY9y3xjsixTgbd9AorQxH7VtMpzfx2HaWtsoUYecaX" {
		t.Error("ipfs upload error: hash not match")
	}
}

func TestIpfs_Publish(t *testing.T) {
	err := ipfs.Publish("", "/ipfs/QmY15H1YGyMwaoT6DBBFBEXy6EtJgUDVEZ1T2wZxSEmr8m")
	if err != nil {
		t.Error(err)
	}
}

func TestMain(m *testing.M) {

	setup()
	code := m.Run()

	os.Exit(code)
}
