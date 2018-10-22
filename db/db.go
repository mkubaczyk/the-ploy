package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mkubaczyk/theploy/models"
)

var DB *gorm.DB
var err error

func Init() {
	DB, err = gorm.Open("sqlite3", "local.sqlite")
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.DeploymentModel{})
}
