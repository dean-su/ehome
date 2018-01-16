package controllers

import (
	"ehome/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type PriceController struct {
	beego.Controller
}

// URLMapping ...
func (c *PriceController) URLMapping() {
	c.Mapping("Request", c.Request)
	c.Mapping("Fixtype", c.Fixtype)
}

func GetFixtypeId(fixtype string) int {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id")
	query["Title"] = fixtype

	l, err := models.GetAllEhomeCategory(query, fields, sortby, order, offset, limit)

	if err != nil {
		beego.Error("Can't find fixtype ", fixtype)
		return 0
	}
	if len(l) <= 0 {
		beego.Error("Can't find fixtype ", fixtype)
		return 0
	}

	return (l[0].(map[string]interface{}))["Id"].(int)

}

func GetLabourlist(fixid int) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id", "Desc", "Price")
	query["Catid"] = strconv.Itoa(fixid)

	ml, err = models.GetAllEhomeLabour(query, fields, sortby, order, offset, limit)
	return
}

func GetMateriallist(fixid int) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id", "Desc", "Price", "Unit")
	query["Catid"] = strconv.Itoa(fixid)

	ml, err = models.GetAllEhomeMaterial(query, fields, sortby, order, offset, limit)
	return
}

func GetFixtypePrice(fixtype string) (m map[string]interface{}, err error) {
	fixid := GetFixtypeId(fixtype)
	m = make(map[string]interface{})
	var tmp1, tmp2 []interface{}

	tmp1, err = GetLabourlist(fixid)
	if err != nil {
		return
	}

	tmp2, err = GetMateriallist(fixid)
	if err != nil {
		return
	}
	m["labourprices"] = tmp1

	m["materialprices"] = tmp2

	return
}

// Request ...
// @Title Request
// @Description get HotTopic
// @Param	Mobileno  Mobileno string	false	"Mobileno"
// @Param	Usertype  Usertype  string	false	"Usertype"
// @Param	Token     Token     string	false	"Token"
// @Param   Fixtype   Fixtype   string  false    "Fixtype"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router /request [get]
func (c *PriceController) Request() {

	/*
		mobile := c.GetString("Mobileno")
		token := c.GetString("Token")
	*/
	fixtype := c.GetString("Fixtype")

	//reqtime, err := c.GetInt64("Reqtime")
	var err error
	map2 := make(map[string]interface{})

	var fixtypeslice []string
	var ret map[string]interface{}
	var i int

	/*
		if err != nil {
			map2["status"] = 1
			goto BOTTOM
		}

		_, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			map2["status"] = 2
			goto BOTTOM
		}
	*/

	fixtypeslice = strings.Split(fixtype, ",")
	if len(fixtypeslice) <= 0 {
		map2["status"] = 3
		goto BOTTOM
	}

	for i = 0; i < len(fixtypeslice); i++ {
		ret, err = GetFixtypePrice(fixtypeslice[i])
		if err != nil {
			map2["status"] = 4
			goto BOTTOM
		}
		map2[fixtypeslice[i]] = ret
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router /fixtype [get]
func (c *PriceController) Fixtype() {
	map2 := make(map[string]interface{})

	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id", "Title")

	l, err := models.GetAllEhomeCategory(query, fields, sortby, order, offset, limit)

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeCategory error! %v", err)
		goto BOTTOM
	}

	map2["records"] = l
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
