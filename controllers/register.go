package controllers

import (
	"crypto/md5"
	"ehome/db"
	"ehome/models"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// EcsUsersController oprations for EcsUsers
type EcsUsersController struct {
	beego.Controller
}

// URLMapping ...
func (c *EcsUsersController) URLMapping() {
	c.Mapping("RegisterInit", c.RegisterInit)
	c.Mapping("RegisterRequest", c.RegisterRequest)
	c.Mapping("ForgetInit", c.ForgetInit)
	c.Mapping("ForgetRequest", c.ForgetRequest)
	c.Mapping("SendSms", c.SendSms)
}

// Get ...
// @Title Init
// @Description  init
// @Param	Mobileno  Phone	string	false	"Filter ..."
// @Param	UserType  usertype	string	false	"Filter ..."
// @Success 201 {int} models.EcsUsers
// @Failure 403 body is empty
// @router /sendsms [get]
func (c *EcsUsersController) SendSms() {

	mobile := c.GetString("Mobileno")
	map2 := make(map[string]interface{})

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rndnum := rnd.Int() % 10000

	curtime := time.Now().Format("20060102150405")
	str := string("易居乐筑yjlz789") + curtime

	sign := md5.Sum([]byte(str))
	signstr := hex.EncodeToString(sign[0:16])
	beego.Info(str, signstr)
	content := fmt.Sprintf("您的验证码是：%d。请不要把验证码泄露给其他人。", rndnum)

	if len(mobile) != 11 {
		map2["Status"] = 1
	} else {
		e := models.SendSms(mobile, content)

		if e != nil {
			beego.Error("PostForm error !")
			map2["Status"] = 2
			map2["Errmsg"] = fmt.Sprintf("%v", e)
		} else {
			map2["Status"] = 0
			map2["Code"] = rndnum
		}
	}

	c.Data["json"] = map2
	c.ServeJSON()
}

// Get ...
// @Title Init
// @Description  init
// @Param	Mobileno  Phone	string	false	"Filter ..."
// @Param	UserType  usertype	string	false	"Filter ..."
// @Success 201 {int} models.EcsUsers
// @Failure 403 body is empty
// @router /init [get]
func (c *EcsUsersController) RegisterInit() {

	mobile := c.GetString("Mobileno")
	//usrtype := c.GetString("UserType")
	v, _ := models.GetEcsUsersByPhone(mobile)
	map2 := make(map[string]interface{})

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rndnum := rnd.Int() % 10000

	curtime := time.Now().Format("20060102150405")
	str := string("易居乐筑yjlz789") + curtime

	sign := md5.Sum([]byte(str))
	signstr := hex.EncodeToString(sign[0:16])
	beego.Info(str, signstr)
	content := fmt.Sprintf("您的验证码是：%d。请不要把验证码泄露给其他人。", rndnum)
	//content := fmt.Sprintf("验证码%d【易居乐筑】", rndnum)

	if v > 0 {
		map2["Status"] = 1
	} else if len(mobile) != 11 {
		map2["Status"] = 11
	} else {
		e := models.SendSms(mobile, content)

		if e != nil {
			beego.Error("PostForm error !")
			map2["Status"] = 2
		} else {
			redis := db.RedisClient.Get()
			defer redis.Close()
			_, rerr := redis.Do("SET", "REGIDENT_"+mobile, strconv.Itoa(rndnum))
			if rerr != nil {
				map2["Status"] = 3
				beego.Error("redis DO err", rerr)
				goto END
			} else {
				redis.Do("expire", "REGIDENT_"+mobile, 600)
			}

			map2["Status"] = 0
			map2["Cryptid"] = "123456"
		}
	}

END:
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
func (c *EcsUsersController) RegisterRequest() {
	/*
		Mobileno=手机号码
		Usertype=用户类型（1：普通用户，　2: 师傅）
		Identificationcode=短信验证码
		Passwd=加密后的密码（先以初始流程返回的Cryptid为密钥用aes算法加密，再用base64编码）
		Username=用名
	*/
	var v models.EcsUsers
	v.UserName = c.GetString("Username")
	v.MobilePhone = c.GetString("Mobileno")
	v.Password = c.GetString("Passwd")
	Identificationcode, e := c.GetInt("Identificationcode")
	map2 := make(map[string]interface{})

	if e != nil {
		beego.Info("Identificationcode error")
		map2["Status"] = 3
	} else {
		redis := db.RedisClient.Get()
		defer redis.Close()
		rv, rerr := redis.Do("GET", "REGIDENT_"+v.MobilePhone)
		if rerr != nil {
			map2["Status"] = 1
		} else {

			if rv != nil {
				iden, _ := strconv.Atoi(string(rv.([]byte)))
				if Identificationcode != iden {
					map2["Status"] = 5
					goto BOTTOM
				}
			} else {
				map2["Status"] = 11
				goto BOTTOM
			}

			beego.Info("Identificationcode ", Identificationcode)
			//usrtype := c.GetString("UserType")
			val, _ := models.GetEcsUsersByPhone(v.MobilePhone)
			beego.Info("mobile :2", val, v.MobilePhone)

			if val > 0 {
				map2["Status"] = 2
			} else {

				beego.Info(v)
				if _, err := models.AddEcsUsers(&v); err == nil {
					map2["Status"] = 0
				} else {
					beego.Info("AddEcsUser error", err)
					map2["Status"] = 4
				}
			}
		}
	}

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Get ...
// @Title Init
// @Description  init
// @Param	Mobileno  Phone	string	false	"Filter ..."
// @Param	UserType  usertype	string	false	"Filter ..."
// @Success 201 {int}
// @Failure 403 body is empty
// @router /forgetinit [get]
func (c *EcsUsersController) ForgetInit() {

	mobile := c.GetString("Mobileno")
	usertype := c.GetString("Usertype")
	t, _ := strconv.Atoi(usertype)

	//usrtype := c.GetString("UserType")
	beego.Info("mobile :", mobile)
	id, _ := models.GetIdByNo(mobile, t)
	beego.Info("mobile :2", t, mobile)
	map2 := make(map[string]interface{})

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rndnum := rnd.Int() % 10000

	content := fmt.Sprintf("您的验证码是：%d。请不要把验证码泄露给其他人。", rndnum)
	//content := fmt.Sprintf("验证码%d【易居乐筑】", rndnum)

	if id <= 0 || len(mobile) != 11 {
		map2["Status"] = 1
		c.Data["json"] = map2
	} else if usertype != "1" && usertype != "2" {
		map2["Status"] = PARAM_ERR
		c.Data["json"] = map2
	} else {
		e := models.SendSms(mobile, content)

		if e != nil {
			beego.Error("PostForm error !")
			map2["Status"] = 2
			c.Data["json"] = map2
		} else {
			redis := db.RedisClient.Get()
			defer redis.Close()
			_, rerr := redis.Do("SET", "FORGET_"+usertype+mobile, strconv.Itoa(rndnum))
			if rerr != nil {
				map2["Status"] = 3
				beego.Error("redis DO err", rerr)
				goto END
			} else {
				redis.Do("expire", "FORGET_"+usertype+mobile, 600)
			}

			map2["Status"] = 0
			map2["Cryptid"] = "123456"
			c.Data["json"] = map2
		}
	}

END:
	c.ServeJSON()
}

// Get ...
// @Title request
// @Description  request
// @Param	Mobileno  Phone	string	false	"Filter ..."
// @Param	UserType  usertype	string	false	"Filter ..."
// @Param	Identificationcode  Identificationcode string	false	"Filter ..."
// @Param	Passwd Passwd  string	false	"Filter ..."
// @Success 201 {int} models.EcsUsers
// @Failure 403 body is empty
// @router /forgetrequest [get]
func (c *EcsUsersController) ForgetRequest() {
	/*
		Mobileno=手机号码
		Usertype=用户类型（1：普通用户，　2: 师傅）
		Identificationcode=短信验证码
		Passwd=加密后的密码（先以初始流程返回的Cryptid为密钥用aes算法加密，再用base64编码）
	*/
	var v models.EcsUsers
	v.MobilePhone = c.GetString("Mobileno")
	v.Password = c.GetString("Passwd")
	usertype := c.GetString("Usertype")
	Identificationcode, e := c.GetInt("Identificationcode")
	map2 := make(map[string]interface{})

	if e != nil {
		beego.Info("Identificationcode error")
		map2["Status"] = 3
	} else {
		redis := db.RedisClient.Get()
		defer redis.Close()
		rv, rerr := redis.Do("GET", "FORGET_"+usertype+v.MobilePhone)
		if rerr != nil {
			map2["Status"] = 1
		} else {

			if rv != nil {
				iden, _ := strconv.Atoi(string(rv.([]byte)))
				if Identificationcode != iden {
					map2["Status"] = 5
					goto BOTTOM
				}

			} else {
				map2["Status"] = 11
				map2["errmsg"] = "Identification not exist!"
				goto BOTTOM
			}

			if usertype == "1" {
				v.Id, rerr = models.GetUseridByNo(v.MobilePhone)

				if rerr != nil {
					map2["Status"] = 2
				} else {

					if err := models.UpdateEcsUsersById(&v); err == nil {
						map2["Status"] = 0
					} else {
						beego.Info("UpdateEcsUsersById error", err)
						map2["Status"] = 4
					}
				}
			} else {
				m, rerr := models.GetEhomeMasterByPhone(v.MobilePhone)
				if rerr != nil || len(m) == 0 {
					map2["Status"] = 2
				} else {
					tmp := (models.EhomeMaster((m[0].(models.EhomeMaster))))
					tmp.Password = c.GetString("Passwd")
					if err := models.UpdateEhomeMasterById(&tmp); err == nil {
						map2["Status"] = 0
					} else {
						map2["Status"] = 4
					}
				}
			}
		}
	}
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
