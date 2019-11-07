package routers

import (
	"flight/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/data", &controllers.DataController{})
}
