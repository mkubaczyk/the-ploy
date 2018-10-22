package main

import (
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq"
	"github.com/mkubaczyk/theploy/config"
	"github.com/mkubaczyk/theploy/controllers"
	"github.com/mkubaczyk/theploy/db"
	"time"
)

type Consumer struct {
	name string
	dateCreated time.Time
}

func main() {
	db.Init()
	defer db.DB.Close()
	config.Init()
	config.TaskQueue.StartConsuming(10, time.Second)
	consumer := Consumer{name: "1", dateCreated: time.Now()}
	config.TaskQueue.AddConsumer("1", &consumer)
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