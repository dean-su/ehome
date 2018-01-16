package controllers

import (
	"ehome/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type StyleController struct {
	beego.Controller
}

// URLMapping ...
func (c *StyleController) URLMapping() {
	c.Mapping("GetStyle", c.GetStyle)
	c.Mapping("GetStylePage", c.GetStylePage)
	c.Mapping("GetStylePrice", c.GetStylePrice)
}

func AllStyle(limit, offset int64) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	fields = append(fields, "Id", "Title", "Introduction", "Sampleimg", "Totalbooking")
	query["Status"] = "1"

	m, err = models.GetAllEhomeStyle(query, fields, sortby, order, offset, limit)

	if err != nil {
		err = fmt.Errorf("GetAllEhomeStyle error! %v", err)
		return
	}

	if len(m) == 0 {
		m = make([]interface{}, 0)
	}
	return
}

// Get ...
// @Title GetStyle
// @Description  GetStyle
// @Success 201 {int}
// @Failure 403 body is empty
// @router / [get]
func (c *StyleController) GetStyle() {

	map2 := make(map[string]interface{})
	m, err := AllStyle(int64(0), int64(0))

	if err != nil {
		SetError(map2, DB_ERROR, "AllStyle terror %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Get ...
// @Title GetStyle
// @Description  GetStylePage
// @Success 201 {int}
// @Failure 403 body is empty
// @router /page [get]
func (c *StyleController) GetStylePage() {
	var limit, offset int64

	var err error
	var ml []interface{}
	map2 := make(map[string]interface{})
	n, e := c.GetInt64("Pagenum")
	m, ee := c.GetInt64("Recperpage")

	if e != nil || ee != nil || m < 1 || n < 1 {
		SetError(map2, PARAM_ERR, "Param Pagenum or Recperpage not valid [%d:%d]", m, n)
		goto BOTTOM
	}

	n = n - 1
	limit = m
	offset = m * n
	ml, err = AllStyle(limit, offset)

	if err != nil {
		SetError(map2, DB_ERROR, "AllStyle error %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = ml

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()

}

// Get ...
// @Title GetStyle
// @Description  GetStyle
// @Success 201 {int}
// @Failure 403 body is empty
// @router /price [get]
func (c *StyleController) GetStylePrice() {
	map2 := make(map[string]interface{})
	id := c.GetString("Id")
	var i int
	var v *models.EhomeStyle
	m, err := StylePrice(id)

	if err != nil {
		SetError(map2, DB_ERROR, "StylePrice error! %v", err)
		goto BOTTOM
	}
	i, err = strconv.Atoi(id)
	if err != nil {
		SetError(map2, PARAM_ERR, "param ID error! %v", err)
		goto BOTTOM
	}
	v, err = models.GetEhomeStyleById(i)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeStyleById error! %v", err)
		goto BOTTOM
	}

	map2["Desc"] = v.Desc
	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func StylePrice(id string) (m []interface{}, err error) {
	var limit, offset int64
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	fields = append(fields, "Id", "Introduction", "Price")
	query["Catid"] = id
	query["Status"] = "1"

	m, err = models.GetAllEhomeStylePrice(query, fields, sortby, order, offset, limit)

	if err != nil {
		err = fmt.Errorf("GetAllEhomeStyle error! %v", err)
		return
	}

	if len(m) == 0 {
		m = make([]interface{}, 0)
	}

	return
}
