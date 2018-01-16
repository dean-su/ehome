package controllers

import (
	"ehome/models"
	_ "errors"
	"github.com/astaxie/beego"
	"strconv"
	_ "strings"
)

type ProjSampleController struct {
	beego.Controller
}

func (c *ProjSampleController) URLMapping() {
	c.Mapping("GetAllProjSample", c.GetAllProjSample)
	c.Mapping("GetProjSampleNum", c.GetProjSampleNum)
	c.Mapping("GetProjSamplePage", c.GetProjSamplePage)
}

// GetAllProjSample...
// @Title Get all project sample
// @Description get all project sample
// @Success 200 {object}
// @Failure 403 body is empty
// @router /all [get]
func (c *ProjSampleController) GetAllProjSample() {
	var limit int64
	var offset int64

	l, err := models.ProjSampleList(limit, offset)

	map2 := make(map[string]interface{})

	if err != nil {
		map2["status"] = 1
		c.Data["json"] = map2
	} else {
		map2["status"] = 0
		map2["records"] = l
		c.Data["json"] = map2
	}
	c.ServeJSON()
}

//GetProjSampleNum...
// @Title Get order  number
// @Description get all order numbers of a specific user
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /number [get]
func (c *ProjSampleController) GetProjSampleNum() {

	l, err := models.ProjSampleNum()
	map2 := make(map[string]interface{})
	if err != nil {
		map2["status"] = 1
		map2["Number"] = l
		c.Data["json"] = map2
	} else {
		map2["status"] = 0
		map2["Number"] = l
		c.Data["json"] = map2
	}
	c.ServeJSON()
}

// GetProjSamplePage...
// @Title Get project sample page
// @Description get project sample of a page of a specific user
// @Param   Pagenum         Pagenum  string true        "page number"
// @Param   Recperpage      Recperpage string true      "rec per page"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /page [get]
func (c *ProjSampleController) GetProjSamplePage() {

	var limit int64
	var offset int64

	pagenumstr := c.GetString("Pagenum")
	recperpagestr := c.GetString("Recperpage")
	beego.Info("page :", pagenumstr, "|", recperpagestr)
	n, e := strconv.Atoi(pagenumstr)
	m, ee := strconv.Atoi(recperpagestr)

	map2 := make(map[string]interface{})

	if e != nil || ee != nil || m < 1 || n < 1 {
		map2["status"] = 2
		c.Data["json"] = map2

	} else {
		n = n - 1
		limit = int64(m)
		offset = int64(m * n)
		l, err := models.ProjSampleList(limit, offset)

		map2 := make(map[string]interface{})
		if err != nil {
			map2["status"] = 1
			c.Data["json"] = map2
		} else {
			map2["status"] = 0
			map2["records"] = l
			c.Data["json"] = map2
		}
	}
	c.ServeJSON()
}
