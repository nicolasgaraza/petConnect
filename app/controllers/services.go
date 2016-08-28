package controllers

import (
	"github.com/nicolasgaraza/petConnect/app/models"

	"github.com/nicolasgaraza/petConnect/app"

	"github.com/revel/revel"
)

//Constructor
type Services struct {
	*revel.Controller
}

func (c Services) Get() revel.Result {

	var service models.Service

	app.Gdb.First(&service, 1)

	return c.RenderJson(service)
}

func (c Services) GetById() revel.Result {

	var id int
	c.Params.Bind(&id, "id")

	revel.INFO.Println("Value of :id", id)

	var service models.Service

	app.Gdb.First(&service, id)

	return c.RenderJson(service)
}
