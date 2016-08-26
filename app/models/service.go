package models

import "github.com/jinzhu/gorm"

type Service struct { // example user fields
	gorm.Model
	Name string
}
