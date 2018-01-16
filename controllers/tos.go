package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

type TosController struct {
	beego.Controller
}

// URLMapping ...
func (c *TosController) URLMapping() {
	c.Mapping("GetTos", c.GetTos)
}

// Get ...
// @Title GetTos
// @Description  GetTos
// @Success 201 {int}
// @Failure 403 body is empty
// @router /all [get]
func (c *TosController) GetTos() {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Id", "Title")

	m, err := models.GetAllEhomeTos(query, fields, sortby, order, offset, limit)

	var map2 map[string]interface{}
	map2 = make(map[string]interface{})

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeTos error %v", err)
		goto BOTTOM
	}
	if len(m) == 0 {
		m = make([]interface{}, 0)
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
