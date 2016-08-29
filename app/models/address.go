package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	AddressLine string
	City        string
	Country     string
	ZipCode     string
}
