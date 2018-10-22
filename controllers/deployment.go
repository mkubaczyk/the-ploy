package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mkubaczyk/theploy/config"
	"github.com/mkubaczyk/theploy/db"
	"github.com/mkubaczyk/theploy/models"
	"net/http"
)

type DeploymentParams struct {
	Cloud    string `json:"cloud"`
	Provider string `json:"provider"`
}

type DeploymentResult struct {
	Id       int
	Cloud    string
	Provider string
}

type DeploymentTask struct {
	Id int
}

func CreateDeploymentEndpoint(c *gin.Context) {
	var deploymentData DeploymentParams
	c.BindJSON(&deploymentData)
	create := db.DB.Create(&models.DeploymentModel{Cloud: deploymentData.Cloud, Provider: deploymentData.Provider})
	id := create.Value.(*models.DeploymentModel).ID
	var deployment models.DeploymentModel
	var queryResult DeploymentResult
	db.DB.First(&deployment, id).Scan(&queryResult)
	task := DeploymentTask{queryResult.Id}
	b, _ := json.Marshal(task)
	config.TaskQueue.Publish(string(b))
	c.JSON(http.StatusOK, queryResult)
}

func GetDeploymentEndpoint(c *gin.Context) {
	id := c.Param("id")
	var deployment models.DeploymentModel
	var queryResult DeploymentResult
	resp := db.DB.First(&deployment, id)
	if resp.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": resp.Error.Error()})
		c.Abort()
		return
	}
	resp.Scan(&queryResult)
	c.JSON(200, queryResult)
}
