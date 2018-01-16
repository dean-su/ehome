package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

// MasteruserController oprations for EcsUsers
type MasteruserController struct {
	beego.Controller
}

// URLMapping ...
func (c *MasteruserController) URLMapping() {
	c.Mapping("Info", c.Info)
}

// Get ...
// @Title Init
// @Description  init
// @Param	Masterid  Masterid string	false	"Filter ..."
// @Success 201 {int}
// @Failure 403 body is empty
// @router /info [get]
func (c *MasteruserController) Info() {

	var self bool
	var tmpid int
	var url string
	var ehomeu *models.EhomeMaster
	map2 := make(map[string]interface{})

	self = false

	masterid, e := c.GetInt("Masterid")
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)

	if err != nil {
		SetError(map2, INVALID_USER, "user not valid")
		goto BOTTOM
	}
	if usertype == 2 {
		tmpid, _ = models.ValidateMaster(mobile, token, reqtime)
		if masterid == 0 {
			masterid = tmpid
		}
		if masterid == tmpid {
			self = true
		}
	} else if e != nil {
		SetError(map2, PARAM_ERR, "Param Masterid not valid!")
		goto BOTTOM
	}

	ehomeu, err = models.GetEhomeMasterById(masterid)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeMasterById error!")
		goto BOTTOM
	}

	if ehomeu.Headimageid != 0 {
		url, err = models.GetImageById(ehomeu.Headimageid)
		if err != nil {
			SetError(map2, PARAM_ERR, "GetImageById error!%v", err)
			goto BOTTOM
		}
	}

	map2["status"] = 0
	map2["Name"] = ehomeu.Name
	map2["Headimage"] = url
	map2["Audited"] = ehomeu.Audited
	if self {
		map2["Balance"] = ehomeu.Balance
		map2["Bonuspoint"] = ehomeu.Bonuspoint
	}

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
