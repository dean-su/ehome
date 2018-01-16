package controllers

import (
	"ehome/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type MasterbankController struct {
	beego.Controller
}

func (c *MasterbankController) URLMapping() {
	c.Mapping("Binding", c.Binding)
	c.Mapping("ModifyBinding", c.ModifyBinding)
	c.Mapping("DeleteBinding", c.DeleteBinding)
	c.Mapping("List", c.List)
}

// Binding account
// @Title binding account
// @Success 200 {object}
// @Failure 403 body is empty
// @router /binding [get]
func (c *MasterbankController) Binding() {
	map2 := make(map[string]interface{})
	var masterid int
	var v models.EhomeBindingAccount

	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	masterid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid master %s, %v", mobile, err)
		goto BOTTOM
	}

	v.Masterid = masterid
	v.Bank = c.GetString("Bank")
	v.Branch = c.GetString("Branch")
	v.Account = c.GetString("Account")
	if len(v.Bank) == 0 {
		SetError(map2, PARAM_ERR, "Param Bank error!")
		goto BOTTOM
	}

	if len(v.Account) == 0 {
		SetError(map2, PARAM_ERR, "Param Account error!")
		goto BOTTOM
	}

	_, err = models.AddEhomeBindingAccount(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeBindingAccount error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// modify Binding account
// @Title modify binding account
// @Success 200 {object}
// @Failure 403 body is empty
// @router /modifybinding [get]
func (c *MasterbankController) ModifyBinding() {
	map2 := make(map[string]interface{})
	var masterid int
	var v models.EhomeBindingAccount

	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	masterid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid master %s, %v", mobile, err)
		goto BOTTOM
	}

	v.Masterid = masterid
	v.Bank = c.GetString("Bank")
	v.Branch = c.GetString("Branch")
	v.Account = c.GetString("Account")
	if len(v.Bank) == 0 {
		SetError(map2, PARAM_ERR, "Param Bank error!")
		goto BOTTOM
	}

	if len(v.Account) == 0 {
		SetError(map2, PARAM_ERR, "Param Account error!")
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param  Id error!")
		goto BOTTOM
	}
	err = CheckBindingAccountUser(v.Masterid, v.Id)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckBindingAccountUser error! %v", err)
		goto BOTTOM
	}

	err = models.UpdateEhomeBindingAccountById(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeBindingAccountById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Delete Binding account
// @Title delete binding account
// @Success 200 {object}
// @Failure 403 body is empty
// @router /deletebinding [get]
func (c *MasterbankController) DeleteBinding() {
	map2 := make(map[string]interface{})
	var masterid int
	var id int

	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	masterid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid master %s, %v", mobile, err)
		goto BOTTOM
	}

	id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param  Id error!")
		goto BOTTOM
	}
	err = CheckBindingAccountUser(masterid, id)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckBindingAccountUser error! %v", err)
		goto BOTTOM
	}

	err = models.DeleteEhomeBindingAccount(id)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeBindingAccountById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// List ...
// @Title List binding bank
// @Description get all bank
// @Success 200 {object}
// @Failure 403 body is empty
// @router /list [get]
func (c *MasterbankController) List() {
	map2 := make(map[string]interface{})
	var masterid int
	var ml []interface{}

	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	masterid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid master %s, %v", mobile, err)
		goto BOTTOM
	}

	ml, err = MasterAccountlist(masterid)
	if err != nil {
		SetError(map2, DB_ERROR, "MasterAccountlist error %d, %v", masterid, err)
		goto BOTTOM
	}

	map2["records"] = ml
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ...
// @Title cash
// @Description get all bank
// @Success 200 {object}
// @Failure 403 body is empty
// @router /cash [get]
func (c *MasterbankController) Cash() {
	map2 := make(map[string]interface{})
	var masterid int

	var v models.EhomeMasterCashLog

	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	masterid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid master %s, %v", mobile, err)
		goto BOTTOM
	}

	v.Masterid = masterid
	v.Bank = c.GetString("Bank")
	v.Branch = c.GetString("Branch")
	v.Account = c.GetString("Account")
	v.Amount, err = c.GetFloat("Amount")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Amount invalid [%s]", c.GetString("Amount"))
		goto BOTTOM
	}
	if len(v.Bank) == 0 {
		SetError(map2, PARAM_ERR, "param  BANK invalid")
		goto BOTTOM
	}
	if len(v.Account) == 0 {
		SetError(map2, PARAM_ERR, "param  Account invalid")
		goto BOTTOM
	}

	_, err = models.AddEhomeMasterCashLog(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeBindingAccount error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func MasterAccountlist(masterid int) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Id", "Bank", "Branch", "Account")
	query["Masterid"] = strconv.Itoa(masterid)

	ml, err = models.GetAllEhomeBindingAccount(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		ml = make([]interface{}, 0)
	}
	return
}

func CheckBindingAccountUser(userid int, id int) (err error) {
	var v *models.EhomeBindingAccount
	v, err = models.GetEhomeBindingAccountById(id)
	if err != nil {
		return
	}

	if v.Masterid != userid {
		err = fmt.Errorf("Record Id %d does't not belong to user %d", id, userid)
	}

	return
}
