package controllers

import (
	"ehome/models"
	"github.com/astaxie/beego"
	"strconv"
)

type FixAddressController struct {
	beego.Controller
}

// URLMapping ...
func (c *FixAddressController) URLMapping() {
	c.Mapping("Request", c.Request)
	c.Mapping("Add", c.Add)
	c.Mapping("Modify", c.Modify)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Default", c.Default)
}

func GetUserAddressList(userid int) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Id", "Contactname", "Contactaddr", "Phone", "IsDefaultaddr", "Region", "Cityid")
	query["Userid"] = strconv.Itoa(userid)

	m, err = models.GetAllEhomeFixAddress(query, fields, sortby, order, offset, limit)

	if err != nil {
		beego.Error("GetAllEhomeFixAddress error!", err)
		return
	}

	for i := 0; i < len(m); i++ {
		tmp := (m[i].(map[string]interface{}))["Region"].(string)
		province, city, region, e := models.GetRegionDetailById(tmp)
		if e != nil {
			beego.Error("GetRegionDetailById error! %s", tmp)
			err = e
			return
		}
		(m[i].(map[string]interface{}))["Region"] = province + city + region
		(m[i].(map[string]interface{}))["Regionid"] = tmp
		tm, e := models.GetCityById((m[i].(map[string]interface{}))["Cityid"].(string))
		if e != nil {
			beego.Error("GetCityById error! %v %v", (m[i].(map[string]interface{}))["Cityid"], e)
			return
		}
		(m[i].(map[string]interface{}))["Provinceid"] = tm[0].(models.EhomeCity).Fatherid
	}

	if len(m) > 0 {
		return
	}

	m = make([]interface{}, 0)

	return

}

// Request ...
// @Title Request
// @Description get HotTopic
// @Param	Mobileno  Mobileno string	false	"Mobileno"
// @Param	Usertype  Usertype  string	false	"Usertype"
// @Param	Token     Token     string	false	"Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Success 200 {object}
// @Failure 403
// @router /request [get]
func (c *FixAddressController) Request() {

	map2 := make(map[string]interface{})

	var err error
	var m []interface{}
	var userid int
	Admin := c.GetString("Name")

	if Admin == "" {
		mobile, token, reqtime, _, err := GetComUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "user %s is not valid! [%v]", mobile, err)
			goto BOTTOM
		}
	} else {
		mobile := c.GetString("Mobileno")
		name, token, reqtime, err := GetAdminUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		err = models.CheckAdminUser(name, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "admin user %s is invalid", name)
			goto BOTTOM
		}

		userid, err = models.GetUseridByNo(mobile)
		if err != nil {
			SetError(map2, INVALID_USER, "mobileno [%s] is invalid!", mobile)
			goto BOTTOM
		}

	}
	m, err = GetUserAddressList(userid)

	if err != nil {
		SetError(map2, DB_ERROR, "GetUserAddressList error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Add ...
// @Title  Add
// @Description Add address
// @Param	Mobileno  Mobileno string	false	"Mobileno"
// @Param	Usertype  Usertype  string	false	"Usertype"
// @Param	Token     Token     string	false	"Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Param   Phone     Phone     string  true     "Phone"
// @Param   Contactname    Contactname   string  true    "Contactname"
// @Param   Contactaddr    Contactaddr  string  true     "Contactaddr"
// @Param   IsDefaultaddr  IsDefaultaddr string  true     "IsDefaultaddr"
// @Success 200 {object}
// @Failure 403
// @router /add [get]
func (c *FixAddressController) Add() {
	map2 := make(map[string]interface{})

	var m []interface{}
	var userid int
	var v models.EhomeFixAddress
	var area []interface{}

	var id int64

	Admin := c.GetString("Name")
	var err error

	if Admin == "" {
		mobile, token, reqtime, _, err := GetComUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "user %s is not valid! [%v]", mobile, err)
			goto BOTTOM
		}
	} else {
		mobile := c.GetString("Mobileno")
		name, token, reqtime, err := GetAdminUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		err = models.CheckAdminUser(name, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "admin user %s is invalid", name)
			goto BOTTOM
		}
		userid, err = models.GetUseridByNo(mobile)
		if err != nil {
			SetError(map2, INVALID_USER, "mobileno [%s] is invalid!", mobile)
			goto BOTTOM
		}

	}

	v.Userid = userid
	v.Contactname = c.GetString("Contactname")
	if len(v.Contactname) == 0 {
		SetError(map2, PARAM_ERR, "param Contactname error!len is 0")
		goto BOTTOM
	}
	v.Phone = c.GetString("Phone")
	if len(v.Phone) == 0 {
		SetError(map2, PARAM_ERR, "param Phone error!len is 0")
		goto BOTTOM
	}

	v.Contactaddr = c.GetString("Contactaddr")
	if len(v.Contactaddr) == 0 {
		SetError(map2, PARAM_ERR, "param Contactaddr error!len is 0")
		goto BOTTOM
	}

	v.Region = c.GetString("Region")
	area, err = models.GetRegionById(v.Region)
	if err != nil {
		SetError(map2, PARAM_ERR, "param Region error! %v", err)
		goto BOTTOM
	}
	v.Cityid = strconv.Itoa(area[0].(models.EhomeArea).Fatherid)

	if len(v.Contactaddr) == 0 {
		SetError(map2, PARAM_ERR, "param  Contactaddr error! len is 0!")
		goto BOTTOM
	}

	v.IsDefaultaddr, err = c.GetInt8("IsDefaultaddr")
	if err != nil {
		SetError(map2, PARAM_ERR, "param IsDefaultaddr error! %s", c.GetString("IsDefaultaddr"))
		goto BOTTOM
	}

	id, err = models.AddEhomeFixAddress(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeFixAddress error! %v", err)
		goto BOTTOM
	}

	err = models.UpdateDefaultAddress(id, userid, v.IsDefaultaddr)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateDefaultAddress error! %v", err)
		goto BOTTOM
	}

	m, err = GetUserAddressList(userid)

	if err != nil {
		SetError(map2, DB_ERROR, "GetUserAddressList error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Modify ...
// @Title Modify
// @Description modify address
// @Param	Mobileno  Mobileno string	false	"Mobileno"
// @Param	Usertype  Usertype  string	false	"Usertype"
// @Param	Token     Token     string	false	"Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Param   Phone     Phone     string  true     "Phone"
// @Param   Contactaddr    Contactaddr  string  true     "Contactaddr"
// @Param   IsDefaultaddr  IsDefaultaddr string  true     "IsDefaultaddr"
// @Success 200 {object}
// @Failure 403
// @router /modify [get]
func (c *FixAddressController) Modify() {

	map2 := make(map[string]interface{})

	var m []interface{}
	var userid int
	var v models.EhomeFixAddress
	var area []interface{}
	var err error

	Admin := c.GetString("Name")

	if Admin == "" {
		mobile, token, reqtime, _, err := GetComUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "user %s is not valid! [%v]", mobile, err)
			goto BOTTOM
		}
	} else {
		mobile := c.GetString("Mobileno")
		name, token, reqtime, err := GetAdminUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		err = models.CheckAdminUser(name, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "admin user %s is invalid", name)
			goto BOTTOM
		}
		userid, err = models.GetUseridByNo(mobile)
		if err != nil {
			SetError(map2, INVALID_USER, "mobileno [%s] is invalid!", mobile)
			goto BOTTOM
		}
	}
	v.Userid = userid
	v.Contactname = c.GetString("Contactname")
	if len(v.Contactname) == 0 {
		SetError(map2, PARAM_ERR, "param Contactname invalid! len is 0")
		goto BOTTOM
	}

	v.Phone = c.GetString("Phone")
	if len(v.Phone) == 0 {
		SetError(map2, PARAM_ERR, "param Contactname invalid! len is 0")
		goto BOTTOM
	}

	v.Contactaddr = c.GetString("Contactaddr")
	if len(v.Contactaddr) == 0 {
		SetError(map2, PARAM_ERR, "param  Contactaddr invalid! len is 0")
		goto BOTTOM
	}

	v.Region = c.GetString("Region")
	area, err = models.GetRegionById(v.Region)
	if err != nil {
		SetError(map2, PARAM_ERR, "param Region error! %v", err)
		goto BOTTOM
	}
	v.Cityid = strconv.Itoa(area[0].(models.EhomeArea).Fatherid)

	if len(v.Contactaddr) == 0 {
		SetError(map2, PARAM_ERR, "param  Contactaddr error! len is 0!")
		goto BOTTOM
	}

	v.IsDefaultaddr, err = c.GetInt8("IsDefaultaddr")
	if err != nil {
		SetError(map2, PARAM_ERR, "param  IsDefaultaddr error! len is 0!")
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param Id is invalid! %s", c.GetString("Id"))
		goto BOTTOM
	}

	if !models.AddressIdBelongToUser(userid, v.Id) {
		SetError(map2, PARAM_ERR, "user %d do not have this address id %d", userid, v.Id)
		goto BOTTOM
	}

	err = models.UpdateEhomeFixAddressById(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeFixAddressById error! %v", err)
		goto BOTTOM
	}

	err = models.UpdateDefaultAddress(int64(v.Id), userid, v.IsDefaultaddr)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateDefaultAddress error! %v", err)
		goto BOTTOM
	}

	m, err = GetUserAddressList(userid)

	if err != nil {
		SetError(map2, DB_ERROR, "GetUserAddressList error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete address
// @Param	Mobileno  Mobileno string	false	"Mobileno"
// @Param	Usertype  Usertype  string	false	"Usertype"
// @Param	Token     Token     string	false	"Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Param   Id        Id        string  true     "Id"
// @Success 200 {object}
// @Failure 403
// @router /delete [get]
func (c *FixAddressController) Delete() {

	map2 := make(map[string]interface{})

	var m []interface{}
	var userid int
	var id int
	var err error

	Admin := c.GetString("Name")

	if Admin == "" {
		mobile, token, reqtime, _, err := GetComUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "user %s is not valid! [%v]", mobile, err)
			goto BOTTOM
		}
	} else {
		mobile := c.GetString("Mobileno")
		name, token, reqtime, err := GetAdminUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		err = models.CheckAdminUser(name, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "admin user %s is invalid", name)
			goto BOTTOM
		}
		userid, err = models.GetUseridByNo(mobile)
		if err != nil {
			SetError(map2, INVALID_USER, "mobileno [%s] is invalid!", mobile)
			goto BOTTOM
		}

	}

	id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param Id not valid! [%s]", c.GetString("Id"))
		goto BOTTOM
	}

	if !models.AddressIdBelongToUser(userid, id) {
		SetError(map2, DB_ERROR, "user %d do not have this address id %d", userid, id)
		goto BOTTOM
	}

	err = models.DeleteEhomeFixAddress(id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeFixAddress error!%v", err)
		goto BOTTOM
	}

	m, err = GetUserAddressList(userid)

	if err != nil {
		SetError(map2, DB_ERROR, "GetUserAddressList error!%v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Default ...
// @Title Default
// @Description get Default
// @Param	Mobileno  Mobileno string	false	"Mobileno"
// @Param	Usertype  Usertype  string	false	"Usertype"
// @Param	Token     Token     string	false	"Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Success 200 {object}
// @Failure 403
// @router /default [get]
func (c *FixAddressController) Default() {

	var err error
	map2 := make(map[string]interface{})

	var m []interface{}
	var userid int

	Admin := c.GetString("Name")

	if Admin == "" {
		mobile, token, reqtime, _, err := GetComUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "user %s is not valid! [%v]", mobile, err)
			goto BOTTOM
		}
	} else {
		mobile := c.GetString("Mobileno")
		name, token, reqtime, err := GetAdminUser(c, map2)
		if err != nil {
			goto BOTTOM
		}

		err = models.CheckAdminUser(name, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "admin user %s is invalid", name)
			goto BOTTOM
		}
		userid, err = models.GetUseridByNo(mobile)
		if err != nil {
			SetError(map2, INVALID_USER, "mobileno [%s] is invalid!", mobile)
			goto BOTTOM
		}

	}
	m, err = models.GetDefaultAddress(userid)

	if err != nil {
		SetError(map2, DB_ERROR, "GetDefaultAddress error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["records"] = m

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func CheckRegionId(region string) (err error) {
	_, err = models.GetRegionById(region)
	return
}
