package app

import (
	"github.com/nicolasgaraza/petConnect/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

func initDevData(db *gorm.DB) {

	revel.INFO.Println("Update DB Model")
	sr := &models.Service{}
	db.DropTableIfExists(sr)
	db.AutoMigrate(sr)
	revel.INFO.Println("Init DB dev data")

	db.Create(&models.Service{Name: "ServiceTest"})

}
