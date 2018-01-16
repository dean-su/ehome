package admins

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("GetUser", c.GetUser)
	c.Mapping("Infobyno", c.Infobyno)
	c.Mapping("Register", c.Register)
	c.Mapping("Modify", c.Modify)

}

func GetAllUser(limit, offset int64, cond *models.EcsUsers_Str) (ml []interface{}, err error) {

	ml, err = models.EcsUsers_list(limit, offset, cond)

	if len(ml) <= 0 {
		ml = make([]interface{}, 0)
	}

	return
}

// @param
// @Failure 403 body is empty
// @router /list [post]
func (c *UserController) GetUser() {
	map2 := make(map[string]interface{})
	var cond models.EcsUsers_Str

	var m int64
	limit, offset := GetPage(c)

	cond.UserName = c.GetString("Username")
	cond.MobilePhone = c.GetString("Userphone")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	map2["records"], err = GetAllUser(limit, offset, &cond)

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllUser error! %s", err)
		delete(map2, "records")
		goto BOTTOM
	}

	m, err = models.EcsUsers_num(&cond)

	if err != nil {
		SetError(map2, DB_ERROR, "GetUserNum error!")
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

// @param
// @Failure 403 body is empty
// @router /infobyno [post]
func (c *UserController) Infobyno() {
	map2 := make(map[string]interface{})
	var cond models.EcsUsers_Str

	var m []interface{}

	cond.MobilePhone = c.GetString("Mobileno")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	m, err = GetAllUser(0, 0, &cond)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllUser error!%v", err)
		goto BOTTOM
	}

	if len(m) == 0 {
		SetError(map2, PARAM_ERR, "param mobileno [%s] error! no this user", cond.MobilePhone)
		goto BOTTOM
	}
	map2 = m[0].(map[string]interface{})

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /register [post]
func (c *UserController) Register() {
	map2 := make(map[string]interface{})

	var v models.EcsUsers

	v.UserName = c.GetString("Username")
	v.MobilePhone = c.GetString("Mobileno")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	if len(v.MobilePhone) != 11 {
		SetError(map2, MOBILE_LEN_ERR, "param Mobileno %s is invalid", v.MobilePhone)
		goto BOTTOM
	}

	v.Selfreg = 0

	_, err = models.AddEcsUsers(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEcsUsers error!%v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modify [post]
func (c *UserController) Modify() {
	map2 := make(map[string]interface{})

	var e error
	var v models.EcsUsers

	v.Id, e = c.GetInt("Id")
	v.UserName = c.GetString("Username")
	v.MobilePhone = c.GetString("Mobileno")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if e != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]\n", c.GetString("Id"))
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	if len(v.MobilePhone) != 11 {
		SetError(map2, MOBILE_LEN_ERR, "param Mobileno %s is invalid", v.MobilePhone)
		goto BOTTOM
	}

	v.Selfreg, err = c.GetInt8("Selfreg")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Selfreg error! [%s]\n", c.GetString("Selfreg"))
		goto BOTTOM
	}

	err = models.UpdateEcsUsersById(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEcsUsersById error!%v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
