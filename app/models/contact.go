package models

import "github.com/jinzhu/gorm"

type Contact struct {
	gorm.Model
	Email       string
	PhoneNumber string
}
