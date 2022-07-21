package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"
)

type MysqlConfig struct {
	User     string `json:"user"`
	DB       string `json:"db"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Db       int    `json:"db"`
	Password string `json:"password"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type AppConfiguration struct {
	ENV        string
	MysqlConf  MysqlConfig
	RedisConf  RedisConfig
	ServerConf ServerConfig
}

var AppConf AppConfiguration

func Init() {
	basePath, concurrentOsTag := GetBasePath()
	env := os.Getenv("ENV")
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		return
	}
	if env == "prod" {
		viper.SetConfigName("config.prod")
	} else {
		viper.SetConfigName("config.dev")
	}
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logrus.Infof("path = %v", path)
	idx := strings.Index(path, fmt.Sprintf("%scourseSystem", concurrentOsTag))
	if idx < 0 {
		panic(fmt.Errorf("project path /courseSystem idx = %v", idx))
	}
	viper.AddConfigPath(basePath + fmt.Sprintf("%sconfig", concurrentOsTag))
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config, %s", err))
	}
	err = viper.Unmarshal(&AppConf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into appConf, %v", err))
	}
	logrus.Infof("AppConf = %v", AppConf)
	AppConf.ENV = env
	logrus.Infof("init courseSystem success: env = %v, conf = %v", AppConf.ENV, AppConf)
}

func GetBasePath() (path string, pathTag string) {
	var concurrentOsTag string
	if strings.Compare("windows", runtime.GOOS) == 0 {
		concurrentOsTag = "\\"
	} else {
		concurrentOsTag = "/"
	}
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	logrus.Infof("path = %v", path)
	idx := strings.Index(path, fmt.Sprintf("%scourseSystem", concurrentOsTag))
	if idx < 0 {
		panic(fmt.Errorf("project path /courseSystem idx = %v", idx))
	}
	basePath := path[0:idx] + fmt.Sprintf("%scourseSystem", concurrentOsTag)
	return basePath, concurrentOsTag
}
