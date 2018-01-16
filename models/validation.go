package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

func GetIdByNo(mobile string, usertype int) (userid int, err error) {
	if usertype == 1 {
		userid, err = GetUseridByNo(mobile)
	} else {
		var v *EhomeMaster
		v, err = GetMasterByNo(mobile)
		if err != nil {
			return
		}
		userid = v.Id
	}
	return
}

func GetUserByNo(mobile string) (v *EcsUsers, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	var ehomeu EcsUsers

	fields = append(fields, "Id", "UserName", "Password")
	query["MobilePhone"] = mobile

	var ml []interface{}

	ml, err = GetAllEcsUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No this user %s", mobile)
	} else {
		ehomeu.Id = ((ml[0].(map[string]interface{}))["Id"]).(int)
		ehomeu.UserName = ((ml[0].(map[string]interface{}))["UserName"]).(string)
		ehomeu.Password = ((ml[0].(map[string]interface{}))["Password"]).(string)
		v = &ehomeu
	}

	return
}

func GetUseridByNo(mobile string) (userid int, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id")
	query["MobilePhone"] = mobile

	var ml []interface{}

	ml, err = GetAllEcsUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No this user %s", mobile)
	} else {
		userid = ((ml[0].(map[string]interface{}))["Id"]).(int)
	}

	return
}

func GetMasterByNo(mobileno string) (v *EhomeMaster, err error) {
	var m []interface{}
	m, err = GetEhomeMasterByPhone(mobileno)
	if len(m) == 0 && err == nil {
		err = fmt.Errorf("Master %s not exist!", mobileno)
		return
	}
	var a EhomeMaster
	a = m[0].(EhomeMaster)
	v = &a

	return
}

func ValidateUser2(mobileno string, token string, reqtime int64, usertype int) (v *EcsUsers, err error) {
	v, err = GetUserByNo(mobileno)
	if err != nil {
		return
	}

	err = CheckUserToken(mobileno, token, reqtime)

	return
}

func ValidateUser(mobileno string, token string, reqtime int64) (uid int, err error) {
	uid, err = GetUseridByNo(mobileno)
	if err != nil {
		return
	}

	err = CheckUserToken(mobileno, token, reqtime)
	return
}

func ValidateMaster(mobileno string, token string, reqtime int64) (uid int, err error) {
	var m []interface{}
	m, err = GetEhomeMasterByPhone(mobileno)
	if len(m) == 0 && err == nil {
		err = fmt.Errorf("Master %s not exist!", mobileno)
		return
	}
	uid = (m[0].(EhomeMaster)).Id

	err = CheckMasterToken(mobileno, token, reqtime)

	return
}

func ValidateMaster2(mobileno string, token string, reqtime int64) (v *EhomeMaster, err error) {

	v, err = GetMasterByNo(mobileno)
	if err != nil {
		return
	}

	err = CheckMasterToken(mobileno, token, reqtime)
	return
}

func IsMasterAudited(userid int) (err error) {
	var v *EhomeMaster
	v, err = GetEhomeMasterById(userid)
	if err != nil {
		return
	}
	if v.Audited != 1 {
		err = fmt.Errorf("Master not audit!")
		return
	}
	return
}

func GetCatidByTitle(fixtype string) (catid int, err error) {

	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id")
	query["Title"] = fixtype

	var ml []interface{}

	ml, err = GetAllEhomeCategory(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No this fixtype %s", fixtype)
	} else {
		catid = ((ml[0].(map[string]interface{}))["Id"]).(int)
	}

	return

}

func GetCatidlist(catidlist string) (idstr string, err error) {
	if len(catidlist) == 0 {
		return
	}
	sl := strings.Split(catidlist, ",")
	var i int
	var id int

	for i = 0; i < len(sl); i++ {
		if i != 0 {
			idstr = idstr + ","
		}
		sl[i] = strings.TrimSpace(sl[i])
		id, err = GetCatidByTitle(sl[i])
		if err != nil {
			return
		}

		idstr = idstr + strconv.Itoa(int(id))
	}

	return
}

func GetFixtypelistbyIdlist(idlist string) (fixtypestr string, err error) {
	var id, i int
	var cat *EhomeCategory
	sl := strings.Split(idlist, ",")
	for i = 0; len(idlist) > 0 && i < len(sl); i++ {
		id, err = strconv.Atoi(sl[i])
		if err != nil {
			return
		}
		cat, err = GetEhomeCategoryById(id)
		if err != nil {
			beego.Error("GetEhomeCategoryById error! id[%s]", sl[i])
			return
		}

		if i != 0 {
			fixtypestr = fixtypestr + ","
		}
		fixtypestr = fixtypestr + cat.Title
	}
	return
}

func CalTotalPrice(Type string, stylepriceid int, size float64, labouridlist string, materialidlist string) (totalprice float64, err error) {
	if Type == "1" {
		var v *EhomeStylePrice
		v, err = GetEhomeStylePriceById(stylepriceid)
		if err != nil {
			return
		}
		totalprice = v.Price * size
		return
	} else {
		var i int
		var id, num int
		var lab *EhomeLabour
		var mat *EhomeMaterial
		sl := strings.Split(labouridlist, ",")

		for i = 0; len(labouridlist) != 0 && i < len(sl); i++ {
			idnum := strings.Split(sl[i], ":")
			if len(idnum) != 2 {
				beego.Error("GetTotalPrice format error! [%s] [%s]", labouridlist, idnum)
				return
			}
			id, err = strconv.Atoi(idnum[0])
			if err != nil {
				return
			}
			if id == 0 {
				continue
			}

			num, err = strconv.Atoi(idnum[1])
			if err != nil {
				return
			}

			lab, err = GetEhomeLabourById(id)
			if err != nil {
				beego.Error("labour id not exist! ", id)
				return
			}
			totalprice = totalprice + lab.Price*float64(num)
		}

		sl = strings.Split(materialidlist, ",")

		for i = 0; len(materialidlist) != 0 && i < len(sl); i++ {
			idnum := strings.Split(sl[i], ":")
			if len(idnum) != 2 {
				beego.Error("GetTotalPrice format error! [%s] [%s]", materialidlist, idnum)
				return
			}
			id, err = strconv.Atoi(idnum[0])
			if err != nil {
				return
			}

			if id == 0 {
				continue
			}

			num, err = strconv.Atoi(idnum[1])
			if err != nil {
				return
			}

			mat, err = GetEhomeMaterialById(id)
			if err != nil {
				beego.Error("material id not exist! ", id)
				return
			}
			totalprice = totalprice + mat.Price*float64(num)
		}
		totalprice = totalprice * 1.2
	}
	return
}

func GetStatusById(id int, usertype int) (status string, err error) {

	if id == ORDER_CANCEL {
		status = "已取消"
		return
	}
	var v *EhomeOrderStatus

	v, err = GetEhomeOrderStatusById(id)
	if err != nil {
		return
	}

	if usertype == 1 {
		status = v.Introduction
	} else {
		status = v.Masterintro
	}

	return
}

func GetImageById(id int) (url string, err error) {
	if id == 0 {
		url = ""
		return
	}

	var v *EhomeImage
	v, err = GetEhomeImageById(id)
	if err != nil {
		return
	}

	url = v.Filename
	return
}

func AddressIdBelongToUser(userid int, id int) (b bool) {
	var v *EhomeFixAddress
	var err error

	v, err = GetEhomeFixAddressById(id)
	if err != nil {
		beego.Error("GetEhomeFixAddressById error!", id, err)
		b = false
		return
	}
	if userid != v.Userid {
		b = false
		return
	}

	b = true
	return
}

func GetDefaultAddress(userid int) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Id", "Userid", "Contactname", "Phone", "Contactaddr", "Region")
	query["Userid"] = strconv.Itoa(userid)

	sortby = append(sortby, "IsDefaultaddr")
	order = append(order, "desc")

	//var ml []interface{}

	limit = 1

	m, err = GetAllEhomeFixAddress(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	//var h EhomeFixAddress

	if len(m) <= 0 {
		err = fmt.Errorf("No  fix address for userid %d", userid)
	}

	for i := 0; i < len(m); i++ {
		tmp := (m[i].(map[string]interface{}))["Region"].(string)
		province, city, region, e := GetRegionDetailById(tmp)
		if e != nil {
			beego.Error("GetRegionDetailById error! %s", tmp)
			err = e
			return
		}
		(m[i].(map[string]interface{}))["Region"] = province + city + region
		(m[i].(map[string]interface{}))["Regionid"] = tmp
	}

	/*else {



		tmp := ml[0].(map[string]interface{})
		h.Userid = tmp["Userid"].(int)
		h.Contactaddr = tmp["Contactaddr"].(string)
		h.Contactname = tmp["Contactname"].(string)
		h.Phone = tmp["Phone"].(string)
	}
	*/

	return
}

func UpdateDefaultAddress(addrid int64, userid int, isdefault int8) (err error) {
	if isdefault != 1 {
		return
	}

	UnsetDefaultAddress(addrid, userid)
	return
}

func GetImageList(idlist string) (ml []interface{}, err error) {
	if len(idlist) == 0 {
		ml = make([]interface{}, 0)
		return
	}
	var id, i int
	var v *EhomeImage
	sl := strings.Split(idlist, ",")
	ml = make([]interface{}, 0)
	for i = 0; i < len(sl); i++ {
		id, err = strconv.Atoi(sl[i])
		if err != nil {
			return
		}
		if id == 0 {
			continue
		}
		v, err = GetEhomeImageById(id)
		if err != nil {
			return
		}
		ml = append(ml, v.Filename)
	}
	return
}

func GetVoiceList(idlist string) (ml []interface{}, err error) {
	var id, i int
	var v *EhomeVoice
	ml = make([]interface{}, 0)
	if len(idlist) == 0 {
		return
	}
	sl := strings.Split(idlist, ",")
	for i = 0; i < len(sl); i++ {
		id, err = strconv.Atoi(sl[i])
		if err != nil {
			return
		}
		v, err = GetEhomeVoiceById(id)
		if err != nil {
			return
		}
		ml = append(ml, v.Filename)
	}
	return
}

func GetRegionById(id string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	query["Regionid"] = id

	limit = 1

	ml, err = GetAllEhomeArea(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No Region for  %s", id)
	}

	return
}

func GetCityById(id string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	query["Cityid"] = id

	limit = 1

	ml, err = GetAllEhomeCity(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No  Citye for  %s", id)
	}

	return
}

func GetProvinceById(id string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	query["Provinceid"] = id

	limit = 1

	ml, err = GetAllEhomeProvince(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No province for %s", id)
	}

	return
}

func GetRegionDetailById(id string) (province string, city string, region string, err error) {
	m, err := GetRegionById(id)
	if err != nil {
		return
	}

	region = (m[0].(EhomeArea)).Region

	m, err = GetCityById(strconv.Itoa((m[0].(EhomeArea)).Fatherid))
	if err != nil {
		return
	}
	city = (m[0].(EhomeCity)).City

	m, err = GetProvinceById(strconv.Itoa((m[0].(EhomeCity)).Fatherid))
	if err != nil {
		return
	}
	province = (m[0].(EhomeProvince)).Province

	return
}

func CheckStyleId(styleid string) (err error) {
	var id int

	id, err = strconv.Atoi(styleid)
	if err != nil {
		return
	}

	v, err := GetEhomeStyleById(id)
	if err != nil {
		return
	}

	if v == nil {
		err = fmt.Errorf("styleid [%s] error!", styleid)
	}

	return
}

func GetWholeFixType(styleid string, stylepriceid int) (msg string, err error) {
	var id int
	id, err = strconv.Atoi(styleid)
	if err != nil {
		return
	}

	v1, err := GetEhomeStyleById(id)
	if err != nil {
		return
	}

	v2, err := GetEhomeStylePriceById(stylepriceid)
	if err != nil {
		return
	}

	msg = v1.Title + " " + v2.Introduction

	return
}

func GetTosList(tosid string) (ml []interface{}, err error) {
	if len(tosid) == 0 {
		ml = make([]interface{}, 0)
		return
	}
	var v *EhomeTos

	sl := strings.Split(tosid, ",")
	var i int
	var id int

	for i = 0; i < len(sl); i++ {
		sl[i] = strings.TrimSpace(sl[i])

		id, err = strconv.Atoi(sl[i])
		v, err = GetEhomeTosById(id)
		ml = append(ml, v.Title)
	}

	return
}

func UpdateTopicClick(Id string) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("update ehome_topic set clicks=clicks+1 where topic_id=%s", Id)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()
	return
}

func AddIncome(v *EhomeIncomeLog) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("update ehome_setting set totalamount=totalamount + %f, Masteramount=Masteramount + %f, platformamount=platformamount+%f  where id=1", v.Price, v.Masteramount, v.Platformamount)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	return
}
