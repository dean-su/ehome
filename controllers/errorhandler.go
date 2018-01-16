package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	map2 := make(map[string]interface{})
	map2["errmsg"] = "page not found"
	map2["status"] = 404

	c.Ctx.Output.Status = 200
	c.Data["json"] = map2
	c.ServeJSON()
}

func (c *ErrorController) Error501() {
	map2 := make(map[string]interface{})
	map2["errmsg"] = "server error"
	map2["status"] = 501

	c.Ctx.Output.Status = 200

	c.Data["json"] = map2
	c.ServeJSON()
}

func (c *ErrorController) ErrorDb() {
	map2 := make(map[string]interface{})
	map2["errmsg"] = "database is now down"
	map2["status"] = 1

	c.Data["json"] = map2
	c.ServeJSON()
}
