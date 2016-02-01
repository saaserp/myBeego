package routers

import (
	"wencai/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user/profile",&controllers.UserController{},`get:Profile`)
}
