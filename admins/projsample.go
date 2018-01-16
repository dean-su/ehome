package admins

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type ProjsampleController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProjsampleController) URLMapping() {

	c.Mapping("Num", c.Num)
	c.Mapping("GetProjsample", c.GetProjsample)
	c.Mapping("AddProjsample", c.AddProjsample)
	c.Mapping("ModifyProjsample", c.ModifyProjsample)
	c.Mapping("DeleteProjsample", c.DeleteProjsample)

}

// @param
// @Failure 403 body is empty
// @router /num [post]
func (c *ProjsampleController) Num() {
	map2 := make(map[string]interface{})

	var l int64
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	l, err = models.ProjSampleNum()

	map2["Total"] = l
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /get [post]
func (c *ProjsampleController) GetProjsample() {
	map2 := make(map[string]interface{})
	var m int64

	limit, offset := GetPage(c)
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllProjsample(limit, offset)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllProjsample error!")
		goto BOTTOM
	}

	m, err = models.ProjSampleNum()
	if err != nil {
		SetError(map2, DB_ERROR, "ProjSampleNum error!")
		goto BOTTOM
	}

	map2["Total"] = m

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()

}

// @param
// @Failure 403 body is empty
// @router /add [post]
func (c *ProjsampleController) AddProjsample() {
	var v models.EhomeProjsample

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	v.Thumbimg = c.GetString("Thumbimg")
	if len(v.Thumbimg) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Thumbimg)
		goto BOTTOM
	}

	v.Intro = c.GetString("Intro")
	if len(v.Intro) == 0 {
		SetError(map2, PARAM_ERR, "param Intro error! [%s]", v.Intro)
		goto BOTTOM
	}

	v.City = c.GetString("City")
	if len(v.City) == 0 {
		SetError(map2, PARAM_ERR, "param City error! [%s]", v.City)
		goto BOTTOM
	}

	v.Style = c.GetString("Style")
	if len(v.Style) == 0 {
		SetError(map2, PARAM_ERR, "param Style error! [%s]", v.Style)
		goto BOTTOM
	}

	v.HousingLayout = c.GetString("HousingLayout")
	if len(v.HousingLayout) == 0 {
		SetError(map2, PARAM_ERR, "param HousingLayout error! [%s]", v.HousingLayout)
		goto BOTTOM
	}

	v.HouseSize = c.GetString("HouseSize")
	if len(v.HouseSize) == 0 {
		SetError(map2, PARAM_ERR, "param HouseSize error! [%s]", v.HouseSize)
		goto BOTTOM
	}

	v.Body = c.GetString("Body")
	if len(v.Body) == 0 {
		SetError(map2, PARAM_ERR, "param Body error! [%s]", v.Body)
		goto BOTTOM
	}

	v.Clicks, err = c.GetInt("Clicks")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Clicks"))
		goto BOTTOM
	}

	v.Visible, err = c.GetInt16("Visible")
	if err != nil {
		SetError(map2, PARAM_ERR, "param visible error! [%s]", c.GetString("visible"))
		goto BOTTOM
	}

	map2["Id"], err = models.AddEhomeProjsample(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeProjsample error! %s", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modify [post]
func (c *ProjsampleController) ModifyProjsample() {
	var v models.EhomeProjsample

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	v.Thumbimg = c.GetString("Thumbimg")
	if len(v.Thumbimg) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Thumbimg)
		goto BOTTOM
	}

	v.Intro = c.GetString("Intro")
	if len(v.Intro) == 0 {
		SetError(map2, PARAM_ERR, "param Intro error! [%s]", v.Intro)
		goto BOTTOM
	}

	v.City = c.GetString("City")
	if len(v.City) == 0 {
		SetError(map2, PARAM_ERR, "param City error! [%s]", v.City)
		goto BOTTOM
	}

	v.Style = c.GetString("Style")
	if len(v.Style) == 0 {
		SetError(map2, PARAM_ERR, "param Style error! [%s]", v.Style)
		goto BOTTOM
	}

	v.HousingLayout = c.GetString("HousingLayout")
	if len(v.HousingLayout) == 0 {
		SetError(map2, PARAM_ERR, "param HousingLayout error! [%s]", v.HousingLayout)
		goto BOTTOM
	}

	v.HouseSize = c.GetString("HouseSize")
	if len(v.HouseSize) == 0 {
		SetError(map2, PARAM_ERR, "param HouseSize error! [%s]", v.HouseSize)
		goto BOTTOM
	}

	v.Body = c.GetString("Body")
	if len(v.Body) == 0 {
		SetError(map2, PARAM_ERR, "param Body error! [%s]", v.Body)
		goto BOTTOM
	}

	v.Clicks, err = c.GetInt("Clicks")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Clicks"))
		goto BOTTOM
	}

	v.Visible, err = c.GetInt16("Visible")
	if err != nil {
		SetError(map2, PARAM_ERR, "param visible error! [%s]", c.GetString("visible"))
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.UpdateEhomeProjsampleById(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeProjsampleById error! %s", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /delete [delete]
func (c *ProjsampleController) DeleteProjsample() {
	var Id int
	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.DeleteEhomeProjsample(Id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeExperienc error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAllProjsample(limit, offset int64) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	sortby = append(sortby, "CreateTime")
	order = append(order, "desc")

	ml, err = models.GetAllEhomeProjsample(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) == 0 {
		ml = make([]interface{}, 0)
	}

	return
}
