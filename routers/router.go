package routers

import (
	"github.com/astaxie/beego"
	"i/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
