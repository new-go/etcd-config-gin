package router

import "github.com/gin-gonic/gin"

/**
*@Author lyer
*@Date 2020/12/19 14:38
*@Describe
**/

func InitRouter() *gin.Engine {
	r := gin.Default()

	PublicRouter := r.Group("")
	{
		InitPublicRouter(PublicRouter)
	}

	EtcdSvcCfgRouter := r.Group("")
	{
		InitEtcdCfgSvcRouter(EtcdSvcCfgRouter)
	}
	return r
}
