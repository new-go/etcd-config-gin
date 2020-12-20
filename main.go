package main

import (
	"github.com/biningo/etcd-config-gin/core"
	"github.com/biningo/etcd-config-gin/global"
)

// https://github.com/flipped-aurora/gin-vue-admin
/**
*@Author lyer
*@Date 2020/12/19 13:45
*@Describe
**/

func main() {
	global.G_VP = core.Viper()
	global.G_ETCD = core.Etcd()
	global.ErrChan = make(chan error)
	go core.RunServer()
	go core.SigListen()
	defer core.ShutDown()
}
