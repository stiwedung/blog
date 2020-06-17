package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/stiwedung/libgo/log"
)

var (
	ROOT       string
	configPath string
	Config     = new(config)
)

func init() {
	execDir, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	ROOT = filepath.Dir(filepath.Dir(execDir))
	configPath = filepath.Join(ROOT, "config/config.toml")
	if fileExist(configPath) {
		if _, err := toml.DecodeFile(configPath, Config); err != nil {
			panic(err)
		}
	} else {
		Config.Common.ServerPort = 8080
	}
}

type dbConfig struct {
	User      string
	Password  string
	MysqlIP   string
	MysqlPort int
	DBName    string
}

type commonConfig struct {
	ServerPort  int
	ReleaseMode bool
	LogCaller   bool
}

type config struct {
	Common commonConfig
	DB     dbConfig
}

func MysqlURL() string {
	if Config.DB.User == "" {
		return ""
	}
	return fmt.Sprintf("%s:%s@(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DB.User,
		Config.DB.Password,
		Config.DB.MysqlIP,
		Config.DB.MysqlPort)
}

func GenConfigFile() {
	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Errorf("create config file failed: %v", err)
		return
	}
	encoder := toml.NewEncoder(file)
	encoder.Encode(Config)
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
