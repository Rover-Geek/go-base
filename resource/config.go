package resource

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

/**
 * @description:
 * @author: Rover
 * @date: 2020/8/30 1:59 下午
 */

var env string

func init()  {
	log.Infof(">>>>>>>>>>开始初始化配置.......")
	env = GetEnv()
	cfgName := "app" + env
	cfgPath := os.Getenv("GOPATH") + "/src/go-base/resource"

	viper.SetConfigName(cfgName)
	viper.AddConfigPath(cfgPath)

	viper.ReadInConfig()

	log.Infof("配置加载成功, env: %s", env)

}

func GetEnv() string {
	env = os.Getenv("env")
	if env == "" {
		return ""
	}
	return env
}

func GetCfg(key string) interface{} {
	port := viper.Get(key)
	if port == nil {
		return 8082
	}
	return port
}