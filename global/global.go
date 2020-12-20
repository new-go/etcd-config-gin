package global

/**
*@Author lyer
*@Date 2020/12/19 13:47
*@Describe
**/
import (
	"github.com/biningo/etcd-config-gin/config"
	"go.etcd.io/etcd/clientv3"
	"github.com/spf13/viper"
	"log"
)

var (
	G_ETCD   *clientv3.Client
	G_VP     *viper.Viper
	G_LOG    log.Logger
	G_CONFIG config.Server
	ErrChan chan error
)
