package models

import "github.com/jinzhu/gorm"

type ServiceInstance struct {
	gorm.Model
	Price       float32
	Description string
	Likes       int
	Dislikes    int
}
