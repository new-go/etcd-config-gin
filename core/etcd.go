package core

import (
	"log"
	"time"

	"github.com/biningo/etcd-config-gin/global"
	"go.etcd.io/etcd/clientv3"
)

/**
*@Author lyer
*@Date 2020/12/19 13:59
*@Describe
**/

func Etcd() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   global.G_CONFIG.Etcd.Endpoints,
		DialTimeout: time.Duration(global.G_CONFIG.Etcd.Timeout),
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}
