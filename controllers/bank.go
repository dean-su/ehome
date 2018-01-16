package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

type BankController struct {
	beego.Controller
}

func (c *BankController) URLMapping() {
	c.Mapping("GetBankList", c.GetBankList)
}

// GetAllBank ...
// @Title Get all bank
// @Description get all bank
// @Success 200 {object}
// @Failure 403 body is empty
// @router /list [get]
func (c *BankController) GetBankList() {

	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	map2 := make(map[string]interface{})
	fields = append(fields, "Bank")

	ml, err := models.GetAllEhomeBank(query, fields, sortby, order, offset, limit)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeBank error! %v", err)
		goto BOTTOM
	}

	if len(ml) <= 0 {
		ml = make([]interface{}, 0)
	}
	map2["records"] = ml
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
