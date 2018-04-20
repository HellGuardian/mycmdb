package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct  {
	beego.Controller
}

func (mc *MainController) Get() {
	// 定义首页模板
	//mc.Data["IsIndex}"] = true
	mc.Data["Website"] = "tanxiaowei@jd.com"
	mc.Data["Email"] = "astaxie@gmail.com"
	mc.TplName = "index.html"

}
