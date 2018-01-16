package controllers

import (
	"ehome/db"
	"ehome/models"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

// MasterController oprations for EcsUsers
type MasterController struct {
	beego.Controller
}

// URLMapping ...
func (c *MasterController) URLMapping() {
	c.Mapping("RegisterInit", c.RegisterInit)
	c.Mapping("RegisterRequest", c.RegisterRequest)
}

// Get ...
// @Title Init
// @Description  init
// @Param	phone  phone string	false	"Filter ..."
// @Success 201 {int}
// @Failure 403 body is empty
// @router /init [get]
func (c *MasterController) RegisterInit() {

	mobile := c.GetString("phone")
	v, err := models.GetEhomeMasterByPhone(mobile)
	beego.Info("mobile :2", v, mobile)
	map2 := make(map[string]interface{})

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndnum := rnd.Int() % 10000

	content := fmt.Sprintf("您的验证码是：%d。请不要把验证码泄露给其他人。", rndnum)

	var Red redis.Conn

	if len(mobile) != 11 {
		SetError(map2, MOBILE_LEN_ERR, "Mobileno  not exist or len error!")
		goto BOTTOM
	}

	if len(v) > 0 {
		SetError(map2, USER_EXIST, "mobile has bee register!")
		goto BOTTOM
	}

	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeMasterByPhone error! %v", err)
		goto BOTTOM
	}

	err = models.SendSms(mobile, content)

	if err != nil {
		beego.Error("SendSms error !")
		SetError(map2, SEND_SMS_ERR, "Send sms error!")
		goto BOTTOM
	} else {
		Red = db.RedisClient.Get()
		defer Red.Close()
		_, err = Red.Do("SET", "REGMASTER_"+mobile, strconv.Itoa(rndnum))
		if err != nil {
			SetError(map2, REDIS_DO_ERR, "Red set REGMASTER err")
			goto BOTTOM
		} else {
			Red.Do("expire", "REGMASTER_"+mobile, 600)
		}

		map2["status"] = 0
		map2["Cryptid"] = "123456"
		c.Data["json"] = map2
	}

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Get ...
// @Title request
// @Description  request
// @Param	Mobileno  Phone	string	false	"Filter ..."
// @Param	UserType  usertype	string	false	"Filter ..."
// @Param	Identificationcode  Identificationcode string	false	"Filter ..."
// @Param	Passwd Passwd  string	false	"Filter ..."
// @Param	Username  Username string	false	"Filter ..."
// @Success 201 {int}
// @Failure 403 body is empty
// @router /request [get]
func (c *MasterController) RegisterRequest() {
	var v models.EhomeMaster
	var vv models.EhomeMasterAuditPending

	v.Phone = c.GetString("phone")
	v.Password = c.GetString("password")
	v.Name = c.GetString("name")
	v.Idcard = c.GetString("idcard")
	v.Tos = c.GetString("tos")
	v.Regionid = c.GetString("region")
	v.Address = c.GetString("address")
	v.Idcardimage = c.GetString("idcardimage")
	v.Certificateimage = c.GetString("certificateimage")

	Identificationcode, err := c.GetInt("identificationcode")
	map2 := make(map[string]interface{})

	var Red redis.Conn
	var rv interface{}
	var ar []interface{}
	//var rv []interface{}
	var tmp int
	var masterid int64
	var area []interface{}

	if err != nil {
		SetError(map2, ID_CODE_ERR, "Identificationcode error!")
		goto BOTTOM
	}

	Red = db.RedisClient.Get()
	defer Red.Close()

	rv, err = Red.Do("GET", "REGMASTER_"+v.Phone)
	if err != nil {
		SetError(map2, REDIS_DO_ERR, "redis GET REGMASTER_ err!")
		goto BOTTOM
	}

	if rv != nil {
		tmp, _ = strconv.Atoi(string(rv.([]byte)))
		if Identificationcode != tmp {
			SetError(map2, ID_CODE_ERR, "Identificationcode error!")
			goto BOTTOM
		}

		beego.Info("test [", string(rv.([]byte)), "]")
	}

	ar, _ = models.GetEhomeMasterByPhone(v.Phone)
	beego.Info("mobile :2", v.Phone)

	if len(ar) > 0 {
		SetError(map2, USER_EXIST, "user exists!")
		goto BOTTOM
	}

	v.Headimageid, _ = c.GetInt("headimageid")

	area, err = models.GetRegionById(v.Regionid)
	if err != nil {
		SetError(map2, PARAM_ERR, "param Region error! %v", err)
		goto BOTTOM
	}
	v.Cityid = strconv.Itoa(area[0].(models.EhomeArea).Fatherid)

	if masterid, err = models.AddEhomeMaster(&v); err != nil {
		SetError(map2, DB_ERROR, "AddEhomeMaster error %s", err)
		goto BOTTOM
	}

	vv.Masterid = int(masterid)
	if _, err = models.AddEhomeMasterAuditPending(&vv); err != nil {
		SetError(map2, DB_ERROR, "AddEhomeMasterAuditPending error %", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
