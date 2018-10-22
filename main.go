package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mkubaczyk/theploy/config"
	"github.com/mkubaczyk/theploy/controllers"
	"github.com/mkubaczyk/theploy/db"
)

func main() {
	db.Init()
	defer db.DB.Close()
	config.Init()
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/deployments/:id", controllers.GetDeploymentEndpoint)
		api.POST("/deployments", controllers.CreateDeploymentEndpoint)
		api.GET("/rmq", controllers.GetRMQStats)
	}
	r.Run()
}
