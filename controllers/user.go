package controllers

import (
	"github.com/astaxie/beego"
)
type UserController struct{
	beego.Controller
}
func (c *UserController) Profile(){
	c.Data["userid"]="wencai"
	c.Data["tag"]="I am wencai"
	c.Data["hobby"]=[] string{"run","code"}
	c.TplName = "user/profile.html"
}