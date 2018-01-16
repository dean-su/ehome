package controllers

import (
	"ehome/models"
	_ "encoding/json"
	_ "errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (c *OrderController) URLMapping() {
	c.Mapping("GetAllOrder", c.GetAllOrder)
	c.Mapping("GetOrderNum", c.GetOrderNum)
	c.Mapping("GetOrderPage", c.GetOrderPage)
	c.Mapping("PlaceOrder", c.PlaceOrder)
	c.Mapping("Status", c.Status)
	c.Mapping("Paycash", c.Paycash)
	c.Mapping("Cancel", c.Cancel)
	c.Mapping("ConfirmPricing", c.ConfirmPricing)
	c.Mapping("CheckAndAccept", c.CheckAndAccept)
}

// GetAllOrder ...
// @Title Get all order
// @Description get all order of a specific user
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /all [get]
func (c *OrderController) GetAllOrder() {
	var limit int64
	var offset int64

	var l []interface{}
	var userid int

	map2 := make(map[string]interface{})
	reqtype := c.GetString("Reqtype")
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if usertype == 1 {
		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			map2["status"] = 2
			goto BOTTOM
		}
	} else {
		l, err = models.GetEhomeMasterByPhone(mobile)
		if len(l) == 0 {
			SetError(map2, PARAM_ERR, "Master %s not exist!", mobile)
			goto BOTTOM
		} else if err != nil {
			SetError(map2, DB_ERROR, "GetEhomeMasterByPhone error! %v", err)
			goto BOTTOM
		}
		userid = (l[0].(models.EhomeMaster)).Id
	}

	l, err = models.OrderList(reqtype, usertype, userid, limit, offset)

	if err != nil {
		SetError(map2, DB_ERROR, "OrderList error! %v", err)
		goto BOTTOM
	} else {
		map2["status"] = 0
		map2["records"] = l
	}
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// GetOrderNum...
// @Title Get order  number
// @Description get all order numbers of a specific user
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   UserType        UserType string true        "user type"
// @Param   Token           Token    string true        "token"
// @Success 200 {object}
// @Failure 403 body is empty
// @router /number [get]
func (c *OrderController) GetOrderNum() {

	l, err := models.OrderNum()
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

// GetOrderNum...
// @Title Get order  page
// @Description get orders of a page of a specific user
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   Usertype        Usertype string true        "user type"
// @Param   Token           Token    string true        "token"
// @Param   Pagenum         Pagenum  string true        "page number"
// @Param   Recperpage      Recperpage string true      "rec per page"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /page [get]
func (c *OrderController) GetOrderPage() {
	var limit int64
	var offset int64

	n, e := c.GetInt64("Pagenum")
	m, ee := c.GetInt64("Recperpage")
	beego.Info("page :", n, "|", m)

	var userid int
	var l []interface{}

	map2 := make(map[string]interface{})
	reqtype := c.GetString("Reqtype")
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if e != nil || ee != nil || m < 1 || n < 1 {
		SetError(map2, PARAM_ERR, "Param Pagenum or Recperpage not valid [%d:%d]", m, n)
		goto BOTTOM
	}

	if usertype == 1 {
		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			map2["status"] = 2
			goto BOTTOM
		}
	} else {
		l, err = models.GetEhomeMasterByPhone(mobile)
		if len(l) == 0 {
			SetError(map2, PARAM_ERR, "Master %s not exist!", mobile)
			goto BOTTOM
		} else if err != nil {
			SetError(map2, DB_ERROR, "GetEhomeMasterByPhone error! %v", err)
			goto BOTTOM
		}
		userid = (l[0].(models.EhomeMaster)).Id
	}

	n = n - 1
	limit = m
	offset = m * n
	l, err = models.OrderList(reqtype, usertype, userid, limit, offset)

	if err != nil {
		SetError(map2, DB_ERROR, "OrderList error! %v", err)
		goto BOTTOM
	} else {
		map2["status"] = 0
		map2["records"] = l
	}

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// PlaceOrder
// @Title place an order
// @Description get orders of a page of a specific user
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   Token           Token    string true        "token"
// @Param   Reqtime         Reqtime  string true        "Reqtime"
// @Param   ImageId         Pagenum  string true        "page number"
// @Param   VoiceID         VoiceID  string true        "rec per page"
// @Param   FixType         FixType  string true        "Fix type"
// @Param   LabourId        LabourId string true        "Labour id"
// @Param   MaterialId      MaterialId string true      "Material id“
// @Param   AppointmenTime  AppointmenTime string true  "appointment time"
// @Param   AddrId          AddrId         string true  "AddrId"
// @Param   Attact          Attact         string  true  "Attact"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /place [get]
func (c *OrderController) PlaceOrder() {
	beego.Info("input:", c.Input())
	mobile := c.GetString("Mobileno")
	token := c.GetString("Token")
	reqtime, err := c.GetInt64("Reqtime")
	imageid := c.GetString("ImageId")
	voiceid := c.GetString("VoiceId")
	attact := c.GetString("Attact")
	appointmenttime, _ := c.GetInt64("AppintmentTime")
	Type := c.GetString("Type")

	var labourid string
	var materialid string
	var fixtype string

	var TotalPrice float64
	var addrid int

	map2 := make(map[string]interface{})

	var province, city, region string

	var ehomeaddr *models.EhomeFixAddress
	var ehomeo models.EhomeOrder
	var userid int
	var tmp []string
	var ti int
	var orderid int64

	var tmp_price float64

	if err != nil {
		map2["Status"] = 1
		//SetErrmsg(map2, "Param Reqtime error! [%s]", c.GetString("Reqtime"))
		goto BOTTOM
	}

	if Type == "1" {
		ehomeo.Type = 1
		ehomeo.Room, err = c.GetInt8("Room")
		if err != nil {
			map2["Status"] = PARAM_ERR
			SetErrmsg(map2, "Param Room erro! [%s]", c.GetString("Room"))
			goto BOTTOM
		}
		ehomeo.Hall, err = c.GetInt8("Hall")
		if err != nil {
			map2["Status"] = PARAM_ERR
			SetErrmsg(map2, "Param Hall erro! [%s]", c.GetString("Hall"))
			goto BOTTOM
		}
		ehomeo.Kitchen, err = c.GetInt8("Kitchen")
		if err != nil {
			map2["Status"] = PARAM_ERR
			SetErrmsg(map2, "Param Kitchen erro! [%s]", c.GetString("Kitchen"))
			goto BOTTOM
		}
		ehomeo.Toilet, err = c.GetInt8("Toilet")
		if err != nil {
			map2["Status"] = PARAM_ERR
			SetErrmsg(map2, "Param Toilet erro! [%s]", c.GetString("Toilet"))
			goto BOTTOM
		}

		ehomeo.Size, err = c.GetFloat("Size")
		/*
			if err != nil {
				map2["Status"] = PARAM_ERR
				SetErrmsg(map2, "Param Size erro! [%s]", c.GetString("Size"))
				goto BOTTOM
			}
		*/

		ehomeo.Catidlist = c.GetString("StyleId")
		err = models.CheckStyleId(ehomeo.Catidlist)
		if err != nil {
			map2["Status"] = PARAM_ERR
			SetErrmsg(map2, "Param StyleId error! [%s]", ehomeo.Catidlist)
			goto BOTTOM
		}

		ehomeo.Stylepriceid, err = c.GetInt("StylePriceId")
		if err != nil {
			map2["Status"] = PARAM_ERR
			SetErrmsg(map2, "Param StylepriceId erro! [%s]", c.GetString("StylepriceId"))
			goto BOTTOM
		}

		fixtype, _ = models.GetWholeFixType(ehomeo.Catidlist, ehomeo.Stylepriceid)

	} else {
		ehomeo.Type = 0
		fixtype = c.GetString("FixType")
		labourid = c.GetString("LabourId")
		materialid = c.GetString("MaterialId")
		ehomeo.Catidlist, err = models.GetCatidlist(fixtype)
		if err != nil {
			map2["Status"] = 5
			SetErrmsg(map2, "Param fixtype error! [%s]", fixtype)
			goto BOTTOM
		}
	}

	TotalPrice, err = c.GetFloat("TotalPrice")
	if err != nil {
		map2["Status"] = 2
		goto BOTTOM
	}
	tmp = strings.Split(imageid, ",")
	/*
		if len(tmp) < 1 {
			map2["Status"] = 3
			goto BOTTOM
		}
	*/
	ti, err = strconv.Atoi(tmp[0])
	/*
		if err != nil {
			map2["Status"] = 4
			goto BOTTOM
		}
	*/

	addrid, err = c.GetInt("AddrId")
	ehomeaddr, err = models.GetEhomeFixAddressById(addrid)
	if err != nil {
		map2["Status"] = 10
		goto BOTTOM
	}

	userid, err = models.ValidateUser(mobile, token, reqtime)
	if err != nil {
		map2["Status"] = 20
		SetErrmsg(map2, "user not register or not login! [%v]", err)
		goto BOTTOM
	}
	ehomeo.Imageid = ti
	ehomeo.Orderno = models.GetOrderNo()
	ehomeo.Status = 1
	ehomeo.Imageidlist = imageid
	ehomeo.Voiceidlist = voiceid
	ehomeo.Labouridlist = labourid
	ehomeo.Materialidlist = materialid
	ehomeo.Userid = int(userid)
	ehomeo.Price = TotalPrice
	ehomeo.CreateTime = time.Now()
	ehomeo.Region = ehomeaddr.Region
	ehomeo.Cityid = ehomeaddr.Cityid
	ehomeo.Contactaddr = ehomeaddr.Contactaddr
	ehomeo.Contactname = ehomeaddr.Contactname
	ehomeo.Contactphone = ehomeaddr.Phone
	ehomeo.Appointmenttime = time.Unix(appointmenttime, 0)
	//beego.Error("mxz ", ehomeo.Appointmenttime.Format("2006-01-02 15:04:05"))
	ehomeo.Attact = attact

	tmp_price, err = models.CalTotalPrice(Type, ehomeo.Stylepriceid, ehomeo.Size, labourid, materialid)
	if err != nil {
		map2["Status"] = 15
		goto BOTTOM
	}

	if !models.IsFloatEqual(tmp_price, ehomeo.Price) {
		map2["Status"] = 6
		SetErrmsg(map2, "not equal %.2f %.2f", tmp_price, ehomeo.Price)
		beego.Error("not equal", tmp_price, ehomeo.Price)
		goto BOTTOM
	}

	orderid, err = models.AddEhomeOrder(&ehomeo)
	if err != nil {
		beego.Error("AddEhomeOrd error", err)
		map2["Status"] = 7
		goto BOTTOM
	}

	err = models.InsertOrderPath(ehomeo.Orderno)
	if err != nil {
		map2["Status"] = 8
		goto BOTTOM
	}

	//	ehomeo.CreateTime.

	map2["Status"] = 0
	map2["Orderno"] = ehomeo.Orderno
	map2["Orderid"] = orderid
	map2["OrderTime"] = ehomeo.CreateTime.Format("20060102150405")
	map2["Price"] = TotalPrice

	province, city, region, err = models.GetRegionDetailById(ehomeo.Region)

	models.SendMail(ehomeo.CreateTime.Format("2006-01-02 15:04:05"), ehomeo.Orderno, ehomeo.Appointmenttime.Format("2006-01-02 15:04:05"), fixtype, TotalPrice, ehomeo.Attact, ehomeo.Contactname, ehomeo.Contactphone, province+city+region, ehomeo.Contactaddr)

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// GetOrderStatus
// @Title get order status
// @Description get orders of a page of a specific user
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   Token           Token    string true        "token"
// @Param   Reqtime         Reqtime  string true        "Reqtime"
// @Param   Orderid         Orderid  string true        "Orderid"
// @Param   Orderno         Orderno  string true        "Orderno"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /status [get]
func (c *OrderController) Status() {
	orderid, oerr := c.GetInt("Orderid")
	orderno := c.GetString("Orderno")
	var v *models.EhomeOrder
	map2 := make(map[string]interface{})

	var province, city, region string
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if oerr != nil {
		beego.Error(err, oerr)
		map2["Status"] = 1
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		map2["Status"] = INVALID_USER
		SetErrmsg(map2, "user not valid")
		goto BOTTOM
	}

	v, err = models.GetEhomeOrderById(orderid)
	if err != nil {
		map2["Status"] = INVALID_ORDER
		SetErrmsg(map2, "orderid invalid %d", orderid)
		goto BOTTOM
	}

	if orderno != v.Orderno {
		map2["Status"] = INVALID_ORDER
		SetErrmsg(map2, "orderid invalid orderid[%d] orderno[%s] ", orderid, orderno)
		goto BOTTOM
	}

	map2["records"], err = GetOrderStatus(orderno, v.Status, v.ModifyTime, usertype)
	if err != nil {
		map2["Status"] = INVALID_ORDER
		SetErrmsg(map2, "GetOrderStatus error! orderid[%d] orderno[%s] ", orderid, orderno)
		delete(map2, "records")
		goto BOTTOM
	}

	map2["Type"] = v.Type
	if v.Type == 0 {
		map2["Fixtype"], err = models.GetFixtypelistbyIdlist(v.Catidlist)
		if err != nil {
			map2["Status"] = INVALID_ORDER
			SetErrmsg(map2, "GetFixtypelistbyIdlist error! catidlist [%s]", v.Catidlist)
			goto BOTTOM
		}
	} else {
		map2["Fixtype"], err = models.GetWholeFixType(v.Catidlist, v.Stylepriceid)
		if err != nil {
			map2["Status"] = INVALID_ORDER
			SetErrmsg(map2, "GetWholeFixType error! catidlist [%s]", v.Catidlist)
			goto BOTTOM
		}
	}

	map2["ImageList"], err = models.GetImageList(v.Imageidlist)
	if err != nil {
		map2["Status"] = 8
		goto BOTTOM
	}

	map2["Voiceidlist"], err = models.GetVoiceList(v.Voiceidlist)
	if err != nil {
		map2["Status"] = 9
		SetErrmsg(map2, "Voiceidlist not valid [%s]", v.Voiceidlist)
		goto BOTTOM
	}
	province, city, region, err = models.GetRegionDetailById(v.Region)
	if err != nil {
		map2["Status"] = DB_ERROR
		SetErrmsg(map2, "region [%s] from order [%s] is not valid!", v.Region, v.Orderno)
		goto BOTTOM
	}

	map2["Appointmenttime"] = v.Appointmenttime.Format("20060102150405")

	map2["Status"] = 0

	map2["Orderno"] = v.Orderno
	map2["Orderid"] = orderid
	map2["OrderTime"] = v.CreateTime.Format("20060102150405")
	map2["Attact"] = v.Attact

	map2["Price"] = v.Price
	map2["Masterprice"] = v.Masterprice
	map2["Contactname"] = v.Contactname
	map2["Contactphone"] = v.Contactphone
	map2["Contactaddr"] = v.Contactaddr

	map2["Region"] = province + city + region
	map2["Regionid"] = v.Region

	map2["Masterid"] = v.Masterid
	map2["Appearance"] = v.Appearance
	map2["Punctual"] = v.Punctual
	map2["Service"] = v.Service
	map2["Quality"] = v.Quality
	map2["Feeback"] = v.Feeback
	map2["Shareimage"], _ = models.GetImageList(v.Shareimage)
	map2["Rejectpricereason"] = v.Rejectpricereason
	map2["Rejectpriceimage"], _ = models.GetImageList(v.Rejectpriceimage)
	map2["Pricereason"] = v.Pricereason
	map2["Priceimage"], _ = models.GetImageList(v.Priceimage)

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetOrderStatus(orderno string, status int, modifyTime time.Time, usertype int) (ml []interface{}, err error) {

	if status == models.ORDER_CANCEL {
		ml = make([]interface{}, 1)
		ml[0] = make(map[string]interface{})
		(ml[0].(map[string]interface{}))["Id"] = 0
		(ml[0].(map[string]interface{}))["Index"] = 1
		(ml[0].(map[string]interface{}))["Introduction"] = "已取消"
		(ml[0].(map[string]interface{}))["FinishTime"] = modifyTime.Format("2006-01-02 15:04:05")
		return
	}

	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Id", "Index", "Introduction", "FinishTime", "Masterintro")
	query["Orderno"] = orderno
	var i int

	sortby = append(sortby, "Index")
	order = append(order, "asc")

	ml, err = models.GetAllEhomeOrderPath(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No this  order %s", orderno)
		return
	}

	for i = 0; i < len(ml); i++ {
		if (ml[i].(map[string]interface{})["Index"]).(int) > status {
			ml[i].(map[string]interface{})["FinishTime"] = ""
		} else {
			ml[i].(map[string]interface{})["FinishTime"] = (ml[i].(map[string]interface{})["FinishTime"].(time.Time)).Format("2006-01-02 15:04:05")

		}
		if status+1 == (ml[i].(map[string]interface{})["Index"]).(int) {
			ml[i].(map[string]interface{})["Next"] = 1
		} else {
			ml[i].(map[string]interface{})["Next"] = 0
		}

		if usertype == 2 {
			ml[i].(map[string]interface{})["Introduction"] = ml[i].(map[string]interface{})["Masterintro"]
		}
		delete(ml[i].(map[string]interface{}), "Masterintro")
	}
	if usertype == 2 {
		ml = ml[1:len(ml)]
	}

	return
}

// GetOrderStatus
// @Title get order status
// @Description get orders of a page of a specific user
// @Param   Mobileno        Mobileno string true        "phone number"
// @Param   Token           Token    string true        "token"
// @Param   Reqtime         Reqtime  string true        "Reqtime"
// @Param   Orderid         Orderid  string true        "Orderid"
// @Param   Orderno         Orderno  string true        "Orderno"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /evaluate [get]
func (c *OrderController) Evaluate() {

	map2 := make(map[string]interface{})

	var tmpi int
	var v models.OrderEvalue
	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	v.Orderid = c.GetString("Orderid")
	v.Orderno = c.GetString("Orderno")
	v.Appearance = c.GetString("Appearance")
	v.Punctual = c.GetString("Punctual")
	v.Service = c.GetString("Service")
	v.Quality = c.GetString("Quality")
	v.Feeback = c.GetString("Feeback")
	v.Shareimage = c.GetString("Shareimage")
	v.Userid, err = models.ValidateUser(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid  user%s", mobile)
		goto BOTTOM
	}
	if len(v.Orderno) == 0 {
		SetError(map2, PARAM_ERR, "Orderno is empty")
		goto BOTTOM
	}

	if len(v.Appearance) == 0 {
		SetError(map2, PARAM_ERR, "Appearance is empty")
		goto BOTTOM
	}

	if len(v.Punctual) == 0 {
		SetError(map2, PARAM_ERR, "Punctual is empty")
		goto BOTTOM
	}

	if len(v.Service) == 0 {
		SetError(map2, PARAM_ERR, "Service is empty")
		goto BOTTOM
	}

	if len(v.Quality) == 0 {
		SetError(map2, PARAM_ERR, "Quality is empty")
		goto BOTTOM
	}

	if len(v.Feeback) == 0 {
		SetError(map2, PARAM_ERR, "Feeback is empty")
		goto BOTTOM
	}

	/*
		if len(v.Shareimage) == 0 {
			SetError(map2, PARAM_ERR, "Shareimage is empty")
			goto BOTTOM
		}
	*/

	tmpi, err = strconv.Atoi(v.Orderid)
	if err != nil {
		SetError(map2, PARAM_ERR, "Param Orderid not valid %s", v.Orderid)
		goto BOTTOM
	}

	err = CheckOrder(v.Userid, tmpi, v.Orderno)
	if err != nil {
		SetError(map2, INVALID_ORDER, "CheckOrder error! %v", err)
		goto BOTTOM
	}

	err = models.EvaluteOrder(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "EvaluteOrder error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Paycase
// @Title get order status
// @Description get orders of a page of a specific user
// @Param   Mobileno        Mobileno string true        "phone number"
// @Param   Token           Token    string true        "token"
// @Param   Reqtime         Reqtime  string true        "Reqtime"
// @Param   Orderid         Orderid  string true        "Orderid"
// @Param   Orderno         Orderno  string true        "Orderno"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /paycash [get]
func (c *OrderController) Paycash() {

	map2 := make(map[string]interface{})

	var tmpi int
	var Userid int
	var Price float64
	mobile, token, reqtime, _, err := GetComUser(c, map2)
	Orderid := c.GetString("Orderid")
	Orderno := c.GetString("Orderno")

	var v *models.EhomeOrder

	if err != nil {
		goto BOTTOM
	}

	Userid, err = models.ValidateUser(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid  user%s", mobile)
		goto BOTTOM
	}
	if len(Orderno) == 0 {
		SetError(map2, PARAM_ERR, "Orderno is empty")
		goto BOTTOM
	}

	tmpi, err = strconv.Atoi(Orderid)
	if err != nil {
		SetError(map2, PARAM_ERR, "Param Orderid not valid %s", Orderid)
		goto BOTTOM
	}

	Price, err = c.GetFloat("Price")
	/*
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Price not float %f", Price)
			goto BOTTOM
		}
	*/

	v, err = CheckOrderPrice(Userid, tmpi, Price)
	if err != nil {
		SetError(map2, INVALID_ORDER, "CheckOrderPrice error! %v", err)
		goto BOTTOM
	}

	err = AddCashLog(tmpi, Orderno, GetOrderPrice(tmpi))
	if err != nil {
		SetError(map2, DB_ERROR, "AddCashLog error! %v", v)
		goto BOTTOM
	}

	v.Status = models.ORDER_CLIENT_PAY
	models.UpdateEhomeOrderById(v)

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetOrderPrice(orderid int) (price float64) {
	v, err := models.GetEhomeOrderById(orderid)
	if err != nil {
		return 0
	}
	return v.Masterprice
}

func CheckOrderPrice(userid int, orderid int, price float64) (*models.EhomeOrder, error) {
	v, err := models.GetEhomeOrderById(orderid)
	if err != nil {
		return nil, err
	}

	if v.Userid != userid {
		return nil, fmt.Errorf("order id doesn't belong to this user")
	}

	/*
		if (v.Price-price < 0.01 && v.Price-price >= 0) || (price-v.Price < 0.01 && price-v.Price >= 0) {
			return v, nil
		} else {
			return nil, fmt.Errorf("Price error! %f != %f", v.Price, price)
		}
	*/
	return v, nil
}

func CheckOrder(userid int, orderid int, orderno string) error {
	v, err := models.GetEhomeOrderById(orderid)
	if err != nil {
		return err
	}

	if v.Userid != userid && userid != 0 {
		return fmt.Errorf("order id doesn't belong to this user")
	}

	if v.Orderno != orderno && orderno != "" {
		return fmt.Errorf("orderid and orderno is not the same order")
	}
	return nil
}

func AddCashLog(Orderid int, Orderno string, Amount float64) (err error) {
	var v models.EhomePaycashLog
	v.Orderid = Orderid
	v.Orderno = Orderno
	v.Amount = Amount
	_, err = models.AddEhomePaycashLog(&v)
	return
}

// Cancel
// @Title cancel order
// @Description cancel an order
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   Token           Token    string true        "token"
// @Param   Reqtime         Reqtime  string true        "Reqtime"
// @Param   Orderid         Orderid  string true        "Orderid"
// @Param   Orderno         Orderno  string true        "Orderno"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /cancel [get]
func (c *OrderController) Cancel() {
	orderid, oerr := c.GetInt("Orderid")
	orderno := c.GetString("Orderno")
	reason := c.GetString("Reason")
	image := c.GetString("Image")
	var v *models.EhomeOrder
	map2 := make(map[string]interface{})

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if oerr != nil {
		SetError(map2, PARAM_ERR, "Param Orderid error! %s", c.GetString("Orderid"))
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, INVALID_USER, "CheckComUser error! %v", err)
		goto BOTTOM
	}

	v, err = models.GetEhomeOrderById(orderid)
	if err != nil {
		SetError(map2, INVALID_ORDER, "GetEhomeOrderById error! %d", orderid)
		goto BOTTOM
	}

	if orderno != v.Orderno {
		SetError(map2, INVALID_ORDER, "orderid or orderno invalid orderid[%d] orderno[%s] ", orderid, orderno)
		goto BOTTOM
	}

	err = models.UpdateOrderStatus(v.Id, v.Orderno, models.ORDER_CANCEL, reason, image)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateOrderStatus erorr! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Confirm Price
// @Title Confirm Price
// @Description confirm Order Price
// @Param	Mobileno		Mobileno string	true		"phone number"
// @Param   Token           Token    string true        "token"
// @Param   Reqtime         Reqtime  string true        "Reqtime"
// @Param   Orderid         Orderid  string true        "Orderid"
// @Param   Orderno         Orderno  string true        "Orderno"
// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /confirmpricing [get]
func (c *OrderController) ConfirmPricing() {

	accepted := c.GetString("Accepted")

	orderid, oerr := c.GetInt("Orderid")
	orderno := c.GetString("Orderno")
	reason := c.GetString("Reason")
	image := c.GetString("Image")

	var status int
	var v *models.EhomeOrder
	map2 := make(map[string]interface{})

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if oerr != nil {
		SetError(map2, PARAM_ERR, "Param Orderid error! %s", c.GetString("Orderid"))
		goto BOTTOM
	}

	if usertype != 1 {
		SetError(map2, PARAM_ERR, "PARA Usertype error! %d", usertype)
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, INVALID_USER, "CheckComUser error! %v", err)
		goto BOTTOM
	}

	v, err = models.GetEhomeOrderById(orderid)
	if err != nil {
		SetError(map2, INVALID_ORDER, "GetEhomeOrderById error! %d", orderid)
		goto BOTTOM
	}

	if orderno != v.Orderno {
		SetError(map2, INVALID_ORDER, "orderid or orderno invalid orderid[%d] orderno[%s] ", orderid, orderno)
		goto BOTTOM
	}

	if accepted == "1" {
		status = models.ORDER_CLIENT_ACCEPT_PRICING
	} else {
		status = models.ORDER_MASTER_ARRIVED
	}

	err = models.UpdateClientConfirmPriceStatus(v.Id, v.Orderno, status, reason, image)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateClientConfirmPriceStatus erorr! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /checkandaccept [get]
func (c *OrderController) CheckAndAccept() {

	orderid, oerr := c.GetInt("Orderid")
	orderno := c.GetString("Orderno")
	reason := c.GetString("Reason")
	image := c.GetString("Image")

	var v *models.EhomeOrder
	map2 := make(map[string]interface{})

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)

	if err != nil {
		goto BOTTOM
	}

	if oerr != nil {
		SetError(map2, PARAM_ERR, "Param Orderid error! %s", c.GetString("Orderid"))
		goto BOTTOM
	}

	if usertype != 1 {
		SetError(map2, PARAM_ERR, "PARA Usertype error! %d", usertype)
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, INVALID_USER, "CheckComUser error! %v", err)
		goto BOTTOM
	}

	v, err = models.GetEhomeOrderById(orderid)
	if err != nil {
		SetError(map2, INVALID_ORDER, "GetEhomeOrderById error! %d", orderid)
		goto BOTTOM
	}

	if orderno != v.Orderno {
		SetError(map2, INVALID_ORDER, "orderid or orderno invalid orderid[%d] orderno[%s] ", orderid, orderno)
		goto BOTTOM
	}

	err = models.UpdateOrderStatus(v.Id, v.Orderno, models.ORDER_CHECK_AND_ACCEPT, reason, image)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateOrderStatus erorr! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
