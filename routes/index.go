package routes

import (
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		helper.ResponseSuccess(c, 200, "Welcome REST API Go!")
	})

	AuthRoutes(router.Group("/o"))
	router.MaxMultipartMemory = 1 << 20 //1mb
	UploadRoutes(router.Group("/upload"))
}

func AuthRoutes(router *gin.RouterGroup) {
	v1 := router.Group("/auth/v1")
	{
		v1.POST("/login", controller.Login)
	}
}

func UploadRoutes(router *gin.RouterGroup) {
	v1 := router.Group("/image/v1")
	{
		v1.POST("/single", middleware.UploadImage)
	}
}
