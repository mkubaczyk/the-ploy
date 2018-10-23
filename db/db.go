package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mkubaczyk/theploy/config"
	"github.com/mkubaczyk/theploy/models"
)

var DB *gorm.DB
var err error

func Init() {
	dbUser := config.GetEnv("MYSQL_USER", "user")
	dbPass := config.GetEnv("MYSQL_PASSWORD", "pass")
	dbName := config.GetEnv("MYSQL_DATABASE", "db")
	dbHost := config.GetEnv("MYSQL_HOST", "127.0.0.1")
	dbPort := config.GetEnv("MYSQL_PORT", "3307")
	dbConn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)
	DB, err = gorm.Open("mysql", dbConn)
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.DeploymentModel{})
}
