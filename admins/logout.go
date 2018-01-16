package admins

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type AdminLogoutController struct {
	beego.Controller
}

// URLMapping ...
func (c *AdminLogoutController) URLMapping() {
	c.Mapping("Logout", c.Logout)
}

// @param
// @Failure 403 body is empty
// @router / [post]
func (c *AdminLogoutController) Logout() {

	map2 := make(map[string]interface{})
	name, token, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	models.AdminLogout(name, token)

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
