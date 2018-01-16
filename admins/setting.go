package admins

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type SettingController struct {
	beego.Controller
}

// URLMapping ...
func (c *SettingController) URLMapping() {
	c.Mapping("GetAdminlist", c.GetAdminlist)
	c.Mapping("ModifyAdmin", c.ModifyAdmin)
	c.Mapping("AddAdmin", c.AddAdmin)
	c.Mapping("ModifySetting", c.ModifySetting)
	c.Mapping("GetSetting", c.GetSetting)
	c.Mapping("DelAdmin", c.DelAdmin)
	c.Mapping("Test", c.Test)
}

func GetAllAdmin(limit, offset int64) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	fields = []string{"Id", "Name", "Phone", "Address", "Email"}

	m, err = models.GetAllEhomeAdmin(query, fields, sortby, order, offset, limit)

	if err != nil {
		return
	}

	if len(m) <= 0 {
		m = make([]interface{}, 0)
	}

	/*
		var tmp []interface{}

		for i := int(0); i < len(m); i++ {
			if name != "" && (m[i].(map[string]interface{}))["Name"].(string) != name {
				continue
			}

			(m[i].(map[string]interface{}))["Province"], (m[i].(map[string]interface{}))["City"], (m[i].(map[string]interface{}))["Region"], err = models.GetRegionDetailById(((m[i].(map[string]interface{}))["Regionid"]).(string))
			(m[i].(map[string]interface{}))["Headimageurl"], err = models.GetImageById(((m[i].(map[string]interface{}))["Headimageid"].(int)))
			tmp = append(tmp, m[i])
		}
		m = tmp
	*/

	return
}

// @param
// @Failure 403 body is empty
// @router /getadminlist [post]
func (c *SettingController) GetAdminlist() {
	map2 := make(map[string]interface{})
	//limit, offset := GetPage(c)

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllAdmin(0, 0)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllAdmin error![%v]", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modifyadmin [post]
func (c *SettingController) ModifyAdmin() {
	var v *models.EhomeAdmin
	map2 := make(map[string]interface{})

	id, e := c.GetInt("Id")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if e != nil {
		SetError(map2, PARAM_ERR, "param Id error! %v", e)
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	v, err = models.GetEhomeAdminById(id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeAdminById error! %v", err)
		goto BOTTOM
	}

	v.Name = c.GetString("AdminName")
	v.Email = c.GetString("Email")
	v.Phone = c.GetString("Phone")
	v.Address = c.GetString("Address")
	if c.GetString("Passwd") != "" {
		v.Passwd = c.GetString("Passwd")
	}

	err = models.UpdateEhomeAdminById(v)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeAdminById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /addadmin [post]
func (c *SettingController) AddAdmin() {

	var v models.EhomeAdmin
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

	v.Name = c.GetString("AdminName")
	v.Email = c.GetString("Email")
	v.Phone = c.GetString("Phone")
	v.Address = c.GetString("Address")
	v.Passwd = c.GetString("Passwd")

	_, err = models.AddEhomeAdmin(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeAdmin error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /deladmin [post]
func (c *SettingController) DelAdmin() {
	var v *models.EhomeAdmin
	map2 := make(map[string]interface{})

	id, e := c.GetInt("Id")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if e != nil {
		SetError(map2, PARAM_ERR, "param Id error! %v", e)
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	v, err = models.GetEhomeAdminById(id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeAdminById error! %v", err)
		goto BOTTOM
	}

	if v.Name == name {
		SetError(map2, PARAM_ERR, "Can't delete yourself!")
		goto BOTTOM
	}

	err = models.DeleteEhomeAdmin(v.Id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeAdmin error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /getsetting [post]
func (c *SettingController) GetSetting() {

	var v *models.EhomeSetting
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

	v, err = models.GetEhomeSettingById(1)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeAdminById error! %v", err)
		goto BOTTOM
	}

	map2["Emailcontent"] = v.Emailcontent

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modifysetting [post]
func (c *SettingController) ModifySetting() {
	var v *models.EhomeSetting
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

	v, err = models.GetEhomeSettingById(1)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeAdminById error! %v", err)
		goto BOTTOM
	}

	v.Emailcontent = c.GetString("Emailcontent")
	beego.Error("err", v)
	err = models.UpdateEhomeSettingById(v)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeSettingById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /test [post]
func (c *SettingController) Test() {
	map2 := make(map[string]interface{})
	err := models.SendMail("2", "r", "3", "4", 5, "8", "7", "8", "9", "10")
	if err != nil {
		SetError(map2, PARAM_ERR, "send mail error! %v", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
