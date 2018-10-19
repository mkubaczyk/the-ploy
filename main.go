package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Deployment struct {
	Id       string `json:"id"`
	Cloud    string `json:"cloud"`
	Provider string `json:"provider"`
}

type DeploymentModel struct {
	gorm.Model
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
		api.GET("/deployments/:id", deploymentEndpoint)
		api.POST("/deployments", createDeploymentEndpoint)
	}
	r.Run()
}

type QueryResult struct {
	Cloud    string
	Provider string
}

func createDeploymentEndpoint(c *gin.Context) {
	var deploymentData Deployment
	c.BindJSON(&deploymentData)
	db.Create(&DeploymentModel{Cloud: deploymentData.Cloud, Provider: deploymentData.Provider})
	var deployment DeploymentModel
	resp := db.First(&deployment, "cloud = ?", "AWS")
	var queryResult QueryResult
	db.Table("deployment_models").Select("cloud, provider").Where("provider = ?", "EB").Scan(&queryResult)
	a, _ := json.Marshal(queryResult)
	fmt.Println(string(a))
	c.JSON(200, gin.H{"ID": resp.Value})
}

func deploymentEndpoint(c *gin.Context) {
	id := c.Param("id")
	resp := Deployment{
		Id:       id,
		Cloud:    "AWS",
		Provider: "EB",
	}
	c.JSON(200, resp)
}
