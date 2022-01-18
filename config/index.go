package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Mongo struct {
	Uri string `json:"uri"`
}

type Mysql struct {
	Uri string `json:"uri"`
}

type Config struct {
	Env   string `json:"env"`
	Mongo Mongo  `json:"mongo"`
	Mysql Mysql  `json:"mysql"`
}

func init() {
}

var Conf Config

func Init() (err error) {
	log.Println("init config with viper ...")

	v := viper.New()
	configType := "json"

	v.AddConfigPath("./config")
	v.SetConfigName("default")
	v.SetConfigType(configType)

	fmt.Println("read config [default] ...")
	if err = v.ReadInConfig(); err != nil {
		fmt.Println("get default.config:", err)
		return err
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, value := range configs {
		v.SetDefault(k, value)
	}
	env := os.Getenv("GO_ENV")
	// 根据配置的env读取相应的配置信息

	if env != "" {
		fmt.Printf("read config [%s] ...\n", env)
		v.SetConfigName(env)
		err = v.ReadInConfig()
		if err != nil {
			return err
		}
	}

	if err = v.Unmarshal(&Conf); err != nil {
		fmt.Println("config Unmarshal err:", err)
		return err
	}
	log.Println("init config done:", Conf)
	return nil
}
