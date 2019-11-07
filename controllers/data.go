package controllers

import (
	"flight/service"
	"github.com/astaxie/beego"
)

type DataController struct {
	beego.Controller
}

func (c *DataController) Get() {
	data := service.GetAll()
	c.Data["json"] = &data
	c.ServeJSON()
}
