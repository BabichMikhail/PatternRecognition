package routers

import (
	"github.com/BabichMikhail/PatternRecognition/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Lab8Controller{}, "get,post:Main")
	beego.Router("/lab2", &controllers.MainController{}, "get,post:Main")
	beego.Router("/lab3", &controllers.Lab3Controller{}, "get,post:Main")
	beego.Router("/lab8", &controllers.Lab8Controller{}, "get,post:Main")
}
