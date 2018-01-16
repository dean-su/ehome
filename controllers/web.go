package controllers

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type WebController struct {
	beego.Controller
}

// URLMapping ...
func (c *WebController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router / [get]
func (c *WebController) Get() {
	id, err := c.GetInt("Id")
	var v *models.EhomeTopic

	err = models.UpdateTopicClick(c.GetString("Id"))

	if err == nil {
		v, err = models.GetEhomeTopicById(id)
	} else {
		beego.Error("UpdateTopicClick error", err)
	}

	if err == nil {
		c.Data["Title"] = v.Title
		c.Data["Desc"] = v.Desc
		c.Data["Content"] = v.Data
	} else {
		beego.Error("GetEhomeTopicById error", id, err)
		c.Data["Title"] = "Error happend"
		c.Data["Desc"] = ""
		c.Data["Content"] = ""
	}
	c.TplName = "projsample.tpl"
	c.Render()
}
