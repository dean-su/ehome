package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

type RegionController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegionController) URLMapping() {
	c.Mapping("GetProvince", c.GetProvince)
	c.Mapping("GetCity", c.GetCity)
	c.Mapping("GetRegion", c.GetRegion)
}

// Get ...
// @Title GetProvince
// @Description  GetProvince
// @Success 201 {int}
// @Failure 403 body is empty
// @router /province [get]
func (c *RegionController) GetProvince() {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Provinceid", "Province")

	m, err := models.GetAllEhomeProvince(query, fields, sortby, order, offset, limit)

	var map2 map[string]interface{}
	map2 = make(map[string]interface{})

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeProvince error %v", err)
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

// Get ...
// @Title  GetCity
// @Description GetCity
// @Success 201 {int}
// @Failure 403 body is empty
// @router /city [get]
func (c *RegionController) GetCity() {

	provinceid := c.GetString("provinceid")

	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Cityid", "City")
	query["Fatherid"] = provinceid

	m, err := models.GetAllEhomeCity(query, fields, sortby, order, offset, limit)

	var map2 map[string]interface{}
	map2 = make(map[string]interface{})

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeCity error %v", err)
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

// Get ...
// @Title GetRegion
// @Description  GetRegion
// @Success 201 {int}
// @Failure 403 body is empty
// @router /region [get]
func (c *RegionController) GetRegion() {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Regionid", "Region")

	cityid := c.GetString("cityid")
	beego.Info(cityid)
	query["Fatherid"] = cityid

	m, err := models.GetAllEhomeArea(query, fields, sortby, order, offset, limit)

	var map2 map[string]interface{}
	map2 = make(map[string]interface{})

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeArea error %v", err)
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
