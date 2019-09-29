package controllers

import (
	"github.com/astaxie/beego"
	"i/service"
)

type DataController struct {
	beego.Controller
}

func (c *DataController) Get() {
	data := service.GetAll()
	c.Data["json"] = &data
	c.ServeJSON()
}
