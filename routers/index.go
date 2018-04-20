package routers

import (
	"cmdb/controllers"
	"github.com/astaxie/beego"
)

func init () {
	// 注册路由
	beego.Router("/", &controllers.MainController{})
}