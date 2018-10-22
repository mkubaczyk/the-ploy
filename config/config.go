package config

import (
	"github.com/adjust/rmq"
	logger "github.com/apsdehal/go-logger"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var RedisConn rmq.Connection
var TaskQueue rmq.Queue
var Logger *logger.Logger

func Init() {
	Logger, _ = logger.New("stdout", 1, os.Stdout)
	Logger.Info("Initializing config")
	redisAddr := getEnv("REDIS_ADDR", "127.0.0.1:6379")
	RedisConn = rmq.OpenConnection("Queue", "tcp", redisAddr, 0)
	TaskQueue = RedisConn.OpenQueue("tasks")
	Logger.Info("Config initialized")
}
