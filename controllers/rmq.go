package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mkubaczyk/theploy/config"
	"net/http"
)

func GetRMQStats(c *gin.Context) {
	layout := c.Param("layout")
	refresh := c.Param("refresh")
	queues := config.RedisConn.GetOpenQueues()
	stats := config.RedisConn.CollectStats(queues)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(stats.GetHtml(layout, refresh)))
}
