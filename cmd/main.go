package main

import (
	"golang-rest-api/config"
	"golang-rest-api/config/mongodb"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/routes"
	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// func welcome(c *gin.Context) {
// 	mongo, err := mongodb.Connect()
// 	helper.PanicIfError(err)

// 	collection := mongo.Database("development").Collection("users")

// 	filter := bson.M{"name": "aku"}
// 	var response bson.M
// 	err = collection.FindOne(c, filter).Decode(&response)
// 	if err != nil {
// 		helper.ResponseAPI(c, 200, err)
// 		return
// 	}

// 	result := gin.H{
// 		"status":  true,
// 		"message": "Welcome API REST GO!",
// 		"result":  response,
// 	}

// 	helper.ResponseAPI(c, 200, result)
// }

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	mongodb.Connect()

	router.ForwardedByClientIP = true
	router.HandleMethodNotAllowed = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(config.CorsHandler())
	router.Use(middleware.LoggerMiddleware())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(helper.Panic)
	router.Static("/static", "web")
	routes.IndexRoutes(router)
	router.NoMethod(helper.MethodNotAllowed)
	router.NoRoute(helper.NotFound)

	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
