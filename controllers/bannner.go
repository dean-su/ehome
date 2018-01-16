package controllers

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type BannerController struct {
	beego.Controller
}

// URLMapping ...
func (c *BannerController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// @param  type   tyep
// @Failure 403 body is empty
// @router / [get]
func (c *BannerController) Get() {
	var err error

	var query = make(map[string]string)
	var sortby []string
	var order []string

	map2 := make(map[string]interface{})
	offset := int64(0)
	limit := int64(0)

	fields := []string{"Img", "Href"}
	t := c.GetString("type")
	if len(t) <= 0 {
		t = "1"
	}
	if t == "1" {
		t = "0"
	} else if t == "2" {
		t = "1"
	} else if t == "3" || t == "4" {
	} else {
		SetError(map2, PARAM_ERR, "PARAM type error! %s", t)
		goto BOTTOM
	}
	query["Type"] = t
	query["status"] = "1"

	map2["records"], err = models.GetAllEhomeBanner(query, fields, sortby, order, offset, limit)

	if err != nil {
		map2["status"] = 1
		delete(map2, "records")
	} else {
		map2["status"] = 0
	}

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
