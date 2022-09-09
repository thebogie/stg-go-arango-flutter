package main

import (
	config "back/configs"
	route "back/routes"
	util "back/utils"
	"log"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	/**
	@description Setup Server
	*/

	router := SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + util.GodotEnv("GO_PORT")))
}

func SetupRouter() *gin.Engine {

	dbc := config.DBconnection{}
	dbc.CreateConnection()

	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/

	log.Printf("GO_ENV:%v", util.GodotEnv("GO_ENV"))

	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://smacktalkgaming.com", "http://localhost:50000", "http://127.0.0.1:50000"},
		AllowMethods: []string{"PUT", "POST", "GET", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		//AllowAllOrigins: true,
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	@description Init All Route
	*/
	route.InitAuthRoutes(&dbc, router)
	/*
		route.InitStudentRoutes(dbc, router)
	*/
	return router
}
