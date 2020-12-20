package core

import (
	"fmt"
	"github.com/biningo/etcd-config-gin/global"
	"github.com/biningo/etcd-config-gin/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/**
*@Author lyer
*@Date 2020/12/19 13:59
*@Describe
**/

func RunServer() {
	r := router.InitRouter()
	if err := r.Run(global.G_CONFIG.System.Addr); err != nil {
		global.ErrChan <- fmt.Errorf("%s", err)
	}
}

func ShutDown() {
	err := <-global.ErrChan
	log.Println("error:", err)

	log.Println("close etcd....")
	global.G_ETCD.Close()

	log.Println("shutdown!")
}

func SigListen() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	global.ErrChan <- fmt.Errorf("%s", <-sig)
}
