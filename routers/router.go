package routers

import (
	"github.com/astaxie/beego"
	"i/flight/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
