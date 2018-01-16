package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

// LoginController oprations for EhomeTopic
type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}

// Get ...
// @Title Get
// @Description create EhomeTopic
// @Param	Mobileno  Mobileno  string true		"tel no"
// @Param   Usertype  Usertype  string true     "user type"
// @Param   Passwd    Passwd    string true      "passwd"
// @Param   Source    Source    string true     "Source"
// @Success 201 {int}
// @Failure 403 body is empty
// @router / [get]
func (c *LoginController) Login() {
	beego.Error(c.Input())
	mobile := c.GetString("Mobileno")
	usertype := c.GetString("Usertype")
	map2 := make(map[string]interface{})
	pass := c.GetString("Passwd")

	if usertype == "1" {
		user, err := models.GetUserByNo(mobile)

		if err != nil {
			beego.Error("GetUserByNo", err)
			map2["Status"] = 1
		} else {
			dbpasswd := user.Password

			beego.Info("db pass", dbpasswd, "input pass", pass)
			if pass != dbpasswd {
				map2["Status"] = 2
			} else {
				map2["Status"] = 0

				map2["Token"] = models.UserToken(mobile)
				map2["Mobileno"] = mobile
			}
		}
	} else {
		sl, err := models.GetEhomeMasterByPhone(mobile)
		if err != nil || len(sl) == 0 {
			map2["Status"] = 1
		} else {
			if pass != sl[0].(models.EhomeMaster).Password {
				map2["Status"] = 2
				map2["errmsg"] = "password not correct!"
			} else {
				map2["Status"] = 0

				map2["Token"] = models.MasterToken(mobile)
				map2["Mobileno"] = mobile
			}
		}
	}

	c.Data["json"] = map2
	c.ServeJSON()
}
