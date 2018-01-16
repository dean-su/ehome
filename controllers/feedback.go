package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

type FeedbackController struct {
	beego.Controller
}

// URLMapping ...
func (c *FeedbackController) URLMapping() {
	c.Mapping("Feedback", c.Feedback)
}

// POST ...
// @Title Feedback
// @Description  Feedback
// @Success 201 {int}
// @Failure 403 body is empty
// @router / [post]
func (c *FeedbackController) Feedback() {
	map2 := make(map[string]interface{})
	var v models.EhomeFeedback

	v.Feedback = c.GetString("Feedback")
	v.Mobileno = c.GetString("Mobileno")

	_, err := models.AddEhomeFeedback(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeFeeback error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
