package router

import (
	v1 "github.com/biningo/etcd-config-gin/api/v1"
	"github.com/gin-gonic/gin"
)

/**
*@Author lyer
*@Date 2020/12/19 14:48
*@Describe
**/
func InitEtcdCfgSvcRouter(router *gin.RouterGroup) {
	EtcdCfgRouter := router.Group("/config")
	{
		EtcdCfgRouter.GET("/get", v1.GetConfig)
		EtcdCfgRouter.GET("/list", v1.ListConfig)
		EtcdCfgRouter.DELETE("/del", v1.DelConfig)
		EtcdCfgRouter.POST("/post", v1.PutConfig)
		EtcdCfgRouter.POST("/upload", v1.UploadConfig)
		EtcdCfgRouter.GET("/download", v1.DownloadConfig)
	}
}
