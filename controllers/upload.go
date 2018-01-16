package controllers

import (
	_ "ehome/models"
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Failure 403 body is empty
// @router / [post]
func (this *UploadController) Post() {
	beego.Info(this.Ctx.Input.Header("Content-Type"))
	//this.Ctx.Input.Method
	beego.Info(string(this.Ctx.Input.RequestBody))
	beego.Info(this.GetString("desc"))
	beego.Info(this.GetString("Filename"))

	f, h, e := this.GetFile("the_file")
	if e == nil {
		f.Close()
	}
	beego.Info("name:", h.Filename)

	err := this.SaveToFile("the_file", "/root/dev/app/src/ehome/static/img/upload.txt")
	if err != nil {
		this.Ctx.WriteString("not ok ")
	} else {

		this.Ctx.WriteString("ok ")
	}

	//this.Ctx.WriteString("hello every body")
}
