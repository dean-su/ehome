package admins

import (
	"ehome/models"
	"fmt"
	"github.com/astaxie/beego"
)

type AdminLoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *AdminLoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}

// @param
// @Failure 403 body is empty
// @router / [get]
func (c *AdminLoginController) Login() {
	name := c.GetString("Name")
	passwd := c.GetString("Password")
	map2 := make(map[string]interface{})

	m, err := GetAdminByName(name)
	if err != nil {
		SetError(map2, PARAM_ERR, "GetAdminByName error! %v", err)
		goto BOTTOM
	}
	if passwd != (m[0].(models.EhomeAdmin)).Passwd {
		SetError(map2, INVALID_PASS, "password is invalid!")
		goto BOTTOM
	}

	map2["status"] = 0
	map2["Token"] = models.AdminToken(name)

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAdminByName(name string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	query["Name"] = name

	limit = 1

	ml, err = models.GetAllEhomeAdmin(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No Admin user %s", name)
	}

	return
}
