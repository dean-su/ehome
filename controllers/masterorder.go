package controllers

import (
	"ehome/models"
	_ "errors"

	"github.com/astaxie/beego"
)

type MasterorderController struct {
	beego.Controller
}

func (c *MasterorderController) URLMapping() {
	c.Mapping("RequestOrder", c.RequestOrder)
	c.Mapping("ConfirmOrder", c.ConfirmOrder)
	c.Mapping("PricingOrder", c.PricingOrder)
	c.Mapping("FinishConstruction", c.FinishConstruction)
	c.Mapping("Setout", c.Setout)
	c.Mapping("Arrived", c.Arrived)
	c.Mapping("ConfirmPaid", c.ConfirmPaid)
}

// RequestOrder ...
// @Title request order
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /request [get]
func (c *MasterorderController) RequestOrder() {

	var userid int
	var v models.EhomeRequestOrder

	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}
	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	userid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}

	err = models.IsMasterAudited(userid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	v.Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}

	v.Masterid = userid
	err = models.RequestOrder(&v)

	/*
		if err != nil {
			SetError(map2, DB_ERROR, "error! %v", err)
			goto BOTTOM
		} else {
			map2["status"] = 0
		}
	*/

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ConfirmOrder ...
// @Title confirm order
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /confirm [get]
func (c *MasterorderController) ConfirmOrder() {

	var userid int
	var v models.EhomeRequestOrder
	var Orderid int
	var Orderno string

	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}
	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	userid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}
	err = models.IsMasterAudited(userid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}
	Orderno = c.GetString("Orderno")

	err = models.CheckOrder(Orderid, Orderno, 0, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	v.Masterid = userid
	err = models.UpdateOrderStatus(Orderid, Orderno, models.ORDER_MASTER_ACCEPTED, "", "")
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}

	/*
		err = models.UpdateOrderMaster(Orderid, Orderno, userid)
		if err != nil {
			SetError(map2, DB_ERROR, "error! %v", err)
			goto BOTTOM
		}
			_, err = models.AddEhomeRequestOrder(&v)

			if err != nil {
				SetError(map2, DB_ERROR, "error! %v", err)
				goto BOTTOM
			}
	*/

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ...Pricing
// @Title confirm order
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /pricing [get]
func (c *MasterorderController) PricingOrder() {

	var Orderid int
	var Orderno string
	var v models.EhomePricingLog

	v.Reason = c.GetString("Reason")
	v.Image = c.GetString("Image")
	price, e := c.GetFloat("Price")
	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if e != nil {
		SetError(map2, PARAM_ERR, "Param Price error!")
		goto BOTTOM
	}

	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	v.Masterid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}

	err = models.IsMasterAudited(v.Masterid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}
	Orderno = c.GetString("Orderno")
	v.Orderid = Orderid
	v.Price = price
	err = models.CheckOrder(v.Orderid, Orderno, v.Masterid, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	err = models.UpdateOrderMasterPrice(Orderid, Orderno, price, v.Reason, v.Image)

	if err != nil {
		SetError(map2, DB_ERROR, "UpdateOrderMasterPrice error! %v", err)
		goto BOTTOM
	}

	_, err = models.AddEhomePricingLog(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "models.AddEhomePricingLog error! %v", err)
		goto BOTTOM
	}

	err = models.UpdateOrderStatus(Orderid, Orderno, models.ORDER_MASTER_PRICING, v.Reason, v.Image)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateOrderStatus error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ...
// @Title setout
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /setout [get]
func (c *MasterorderController) Setout() {

	var userid int
	var v models.EhomeRequestOrder
	var Orderid int
	var Orderno string

	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}
	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	userid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}

	err = models.IsMasterAudited(userid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}
	Orderno = c.GetString("Orderno")

	err = models.CheckOrder(Orderid, Orderno, userid, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	v.Masterid = userid
	err = models.UpdateOrderStatus(Orderid, Orderno, models.ORDER_MASTER_SETOUT, "", "")
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ...
// @Title setout
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /arrived [get]
func (c *MasterorderController) Arrived() {

	var userid int
	var v models.EhomeRequestOrder
	var Orderid int
	var Orderno string

	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}
	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	userid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}

	err = models.IsMasterAudited(userid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}
	Orderno = c.GetString("Orderno")

	err = models.CheckOrder(Orderid, Orderno, userid, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	v.Masterid = userid
	err = models.UpdateOrderStatus(Orderid, Orderno, models.ORDER_MASTER_ARRIVED, "", "")
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ...
// @Title finish construction
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /finishconstruction [get]
func (c *MasterorderController) FinishConstruction() {

	var userid int
	var v models.EhomeRequestOrder
	var Orderid int
	var Orderno string

	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}
	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	userid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}

	err = models.IsMasterAudited(userid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}
	Orderno = c.GetString("Orderno")

	err = models.CheckOrder(Orderid, Orderno, userid, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	v.Masterid = userid
	err = models.UpdateOrderStatus(Orderid, Orderno, models.ORDER_FINISH_CONSTRUCTION, "", "")
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// ...
// @Title finish construction
// @Description request order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /confirmpaid [get]
func (c *MasterorderController) ConfirmPaid() {

	var userid int
	var v models.EhomeRequestOrder
	var Orderid int
	var Orderno string

	map2 := make(map[string]interface{})
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}
	if usertype != 2 {
		SetError(map2, PARAM_ERR, "Param Usertype error!")
		goto BOTTOM
	}

	userid, err = models.ValidateMaster(mobile, token, reqtime)
	if err != nil {
		SetError(map2, PARAM_ERR, "ValidateMaster error! %v", err)
		goto BOTTOM
	}

	err = models.IsMasterAudited(userid)
	if err != nil {
		SetError(map2, 99, "IsMasterAudited error! %v", err)
		goto BOTTOM
	}

	Orderid, err = c.GetInt("Orderid")
	if err != nil {
		SetError(map2, PARAM_ERR, "Param orderid error! %v", err)
		goto BOTTOM
	}
	Orderno = c.GetString("Orderno")

	err = models.CheckOrder(Orderid, Orderno, userid, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	v.Masterid = userid
	err = models.UpdateOrderStatus(Orderid, Orderno, models.ORDER_MASTER_CONFIRM_PAY, "", "")
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
