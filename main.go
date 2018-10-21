package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

var db *gorm.DB
var err error

type DeploymentParams struct {
	Cloud    string `json:"cloud"`
	Provider string `json:"provider"`
}

type DeploymentModel struct {
	gorm.Model
	Cloud    string
	Provider string
}

type DeploymentResult struct {
	Id       int
	Cloud    string
	Provider string
}

func main() {
	db, err = gorm.Open("sqlite3", "local.sqlite")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&DeploymentModel{})
	defer db.Close()
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/deployments/:id", getDeploymentEndpoint)
		api.POST("/deployments", createDeploymentEndpoint)
	}
	r.Run()
}

func createDeploymentEndpoint(c *gin.Context) {
	var deploymentData DeploymentParams
	c.BindJSON(&deploymentData)
	create := db.Create(&DeploymentModel{Cloud: deploymentData.Cloud, Provider: deploymentData.Provider})
	id := create.Value.(*DeploymentModel).ID
	var deployment DeploymentModel
	var queryResult DeploymentResult
	db.First(&deployment, id).Scan(&queryResult)
	c.JSON(http.StatusOK, queryResult)
}

func getDeploymentEndpoint(c *gin.Context) {
	id := c.Param("id")
	var deployment DeploymentModel
	var queryResult DeploymentResult
	resp := db.First(&deployment, id)
	if resp.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": resp.Error.Error()})
		c.Abort()
		return
	}
	resp.Scan(&queryResult)
	c.JSON(200, queryResult)
}
