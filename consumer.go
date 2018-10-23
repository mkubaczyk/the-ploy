package main

import (
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq"
	"github.com/mkubaczyk/theploy/config"
	"github.com/mkubaczyk/theploy/controllers"
	"github.com/mkubaczyk/theploy/db"
	"os"
	"time"
)

type Consumer struct {
}

func main() {
	db.Init()
	defer db.DB.Close()
	config.Init()
	config.TaskQueue.StartConsuming(10, time.Second)
	hostname, _ := os.Hostname()
	config.Logger.Info(fmt.Sprintf("Initializing %v consumer", hostname))
	config.TaskQueue.AddConsumer(hostname, &Consumer{})
	select {}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	var task controllers.DeploymentTask
	err := json.Unmarshal([]byte(delivery.Payload()), &task)
	if err != nil {
		config.Logger.Warning(fmt.Sprintf("rejected %v", task.Id))
		delivery.Reject()
		return
	}
	config.Logger.Info(fmt.Sprintf("performing task %v", task.Id))
	delivery.Ack()
}