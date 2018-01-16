package admins

import (
	"ehome/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

type OrderController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrderController) URLMapping() {
	c.Mapping("GetRegion", c.GetRegion)
	c.Mapping("GetStatusList", c.GetStatusList)
	c.Mapping("OrderList", c.OrderList)
	c.Mapping("Detail", c.Detail)
	c.Mapping("Audit", c.Audit)
	c.Mapping("DealCash", c.DealCash)
	c.Mapping("PlaceOrder", c.PlaceOrder)
	c.Mapping("Modify", c.Modify)
	c.Mapping("Requestmasterlist", c.Requestmasterlist)
	c.Mapping("Assignmaster", c.Assignmaster)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Modifystatusandprice", c.Modifystatusandprice)
}

// @param
// @Failure 403 body is empty
// @router /region [post]
func (c *OrderController) GetRegion() {
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

	map2["records"], err = models.GetOrderRegion()
	if err != nil {
		SetError(map2, DB_ERROR, "GetOrderRegion error! %v", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /statuslist [post]
func (c *OrderController) GetStatusList() {
	map2 := make(map[string]interface{})

	gettotal := c.GetString("Gettotal")
	ordertype := c.GetString("Ordertype")
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = models.GetOrderStatusList(gettotal, ordertype)
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func Containfixtypeid(catlist, catid string) bool {
	sl := strings.Split(catlist, ",")
	for i := 0; i < len(sl); i++ {
		if catid == sl[i] {
			return true
		}
	}
	return false
}

func GetAllOrder(limit, offset int64, cityid, regionid, statusid, ordertype, fixtypeid, mastername, month string) (ml []interface{}, err error) {

	var cond models.EhomeOrder_Str

	cond.Cityid = cityid
	cond.Regionid = regionid
	cond.Status = statusid
	cond.Type = ordertype
	cond.Catidlist = fixtypeid
	cond.Mastername = mastername

	ml, err = models.EhomeOrder_list(limit, offset, &cond)

	if err != nil {
		return
	}

	if len(ml) <= 0 {
		ml = make([]interface{}, 0)
	}

	for i := int(0); i < len(ml); i++ {
		var master *models.EhomeMaster
		if !(ml[i].(map[string]interface{})["Masterid"].(int) == 0) {
			master, err = models.GetEhomeMasterById((ml[i].(map[string]interface{}))["Masterid"].(int))
			if err != nil {
				beego.Error("can't find master by id", (ml[i].(map[string]interface{}))["Masterid"].(int))
				continue
			}
			(ml[i].(map[string]interface{}))["Mastername"] = master.Name
			(ml[i].(map[string]interface{}))["Masterphone"] = master.Phone
		}

		(ml[i].(map[string]interface{}))["Province"], (ml[i].(map[string]interface{}))["City"], (ml[i].(map[string]interface{}))["Region"], err = models.GetRegionDetailById(((ml[i].(map[string]interface{}))["Regionid"]).(string))
		if err != nil {
			beego.Error("region error %v %v", err, (ml[i].(map[string]interface{}))["Regionid"])
			return
		}

		if (ml[i].(map[string]interface{}))["Type"].(int8) == 1 {
			(ml[i].(map[string]interface{}))["Catdesc"], err = models.GetWholeFixType((ml[i].(map[string]interface{}))["Catidlist"].(string), (ml[i].(map[string]interface{}))["Stylepriceid"].(int))
			if err != nil {
				beego.Error("GetWholeFixType error! catidlist ", (ml[i].(map[string]interface{}))["Catidlist"].(string), (ml[i].(map[string]interface{}))["Stylepriceid"].(int))
				return
			}

		} else {
			(ml[i].(map[string]interface{}))["Catdesc"], err = models.GetFixtypelistbyIdlist((ml[i].(map[string]interface{}))["Catidlist"].(string))
			if err != nil {
				//beego.Error("mxz ", (ml[i].(map[string]interface{}))["Catidlist"])
				beego.Error("Catidlist error %v %v", err, (ml[i].(map[string]interface{}))["Catidlist"])
				return
			}
		}

		(ml[i].(map[string]interface{}))["Statusdesc"], err = models.GetStatusById((ml[i].(map[string]interface{}))["Status"].(int), 1)

	}

	return
}

// @param
// @Failure 403 body is empty
// @router /list [post]
func (c *OrderController) OrderList() {
	map2 := make(map[string]interface{})

	var m int64
	limit, offset := GetPage(c)

	regionid := c.GetString("Regionid")
	cityid := c.GetString("Cityid")
	statusid := c.GetString("Statusid")
	ordertype := c.GetString("Ordertype")
	fixtypeid := c.GetString("Fixtypeid")
	mastername := c.GetString("Mastername")
	month := c.GetString("Month")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	map2["records"], err = GetAllOrder(limit, offset, cityid, regionid, statusid, ordertype, fixtypeid, mastername, month)

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllOrder error! %s", err)
		delete(map2, "records")
		goto BOTTOM
	}

	m, err = models.OrderListNum(cityid, regionid, statusid, ordertype, fixtypeid, mastername, month)

	if err != nil {
		SetError(map2, DB_ERROR, "OrderListNum error!")
		goto BOTTOM
	} else {
		map2["status"] = 0
		map2["Total"] = m
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetOrderById(id int) (m map[string]interface{}, err error) {
	var v *models.EhomeOrder
	var u *models.EcsUsers
	v, err = models.GetEhomeOrderById(id)
	if err != nil {
		return
	}

	m = models.Struct2Map(*v)

	u, err = models.GetEcsUsersById(v.Userid)
	if err != nil {
		beego.Error("order's userid is invalid!", id)
		return
	}

	m["Mobileno"] = u.MobilePhone

	if v.Type == 0 {
		m["Fixtype"], err = models.GetFixtypelistbyIdlist(v.Catidlist)
		if err != nil {
			fmt.Errorf("GetFixtypelistbyIdlist error! catidlist [%s]!%v", v.Catidlist, err)
			return
		}
	} else {
		m["Fixtype"], err = models.GetWholeFixType(v.Catidlist, v.Stylepriceid)
		if err != nil {
			err = fmt.Errorf("GetWholeFixType error! catidlist [%s]!%v", v.Catidlist, err)
			return
		}
	}

	m["ImageList"], err = models.GetImageList(v.Imageidlist)
	if err != nil {
		beego.Error("GetImageList error!Imageidlist [%s]!%v", v.Imageidlist, err)
		err = fmt.Errorf("GetImageList error!Imageidlist [%s]!%v", v.Imageidlist, err)
		return
	}

	m["Voiceidlist"], err = models.GetVoiceList(v.Voiceidlist)
	if err != nil {
		err = fmt.Errorf("GetVoiceList error!Voiceidlist [%s]!%v", v.Voiceidlist, err)
		return
	}

	m["Regionid"] = v.Region
	m["Province"], m["City"], m["Region"], err = models.GetRegionDetailById(v.Region)
	if err != nil {
		err = fmt.Errorf("region [%s] from order [%s] is not valid!%v", v.Region, v.Orderno, err)
		return
	}

	m["Shareimagelist"], _ = models.GetImageList(v.Shareimage)
	m["Rejectpriceimagelist"], _ = models.GetImageList(v.Rejectpriceimage)
	m["Priceimagelist"], _ = models.GetImageList(v.Priceimage)

	m["Orderstatus"], err = models.GetOrderStatus(v.Orderno, v.Status, v.ModifyTime, 1)
	if err != nil {
		err = fmt.Errorf("GetOrderStatus error! %v", err)
		return
	}
	m["StatusDesc"], err = models.GetStatusById(v.Status, 1)
	if err != nil {
		err = fmt.Errorf("GetStatusById error! %v", err)
		return
	}

	tm, e := models.GetCityById(v.Cityid)
	if e != nil {
		err = e
		beego.Error("GetCityById error! %v %v", v.Cityid, e)
		return
	}
	m["Provinceid"] = tm[0].(models.EhomeCity).Fatherid

	/*
		m["Province"], m["City"], m["Region"], err = models.GetRegionDetailById(m["Regionid"].(string))
		m["Headimageurl"], err = models.GetImageById(m["Headimageid"].(int))
		m["Certificateimagelist"], err = models.GetImageList(m["Certificateimage"].(string))
		m["Idcardimagelist"], err = models.GetImageList(m["Idcardimage"].(string))
		m["Toslist"], err = models.GetTosList(m["Tos"].(string))

		delete(m, "Password")
	*/

	return
}

// @param
// @Failure 403 body is empty
// @router /detail [post]
func (c *OrderController) Detail() {

	var id int
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

	id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! id[%s]", c.GetString("Id"))
		goto BOTTOM
	}

	map2["records"], err = GetOrderById(id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetOrderById error! %d %v\n", id, err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /audit [post]
func (c *OrderController) Audit() {
	map2 := make(map[string]interface{})

	var masterid int
	var adminid int
	var id int
	var audited int

	commented := c.GetString("Commented")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	adminid, err = models.GetAdminId(name)
	if err != nil {
		SetError(map2, PARAM_ERR, "admin user %s is invalid! %s", name, err)
		goto BOTTOM
	}

	id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id [%s] is invalid! %s", c.GetString("Id"), err)
		goto BOTTOM
	}

	masterid, err = c.GetInt("Masterid")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Masterid [%s] is invalid! %s", c.GetString("Masterid"), err)
		goto BOTTOM
	}

	audited, err = c.GetInt("Audited")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Audited [%s] is invalid! %s", c.GetString("Audited"), err)
		goto BOTTOM
	}

	err = models.AuditEhomeMaster(id, masterid, adminid, audited, commented)
	if err != nil {
		SetError(map2, DB_ERROR, "AuditEhomeMaster error!")
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /dealcash [post]
func (c *OrderController) DealCash() {
	map2 := make(map[string]interface{})

	var masterid int
	var adminid int
	var id int
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	adminid, err = models.GetAdminId(name)
	if err != nil {
		SetError(map2, PARAM_ERR, "admin user %s is invalid! %s", name, err)
		goto BOTTOM
	}

	id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id [%s] is invalid! %s", c.GetString("Id"), err)
		goto BOTTOM
	}

	masterid, err = c.GetInt("Masterid")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Masterid [%s] is invalid! %s", c.GetString("Masterid"), err)
		goto BOTTOM
	}

	err = models.DealCash(id, masterid, adminid)
	if err != nil {
		SetError(map2, DB_ERROR, "DealCash error! %s", err)
		goto BOTTOM
	}

	map2["status"] = 0

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
// @router /add [post]
func (c *OrderController) PlaceOrder() {
	beego.Info("input:", c.Input())

	map2 := make(map[string]interface{})

	mobile := c.GetString("Mobileno")
	imageid := c.GetString("ImageId")
	voiceid := c.GetString("VoiceId")
	attact := c.GetString("Attact")
	appointmenttime := c.GetString("AppintmentTime")
	Type := c.GetString("Ordertype")

	var labourid string
	var materialid string
	var fixtype string

	var TotalPrice float64

	var ehomeaddr *models.EhomeFixAddress
	var ehomeo models.EhomeOrder
	var userid int
	var tmp []string
	var ti int
	var orderid int64
	var addrid int

	var tmp_price float64

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	if Type == "1" {
		ehomeo.Type = 1
		ehomeo.Room, err = c.GetInt8("Room")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Room erro! [%s]", c.GetString("Room"))
			goto BOTTOM
		}
		ehomeo.Hall, err = c.GetInt8("Hall")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Hall erro! [%s]", c.GetString("Hall"))
			goto BOTTOM
		}
		ehomeo.Kitchen, err = c.GetInt8("Kitchen")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Kitchen erro! [%s]", c.GetString("Kitchen"))
			goto BOTTOM
		}
		ehomeo.Toilet, err = c.GetInt8("Toilet")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Toilet erro! [%s]", c.GetString("Toilet"))
			goto BOTTOM
		}

		ehomeo.Size, err = c.GetFloat("Size")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Size erro! [%s]", c.GetString("Size"))
			goto BOTTOM
		}

		ehomeo.Catidlist = c.GetString("StyleId")
		err = models.CheckStyleId(ehomeo.Catidlist)
		if err != nil {
			SetError(map2, PARAM_ERR, "Param StyleId error! [%s]", ehomeo.Catidlist)
			goto BOTTOM
		}

		ehomeo.Stylepriceid, err = c.GetInt("StylePriceId")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param StylepriceId erro! [%s]", c.GetString("StylepriceId"))
			goto BOTTOM
		}

	} else {
		ehomeo.Type = 0
		fixtype = c.GetString("FixType")
		labourid = c.GetString("LabourId")
		materialid = c.GetString("MaterialId")
		ehomeo.Catidlist, err = models.GetCatidlist(fixtype)
		if err != nil {
			SetError(map2, 5, "Param fixtype error! [%s]", fixtype)
			goto BOTTOM
		}
	}

	TotalPrice, err = c.GetFloat("TotalPrice")
	/*
		if err != nil {
			SetError(map2, PARAM_ERR, "TotalPrice error! [%s]", c.GetString("TotalPrice"))
			goto BOTTOM
		}
	*/
	tmp = strings.Split(imageid, ",")
	/*
		if len(tmp) < 1 {
			SetError(map2, PARAM_ERR, "param imageid error! [%s]", c.GetString("imageid"))
			goto BOTTOM
		}
	*/
	ti, err = strconv.Atoi(tmp[0])
	/*
		if err != nil {
			SetError(map2, PARAM_ERR, "param imageid error! [%s]", c.GetString("imageid"))
			goto BOTTOM
		}
	*/

	addrid, err = c.GetInt("AddrId")
	ehomeaddr, err = models.GetEhomeFixAddressById(addrid)
	if err != nil {
		SetError(map2, PARAM_ERR, "param Addrid error! [%s]", c.GetString("AddrId"))
		goto BOTTOM
	}

	userid, err = models.GetUseridByNo(mobile)
	if err != nil {
		SetError(map2, INVALID_USER, "user [%s] not register or not login! [%v]", mobile, err)
		goto BOTTOM
	}

	ehomeo.Remark = c.GetString("Remark")
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
	ehomeo.Appointmenttime, _ = time.Parse("2006-01-02 15:04:05", appointmenttime)
	beego.Error("mxz ", ehomeo.Appointmenttime.Format("2006-01-02 15:04:05"))
	ehomeo.Attact = attact
	ehomeo.Cityid = ehomeaddr.Cityid
	ehomeo.Region = ehomeaddr.Region
	ehomeo.Contactaddr = ehomeaddr.Contactaddr
	ehomeo.Contactname = ehomeaddr.Contactname
	ehomeo.Contactphone = ehomeaddr.Phone

	tmp_price, err = models.CalTotalPrice(Type, ehomeo.Stylepriceid, ehomeo.Size, labourid, materialid)
	if err != nil {
		SetError(map2, 15, "CalTotalPrice error!%v", err)
		goto BOTTOM
	}

	if !models.IsFloatEqual(tmp_price, ehomeo.Price) {
		SetError(map2, 6, "totalprice is not equal [%v] != [%v]", tmp_price, ehomeo.Price)
		beego.Error("not equal", tmp_price, ehomeo.Price)
		goto BOTTOM
	}

	orderid, err = models.AddEhomeOrder(&ehomeo)
	if err != nil {
		beego.Error("AddEhomeOrd error", err)
		SetError(map2, 7, "AddEhomeOrder error!%v", err)
		goto BOTTOM
	}

	err = models.InsertOrderPath(ehomeo.Orderno)
	if err != nil {
		SetError(map2, 8, "InsertOrderPath error!%v", err)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["Orderno"] = ehomeo.Orderno
	map2["Orderid"] = orderid
	map2["OrderTime"] = ehomeo.CreateTime.Format("20060102150405")
	map2["Price"] = TotalPrice

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /modify [post]
func (c *OrderController) Modify() {
	beego.Info("input:", c.Input())

	map2 := make(map[string]interface{})

	var area []interface{}

	Id, _ := c.GetInt("Id")
	mobile := c.GetString("Mobileno")
	imageid := c.GetString("ImageId")
	voiceid := c.GetString("VoiceId")
	attact := c.GetString("Attact")
	appointmenttime := c.GetString("AppintmentTime")
	Type := c.GetString("Ordertype")

	var labourid string
	var materialid string
	var fixtype string

	var TotalPrice float64
	//var addrid int

	//var ehomeaddr *models.EhomeFixAddress
	var tv *models.EhomeOrder
	var ehomeo models.EhomeOrder
	var userid int
	var tmp []string
	var ti int

	var tmp_price float64

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	tv, err = models.GetEhomeOrderById(Id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeOrderById error! [%d]! %v", Id, tv)
		goto BOTTOM
	}

	ehomeo = *tv

	ehomeo.Masterid, _ = c.GetInt("Masterid")
	ehomeo.Masterprice, _ = c.GetFloat("Masterprice")
	ehomeo.Status, _ = c.GetInt("Status")
	if ehomeo.Status < 0 || ehomeo.Status > 11 {
		SetError(map2, PARAM_ERR, "Param Status error! [%s]!valid range[0,11]", c.GetString("Status"))
		goto BOTTOM
	}

	if Type == "1" {
		ehomeo.Type = 1
		ehomeo.Room, err = c.GetInt8("Room")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Room erro! [%s]", c.GetString("Room"))
			goto BOTTOM
		}
		ehomeo.Hall, err = c.GetInt8("Hall")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Hall erro! [%s]", c.GetString("Hall"))
			goto BOTTOM
		}
		ehomeo.Kitchen, err = c.GetInt8("Kitchen")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Kitchen erro! [%s]", c.GetString("Kitchen"))
			goto BOTTOM
		}
		ehomeo.Toilet, err = c.GetInt8("Toilet")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Toilet erro! [%s]", c.GetString("Toilet"))
			goto BOTTOM
		}

		ehomeo.Size, err = c.GetFloat("Size")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param Size erro! [%s]", c.GetString("Size"))
			goto BOTTOM
		}

		ehomeo.Catidlist = c.GetString("StyleId")
		err = models.CheckStyleId(ehomeo.Catidlist)
		if err != nil {
			SetError(map2, PARAM_ERR, "Param StyleId error! [%s]", ehomeo.Catidlist)
			goto BOTTOM
		}

		ehomeo.Stylepriceid, err = c.GetInt("StylePriceId")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param StylepriceId erro! [%s]", c.GetString("StylepriceId"))
			goto BOTTOM
		}

	} else {
		ehomeo.Type = 0
		fixtype = c.GetString("FixType")
		labourid = c.GetString("LabourId")
		materialid = c.GetString("MaterialId")
		ehomeo.Catidlist, err = models.GetCatidlist(fixtype)
		if err != nil {
			SetError(map2, 5, "Param fixtype error! [%s]", fixtype)
			goto BOTTOM
		}
	}

	TotalPrice, err = c.GetFloat("TotalPrice")
	if err != nil {
		SetError(map2, PARAM_ERR, "TotalPrice error! [%s]", c.GetString("TotalPrice"))
		goto BOTTOM
	}
	tmp = strings.Split(imageid, ",")
	/*
		if len(tmp) < 1 {
			SetError(map2, PARAM_ERR, "param imageid error! [%s]", c.GetString("imageid"))
			goto BOTTOM
		}
	*/
	ti, err = strconv.Atoi(tmp[0])
	/*
		if err != nil {
			SetError(map2, PARAM_ERR, "param imageid error! [%s]", c.GetString("imageid"))
			goto BOTTOM
		}
	*/

	/*
		addrid, err = c.GetInt("AddrId")
		ehomeaddr, err = models.GetEhomeFixAddressById(addrid)
		if err != nil {
			SetError(map2, PARAM_ERR, "param Addrid error! [%s]", c.GetString("AddrId"))
			goto BOTTOM
		}
	*/

	userid, err = models.GetUseridByNo(mobile)
	if err != nil {
		SetError(map2, INVALID_USER, "user [%s] not register or not login! [%v]", mobile, err)
		goto BOTTOM
	}

	ehomeo.Remark = c.GetString("Remark")
	ehomeo.Imageid = ti
	//ehomeo.Orderno = models.GetOrderNo()
	ehomeo.Status = 1
	ehomeo.Imageidlist = imageid
	ehomeo.Voiceidlist = voiceid
	ehomeo.Labouridlist = labourid
	ehomeo.Materialidlist = materialid
	ehomeo.Userid = int(userid)
	ehomeo.Price = TotalPrice
	ehomeo.CreateTime = time.Now()

	ehomeo.Region = c.GetString("Regionid")
	ehomeo.Cityid = c.GetString("Cityid")
	ehomeo.Contactaddr = c.GetString("Contactaddress")
	ehomeo.Contactname = c.GetString("Contactname")
	ehomeo.Contactphone = c.GetString("Contactphone")

	if ehomeo.Region == "" {
		SetError(map2, PARAM_ERR, "param Regionid is empty")
		goto BOTTOM
	}
	area, err = models.GetRegionById(ehomeo.Region)
	if err != nil {
		SetError(map2, PARAM_ERR, "param Regionid is invalid! %s [%v]", ehomeo.Region, err)
		goto BOTTOM
	}
	ehomeo.Cityid = strconv.Itoa(area[0].(models.EhomeArea).Fatherid)

	/*
		if ehomeo.Cityid == "" {
			SetError(map2, PARAM_ERR, "param Cityid is empty")
			goto BOTTOM
		}
	*/

	if ehomeo.Contactaddr == "" {
		SetError(map2, PARAM_ERR, "param  Contactaddress is empty")
		goto BOTTOM
	}

	if ehomeo.Contactname == "" {
		SetError(map2, PARAM_ERR, "param  Contactname is empty")
		goto BOTTOM
	}

	if ehomeo.Contactphone == "" {
		SetError(map2, PARAM_ERR, "param  Contactphone is empty")
		goto BOTTOM
	}

	ehomeo.Appointmenttime, _ = time.Parse("2006-01-02 15:04:05", appointmenttime)
	beego.Error("mxz ", ehomeo.Appointmenttime.Format("2006-01-02 15:04:05"))
	ehomeo.Attact = attact

	tmp_price, err = models.CalTotalPrice(Type, ehomeo.Stylepriceid, ehomeo.Size, labourid, materialid)
	if err != nil {
		SetError(map2, 15, "CalTotalPrice error!%v", err)
		goto BOTTOM
	}

	if !models.IsFloatEqual(tmp_price, ehomeo.Price) {
		SetError(map2, 6, "totalprice is not equal [%v] != [%v]", tmp_price, ehomeo.Price)
		beego.Error("not equal", tmp_price, ehomeo.Price)
		goto BOTTOM
	}

	err = models.UpdateEhomeOrderById(&ehomeo)
	if err != nil {
		beego.Error("UpdateEhomeOrderById error", err)
		SetError(map2, 7, "UpdateEhomeOrderById error!%v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func Ordermasterlist(orderid string, name, phone string) (m []interface{}, err error) {
	var cond models.EhomeRequestOrder_Str
	cond.Orderid = orderid
	//cond.name = name
	//cond.phone = phone

	m, err = models.EhomeRequestOrder_list(0, 0, &cond)

	var tmp []interface{}
	for i := int(0); i < len(m); i++ {
		//m[i] = models.Struct2Map(m[i].(models.EhomeRequestOrder))
		v, e := models.GetEhomeMasterById(m[i].(map[string]interface{})["Masterid"].(int))
		if e != nil {
			err = e
			return
		}

		if phone != "" {
			if string(v.Phone[0:(len(phone))]) != phone {
				continue
			}
		}

		if name != "" {
			if v.Name != name {
				continue
			}
		}

		(m[i].(map[string]interface{}))["Cityid"] = v.Cityid
		(m[i].(map[string]interface{}))["Phone"] = v.Phone
		(m[i].(map[string]interface{}))["Regionid"] = v.Regionid
		(m[i].(map[string]interface{}))["Name"] = v.Name
		(m[i].(map[string]interface{}))["Address"] = v.Address
		(m[i].(map[string]interface{}))["Province"], (m[i].(map[string]interface{}))["City"], (m[i].(map[string]interface{}))["Region"], err = models.GetRegionDetailById(((m[i].(map[string]interface{}))["Regionid"]).(string))
		if err != nil {
			beego.Error("can't find region ", m[i].(map[string]interface{})["Regionid"], m[i].(map[string]interface{})["Masterid"])
			return
		}
		tmp = append(tmp, m[i])
	}
	m = tmp

	return
}

// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /requestmasterlist [post]
func (c *OrderController) Requestmasterlist() {
	map2 := make(map[string]interface{})

	mastername := c.GetString("Mastername")
	masterphone := c.GetString("Masterphone")
	orderid := c.GetString("Orderid")
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = Ordermasterlist(orderid, mastername, masterphone)
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /assignmaster [post]
func (c *OrderController) Assignmaster() {
	map2 := make(map[string]interface{})

	var mid int
	var v *models.EhomeMaster

	orderid := c.GetString("Orderid")
	masterid := c.GetString("Masterid")
	var oid int
	var orderno string

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	mid, _ = strconv.Atoi(masterid)
	v, err = models.GetEhomeMasterById(mid)
	if v == nil || err != nil {
		SetError(map2, PARAM_ERR, "masterid %s is invalid!", masterid)
		goto BOTTOM
	}

	oid, _ = strconv.Atoi(orderid)

	orderno, err = models.GetOrderNoByid(oid)
	if err != nil {
		SetError(map2, PARAM_ERR, "orderid [%s] not valid !%v", orderid, err)
		goto BOTTOM
	}

	err = models.UpdateOrderMaster(orderid, mid)
	if err != nil {
		SetError(map2, DB_ERROR, "updateOrderMaster error!%v", err)
		goto BOTTOM
	}

	err = models.UpdateOrderStatusLess(oid, orderno, models.ORDER_MASTER_ACCEPTED)
	if err != nil {
		SetError(map2, DB_ERROR, "error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /delete [post]
func (c *OrderController) Delete() {
	map2 := make(map[string]interface{})

	var o *models.EhomeOrder

	orderid := c.GetString("Orderid")
	var oid int

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	oid, err = strconv.Atoi(orderid)
	if err != nil {
		SetError(map2, PARAM_ERR, "orderid [%s] not valid !%v", orderid, err)
		goto BOTTOM
	}

	o, err = models.GetEhomeOrderById(oid)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeOrderById error!%d %v", oid, err)
		goto BOTTOM
	}

	err = models.DeleteEhomeOrder(oid)
	if err != nil {
		SetError(map2, DB_ERROR, "Delete ehome order [%d] error!%v", oid, err)
		goto BOTTOM
	}
	err = models.DeleteEhomeOrderPathByno(o.Orderno)

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object} models.order
// @Failure 403 body is empty
// @router /modifystatusandprice [post]
func (c *OrderController) Modifystatusandprice() {
	map2 := make(map[string]interface{})

	var o *models.EhomeOrder

	orderid := c.GetString("Orderid")
	var oid int

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	oid, err = strconv.Atoi(orderid)
	if err != nil {
		SetError(map2, PARAM_ERR, "orderid [%s] not valid !%v", orderid, err)
		goto BOTTOM
	}

	o, err = models.GetEhomeOrderById(oid)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeOrderById error!%d %v", oid, err)
		goto BOTTOM
	}

	if c.GetString("Status") != "" {
		o.Status, err = c.GetInt("Status")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param price [%s] not valid !%v", c.GetString("Status"), err)
			goto BOTTOM
		}
		if o.Status < 0 || o.Status > models.ORDER_MASTER_CONFIRM_PAY {
			SetError(map2, PARAM_ERR, "Param price [%s] not valid !it should range from 1 to 11", c.GetString("Status"))
			goto BOTTOM

		}
	}

	if c.GetString("Price") != "" {
		o.Price, err = c.GetFloat("Price")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param price [%s] not valid !%v", c.GetString("Price"), err)
			goto BOTTOM
		}
	}

	if c.GetString("Masterprice") != "" {
		o.Masterprice, err = c.GetFloat("Masterprice")
		if err != nil {
			SetError(map2, PARAM_ERR, "Param price [%s] not valid !%v", c.GetString("Masterprice"), err)
			goto BOTTOM
		}
	}

	err = models.UpdateEhomeOrderById(o)
	if err != nil {
		beego.Error("UpdateEhomeOrderById error", err)
		SetError(map2, 7, "UpdateEhomeOrderById error!%v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
