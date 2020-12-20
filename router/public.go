package router

import "github.com/gin-gonic/gin"

/**
*@Author lyer
*@Date 2020/12/19 14:52
*@Describe
**/

func InitPublicRouter(router *gin.RouterGroup) {
	PublicRouter := router.Group("/public")
	{
		PublicRouter.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "OK")
		})
	}
}
