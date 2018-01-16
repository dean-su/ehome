package controllers

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type ExperienceController struct {
	beego.Controller
}

// URLMapping ...
func (c *ExperienceController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @param   num   num false
// @Failure 403 body is empty
// @router / [get]
func (c *ExperienceController) Get() {
	var query = make(map[string]string)
	var sortby []string
	var order []string

	offset := int64(0)
	limit := int64(0)

	fields := []string{"Title", "Image", "Body"}
	limit, _ = c.GetInt64("num")
	//query["visible"] = "1"

	l, err := models.GetAllEhomeExperience(query, fields, sortby, order, offset, limit)

	map2 := make(map[string]interface{})
	if err != nil {
		map2["status"] = 1
	} else {
		if l == nil {
			l = make([]interface{}, 0)
		}
		map2["status"] = 0
		map2["records"] = l
	}

	c.Data["json"] = map2
	c.ServeJSON()
}
