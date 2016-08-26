package controllers

import "github.com/revel/revel"

//Constructor
type App struct {
	*revel.Controller
}

//Index
func (c App) Index() revel.Result {
	return c.Render()
}
