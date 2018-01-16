package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RedirectController struct {
	beego.Controller
}

// URLMapping ...
func (c *RedirectController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

// @param   num   num false
// @Failure 403 body is empty
// @router / [get]
func (c *RedirectController) Get() {

	beego.Error("redirect Get Data")
	beego.Error(c.Input())
	fmt.Printf("data [%v]\n", c.Ctx.Input.Data)

	beego.Error("redirect Get Body")
	beego.Error(string(c.Ctx.Input.RequestBody))
}

// @param   num   num false
// @Failure 403 body is empty
// @router / [post]
func (c *RedirectController) Post() {

	beego.Error("redirect post Data")
	beego.Error(c.Input())

	beego.Error("redirect post Body")
	beego.Error(string(c.Ctx.Input.RequestBody))
}
