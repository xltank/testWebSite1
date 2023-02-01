package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Mongo struct {
	Uri string `json:"uri"`
}

type Mysql struct {
	Uri string `json:"uri"`
}

type Config struct {
	Env        string `json:"env,omitempty"`
	Port       string `json:"port,omitempty"`
	Mongo      Mongo  `json:"mongo,omitempty"`
	Mysql      Mysql  `json:"mysql,omitempty"`
	DefaultVar string `json:"default_var,omitempty"`
}

func init() {
}

var Conf Config

func Init() (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("init config with viper ...")

	v := viper.New()
	configType := "json"

	v.AddConfigPath("./config")
	v.SetConfigName("default")
	v.SetConfigType(configType)

	// log.Println("read config [default] ...")
	if err = v.ReadInConfig(); err != nil {
		log.Println("get default.config:", err)
		return err
	}

	configs := v.AllSettings()
	// 将default中的配置全部以默认配置写入，否则无法达到local覆盖default的目的。
	for k, value := range configs {
		v.SetDefault(k, value)
	}

	// log.Println("test env var:", os.Getenv("var1"))
	env := os.Getenv("WEBSITE_ENV")
	log.Println("env:", env)
	// 根据配置的env读取相应的配置信息

	if env != "" {
		// log.Printf("read config [%s] ...\n", env)
		v.SetConfigName(env)
		err = v.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// log.Println("default var:", v.GetString("defaultVar"))

	if err = v.Unmarshal(&Conf); err != nil {
		log.Println("config Unmarshal err:", err)
		return err
	}
	log.Println("init config done, defaultVar:", Conf.DefaultVar)
	return nil
}
