package main

import (
	"Genitive/cmd/api/app"
	"Genitive/config"
	"Genitive/services"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	ConfDir string
)

func init() {
	flag.StringVar(&ConfDir, "f", GetProjectRoot()+"/conf", "set config file directory")
}

func getBinAbPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 定义项目根目录
func GetProjectRoot() string {
	dir := getBinAbPath()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		dir = getCallerAbPath()
	}
	return path.Dir(dir)
}

func getCallerAbPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func main() {
	// 解析配置文件
	config.InitConfig(ConfDir + "/config.yaml")
	go services.Runbevm()
	app.Run()
}
