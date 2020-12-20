package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"github.com/biningo/etcd-config-gin/global"
)

/**
*@Author lyer
*@Date 2020/12/19 13:59
*@Describe
**/

func Viper() *viper.Viper {
	path := ""
	flag.StringVar(&path, "c", "", "choose config file")
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.G_CONFIG); err != nil {
			log.Fatal(err)
		}
	})

	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		log.Fatal(err)
	}
	return v
}
