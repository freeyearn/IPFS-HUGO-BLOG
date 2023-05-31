package hugo

import (
	"fmt"
	"github.com/gohugoio/hugo/commands"
	"github.com/gohugoio/hugo/common/hugo"
	"os"
	"path/filepath"
)

func Version() {
	fmt.Println(hugo.CurrentVersion)
}

// GetDir command source route, hugo commands will run in dir.
// For example h.dir="resource" means h.dir will be {your project absolute path}\\resource
func GetDir(dir string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir = filepath.Join(wd, dir)
	if err != nil {
		return "", err
	}
	return dir, nil
}

// Create new hugo site, as same as using command `hugo new site xxx`
func Create(site string) (*commands.Response, error) {
	res := commands.Execute([]string{"new", "site", site})
	if res.Err != nil {
		fmt.Println(res.Err)
	}

	return &res, nil
}

// Build hugo project, as same as using command `hugo`
// the target dir in default is set in HugoManager.
func Build(dir string) (*commands.Response, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	err = os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	res := commands.Execute([]string{})
	// move back to project dir
	err = os.Chdir(currentDir)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetPublicDirPath get the public dir path of hugo. The public dir is the dictionary that hugo compiles pages to.
func GetPublicDirPath(dir string) string {
	return filepath.Join(dir, "public")
}
