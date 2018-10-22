package models

import "github.com/jinzhu/gorm"

type DeploymentModel struct {
	gorm.Model
	Cloud    string
	Provider string
}
