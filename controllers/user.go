package controllers

import (
	"ehome/models"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("UploadImage", c.UploadImage)
	c.Mapping("Info", c.Info)
	c.Mapping("Changepass", c.Changepass)
}

// POST ...
// @Title UploadImage
// @Description  UploadImage
// @Param	Mobileno  Mobileno  string	true     "Mobileno"
// @Param	Usertype  Usertype  string	true     "Usertype"
// @Param	Token     Token     string	true     "Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Param   File      File      string  true     "File"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /uploadimage [post]
func (c *UserController) UploadImage() {

	beego.Error(c.Input())
	beego.Error(string(c.Ctx.Input.RequestBody))
	var map2 map[string]interface{}
	var userid int
	var filename string
	var path string
	var r []interface{}
	var sl []string
	var ehomeu models.EhomeUser
	var ehomei models.EhomeImage
	var id int64

	map2 = make(map[string]interface{})

	f, h, e := c.GetFile("File")
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if e == nil {
		f.Close()
	} else {
		SetError(map2, PARAM_ERR, "param File  not exists")
		goto BOTTOM
	}
	if err != nil {
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, PARAM_ERR, "user %s is invalid! %v", err)
		goto BOTTOM
	}

	sl = strings.Split(h.Filename, ".")
	if len(sl) < 2 {
		SetError(map2, INVALID_FILENAME, "invalid filename %s", h.Filename)
		goto BOTTOM
	}

	ehomei.Filename = models.GetRandFileName(models.ImagePath, sl)
	filename = ehomei.Filename
	path = models.ServerPath + filename
	beego.Info(path)

	e = c.SaveToFile("File", path)

	if e != nil {
		SetError(map2, SAVE_FILE_ERROR, "save file error")
		goto BOTTOM
	}

	if usertype == 1 {
		userid, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "invalid user %s", mobile)
			goto BOTTOM
		}

		r, err = models.GetEhomeUserByUid(userid)
		if err != nil {
			SetError(map2, GET_EHOME_USER_ERROR, "GetEhomeUserByUid error!")
			goto BOTTOM
		}
		ehomeu.Userid = userid
		ehomeu.Image = filename
		ehomeu.Balance = 0
		ehomeu.Points = 0
		if len(r) == 0 {
			_, err = models.AddEhomeUser(&ehomeu)
			if err != nil {
				SetError(map2, ADD_EHOME_USER_ERROR, "AddEhomeUser error！")
				goto BOTTOM
			}
		} else {
			ehomeu.Id = (r[0].(map[string]interface{}))["Id"].(int)
			ehomeu.Balance = (r[0].(map[string]interface{}))["Balance"].(int)
			ehomeu.Points = (r[0].(map[string]interface{}))["Points"].(int)
			err = models.UpdateEhomeUserById(&ehomeu)
			if err != nil {
				SetError(map2, UPDATE_EHOME_USER_ERROR, "UpdateEhomeUser error! %v", err)
				goto BOTTOM
			}
		}
	} else {
		id, err = models.AddEhomeImage(&ehomei)
		if err != nil {
			SetError(map2, DB_ERROR, "AddEhomeImage error! %v", err)
			goto BOTTOM
		}

		err = models.UpdateEhomeMasterHeadimage(mobile, int(id))
		if err != nil {
			SetError(map2, DB_ERROR, "UpdateEhomeMasterHeadimage error! %v", err)
			goto BOTTOM
		}
	}

	map2["status"] = 0
	map2["url"] = filename

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
	return
}

// Get ...
// @Title Get
// @Description  Info
// @Param	Mobileno  Mobileno  string	true     "Mobileno"
// @Param	Usertype  Usertype  string	true     "Usertype"
// @Param	Token     Token     string	true     "Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /info [get]
func (c *UserController) Info() {
	var map2 map[string]interface{}

	var m []interface{}

	map2 = make(map[string]interface{})
	var ecsu *models.EcsUsers

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	ecsu, err = models.ValidateUser2(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, INVALID_USER, "invalid user %s", mobile)
		goto BOTTOM
	}

	map2["status"] = 0
	map2["Username"] = ecsu.UserName
	m, err = models.GetEhomeUserByUid(ecsu.Id)
	if err == nil && len(m) > 0 {
		map2["Balance"] = m[0].(map[string]interface{})["Balance"].(int)
		map2["Points"] = m[0].(map[string]interface{})["Points"].(int)
		map2["Image"] = m[0].(map[string]interface{})["Image"].(string)
	} else {
		map2["Balance"] = 0
		map2["Points"] = 0
		map2["Image"] = ""
	}

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
	return
}

// Get ...
// @Title Get
// @Description  Info
// @Param	Mobileno  Mobileno  string	true     "Mobileno"
// @Param	Usertype  Usertype  string	true     "Usertype"
// @Param	Token     Token     string	true     "Token"
// @Param   Reqtime   Reqtime   string  true     "Reqtime"
// @Param   Oldpasswd  Oldpasswd string true     "Oldpasswd"
// @Param   Newpasswd  Newpasswd string true     "Newpasswd"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /changepasswd [get]
func (c *UserController) Changepass() {
	beego.Error(c.Input())
	map2 := make(map[string]interface{})

	var emas *models.EhomeMaster
	var ecsu *models.EcsUsers
	oldpasswd := c.GetString("Oldpasswd")
	newpasswd := c.GetString("Newpasswd")
	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if usertype == 1 {
		ecsu, err = models.ValidateUser2(mobile, token, reqtime, usertype)
		if err != nil {
			SetError(map2, INVALID_USER, "invalid user %s, %v", mobile, err)
			goto BOTTOM
		}

		if ecsu.Password != oldpasswd {
			SetError(map2, INVALID_PASS, "old Password invalid ")
			goto BOTTOM
		}

		if len(newpasswd) == 0 {
			SetError(map2, INVALID_PASS, "new Password ivalid")
			goto BOTTOM
		}

		ecsu.Password = newpasswd
		err = models.UpdateEcsUsersPasswd(ecsu)
		if err != nil {
			SetError(map2, DB_ERROR, "UpdateEcsUsersPasswd error %v", err)
			goto BOTTOM
		}
	} else {
		emas, err = models.ValidateMaster2(mobile, token, reqtime)
		if err != nil {
			SetError(map2, INVALID_USER, "invalid user %s", mobile)
			goto BOTTOM
		}

		if emas.Password != oldpasswd {
			SetError(map2, INVALID_PASS, "old Password invalid ")
			goto BOTTOM
		}

		if len(newpasswd) == 0 {
			SetError(map2, INVALID_PASS, "new Password ivalid")
			goto BOTTOM
		}

		emas.Password = newpasswd
		err = models.UpdateEhomeMasterPasswd(emas)
		if err != nil {
			SetError(map2, DB_ERROR, "UpdateEhomeMasterPasswd error %v", err)
			goto BOTTOM
		}
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
	return

}
