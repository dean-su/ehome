package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

// LogoutController oprations for EhomeTopic
type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) URLMapping() {
	c.Mapping("Logout", c.Logout)
}

// Get ...
// @Title Get
// @Description Logout
// @Success 201 {int}
// @Failure 403 body is empty
// @router / [get]
func (c *LogoutController) Logout() {
	map2 := make(map[string]interface{})

	mobile, token, _, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	if usertype == 1 {
		models.UserLogout(mobile, token)
	} else {
		models.MasterLogout(mobile, token)
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
