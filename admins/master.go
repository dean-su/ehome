package admins

import (
	"ehome/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MasterController struct {
	beego.Controller
}

// URLMapping ...
func (c *MasterController) URLMapping() {
	c.Mapping("GetMaster", c.GetMaster)
	c.Mapping("Detail", c.Detail)
	c.Mapping("GetRegion", c.GetRegion)
	c.Mapping("GetAuditPending", c.GetAuditPending)
	c.Mapping("Audit", c.Audit)
	c.Mapping("GetCashPending", c.GetCashPending)
	c.Mapping("DealCash", c.DealCash)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Modify", c.Modify)

}

func GetAllMaster(limit, offset int64, cityid, regionid string, name, phone string) (ml []interface{}, err error) {

	var cond models.EhomeMaster_Str
	cond.Cityid = cityid
	cond.Regionid = regionid
	cond.Phone = phone
	cond.Name = name

	ml, err = models.EhomeMaster_list(limit, offset, &cond)

	if len(ml) <= 0 {
		ml = make([]interface{}, 0)
	}

	var tmp []interface{}

	for i := int(0); i < len(ml); i++ {
		//ml[i] = models.Struct2Map(ml[i].(models.EhomeMaster))
		(ml[i].(map[string]interface{}))["Province"], (ml[i].(map[string]interface{}))["City"], (ml[i].(map[string]interface{}))["Region"], err = models.GetRegionDetailById(((ml[i].(map[string]interface{}))["Regionid"]).(string))
		(ml[i].(map[string]interface{}))["Headimageurl"], err = models.GetImageById(((ml[i].(map[string]interface{}))["Headimageid"].(int)))
		tmp = append(tmp, ml[i])
	}
	ml = tmp

	return
}

// @param
// @Failure 403 body is empty
// @router /get [post]
func (c *MasterController) GetMaster() {
	map2 := make(map[string]interface{})

	var m int64
	limit, offset := GetPage(c)

	mastername := c.GetString("Mastername")
	masterphone := c.GetString("Masterphone")
	regionid := c.GetString("Regionid")
	cityid := c.GetString("Cityid")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	map2["records"], err = GetAllMaster(limit, offset, cityid, regionid, mastername, masterphone)

	if err != nil {
		SetError(map2, DB_ERROR, "GetAllMaster error! %s", err)
		delete(map2, "records")
		goto BOTTOM
	}

	m, err = models.GetMasterNum(cityid, regionid, mastername, masterphone)

	if err != nil {
		SetError(map2, DB_ERROR, "GetMasterNum error!")
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

func GetMasterById(id int) (m map[string]interface{}, err error) {
	var v *models.EhomeMaster
	v, err = models.GetEhomeMasterById(id)
	if err != nil {
		return
	}

	m = models.Struct2Map(*v)
	m["Province"], m["City"], m["Region"], err = models.GetRegionDetailById(m["Regionid"].(string))
	m["Headimageurl"], err = models.GetImageById(m["Headimageid"].(int))
	m["Certificateimagelist"], err = models.GetImageList(m["Certificateimage"].(string))
	m["Idcardimagelist"], err = models.GetImageList(m["Idcardimage"].(string))
	m["Toslist"], err = models.GetTosList(m["Tos"].(string))
	k, e := models.GetCityById(m["Cityid"].(string))
	if e != nil {
		err = e
		beego.Error("GetCityById error! %v", err)
		return
	}

	m["Provinceid"] = strconv.Itoa(k[0].(models.EhomeCity).Fatherid)

	delete(m, "Password")

	return
}

// @param
// @Failure 403 body is empty
// @router /detail [post]
func (c *MasterController) Detail() {

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

	map2["records"], err = GetMasterById(id)
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /region [post]
func (c *MasterController) GetRegion() {
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

	map2["records"], err = models.GetAdminRegion()
	if err != nil {
		SetError(map2, DB_ERROR, "GetAdminRegion error! %v", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAllAuditPending(limit, offset int64) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	fields = []string{"Id", "CreateTime", "Masterid"}
	query["DealFlag"] = "0"

	m, err = models.GetAllEhomeMasterAuditPending(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(m) <= 0 {
		m = make([]interface{}, 0)
	}

	for i := int(0); i < len(m); i++ {
		var v *models.EhomeMaster

		//m[i] = models.Struct2Map(m[i].(models.EhomeMasterAuditPending))

		v, err = models.GetEhomeMasterById((m[i].(map[string]interface{}))["Masterid"].(int))
		(m[i].(map[string]interface{}))["Mastername"] = v.Name
		(m[i].(map[string]interface{}))["Masterphone"] = v.Phone
		(m[i].(map[string]interface{}))["Headimageurl"], err = models.GetImageById(v.Headimageid)
	}

	return
}

// @param
// @Failure 403 body is empty
// @router /auditpending [post]
func (c *MasterController) GetAuditPending() {
	map2 := make(map[string]interface{})

	beego.Error("ok ", c.GetString("Name"), "ab")
	limit, offset := GetPage(c)
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllAuditPending(limit, offset)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllAuditPending error! %v", err)
		goto BOTTOM
	}
	map2["Total"], err = models.GetAuditPendingNum()
	if err != nil {
		SetError(map2, DB_ERROR, "GetAuditPendingNum error! %v", err)
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
func (c *MasterController) Audit() {
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

func GetAllCashPending(limit, offset int64) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	fields = []string{"Id", "CreateTime", "Masterid", "Amount", "Bank", "Branch"}
	query["DealFlag"] = "0"

	m, err = models.GetAllEhomeMasterCashLog(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(m) <= 0 {
		m = make([]interface{}, 0)
	}

	for i := int(0); i < len(m); i++ {
		var v *models.EhomeMaster

		//m[i] = models.Struct2Map(m[i].(models.EhomeMasterAuditPending))

		v, err = models.GetEhomeMasterById((m[i].(map[string]interface{}))["Masterid"].(int))
		if err != nil {
			continue
		}
		(m[i].(map[string]interface{}))["Mastername"] = v.Name
		(m[i].(map[string]interface{}))["Masterphone"] = v.Phone
		(m[i].(map[string]interface{}))["Headimageurl"], err = models.GetImageById(v.Headimageid)
	}

	return
}

// @param
// @Failure 403 body is empty
// @router /cashpending [post]
func (c *MasterController) GetCashPending() {
	map2 := make(map[string]interface{})

	limit, offset := GetPage(c)
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllCashPending(limit, offset)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllCashPending error! %v", err)
		goto BOTTOM
	}
	map2["Total"], err = models.GetCashPendingNum()
	if err != nil {
		SetError(map2, DB_ERROR, "GetCashPendingNum error! %v", err)
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
func (c *MasterController) DealCash() {
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

// @param
// @Failure 403 body is empty
// @router /delete [post]
func (c *MasterController) Delete() {
	map2 := make(map[string]interface{})

	var num int64
	//var v *models.EhomeMaster
	id, e := c.GetInt("Id")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}
	if e != nil {
		SetError(map2, PARAM_ERR, "Param Id error! %v, %s", e, c.GetString("Id"))
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	_, err = models.GetEhomeMasterById(id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllEhomeMast error! %v", err)
		goto BOTTOM
	}
	//beego.Error(v)

	num, err = models.GetEhomeMasterOrderNum(id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeMasterOrderNum error! %v [%d]", err, id)
		goto BOTTOM
	}
	if num != 0 {
		SetError(map2, MASTER_HAVE_ORDER, "This master[%d] have order [%d]", id, num)
		goto BOTTOM
	}

	err = models.BakupEhomeMaster(id)
	if err != nil {
		SetError(map2, DB_ERROR, "BakupEhomeMaster error! %v", err)
		goto BOTTOM
	}

	err = models.DeleteEhomeMaster(id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeMaster error! %v [%d]", err, id)
		goto BOTTOM
	}

	err = models.DeleteEhomeMasterAuditPendingbymasterid(id)
	if err != nil {
		beego.Error("DeleteEhomeMasterAuditPendingbymasterid error! %v", err)
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()

}

// @param
// @Failure 403 body is empty
// @router /modify [post]
func (c *MasterController) Modify() {
	map2 := make(map[string]interface{})
	var area []interface{}
	var v *models.EhomeMaster

	id, e := c.GetInt("Id")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if e != nil {
		SetError(map2, PARAM_ERR, "Param Id error! %v, %s", e, c.GetString("Id"))
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	v, err = models.GetEhomeMasterById(id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetEhomeMasterById error! %v", err)
		goto BOTTOM
	}

	v.Phone = c.GetString("Masterphone")
	v.Name = c.GetString("Mastername")
	v.Idcard = c.GetString("Idcard")
	v.Tos = c.GetString("Tos")
	v.Regionid = c.GetString("Regionid")
	area, err = models.GetRegionById(v.Regionid)
	if err != nil {
		SetError(map2, PARAM_ERR, "param Regionid error! %v", err)
		goto BOTTOM
	}
	v.Cityid = strconv.Itoa(area[0].(models.EhomeArea).Fatherid)

	v.Address = c.GetString("Address")
	v.Idcardimage = c.GetString("Idcardimage")
	v.Certificateimage = c.GetString("Certificateimage")
	v.Headimageid, err = c.GetInt("Headimageid")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Headimageid error! %s", c.GetString("Headimageid"))
		goto BOTTOM
	}
	v.Balance, err = c.GetFloat("Balance")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Balance error! %s", c.GetString("Balance"))
		goto BOTTOM
	}
	v.Audited, err = c.GetInt16("Audited")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Audited error!%s", c.GetString("Audited"))
		goto BOTTOM
	}

	if c.GetString("Passwd") != "" {
		v.Password = c.GetString("Passwd")
	}

	err = models.UpdateEhomeMasterById(v)
	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeMasterById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
