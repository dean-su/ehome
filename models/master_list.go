package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type EhomeMaster_Str struct {
	Id          string
	Idcard      string
	Regionid    string
	Headimageid string
	Name        string
	CreateTime  string
	Audited     string
	Phone       string
	Cityid      string
}

func EhomeMaster_list(limit, offset int64, cond *EhomeMaster_Str) (m []interface{}, err error) {
	o := orm.NewOrm()
	var sql string
	var t EhomeMaster
	fieldlist := "Id,Idcard,Regionid,Headimageid,Name,create_time,Audited,Phone,Cityid"

	sql = fmt.Sprintf("select %s from %s where 1=1", fieldlist, t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and Id = %s", sql, cond.Id)
	}
	if cond.Idcard != "" {
		sql = fmt.Sprintf("%s and Idcard = '%s'", sql, cond.Idcard)
	}
	if cond.Regionid != "" {
		sql = fmt.Sprintf("%s and Regionid = '%s'", sql, cond.Regionid)
	}
	if cond.Headimageid != "" {
		sql = fmt.Sprintf("%s and Headimageid = %s", sql, cond.Headimageid)
	}
	if cond.Name != "" {
		sql = fmt.Sprintf("%s and Name = '%s'", sql, cond.Name)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}
	if cond.Audited != "" {
		sql = fmt.Sprintf("%s and Audited = %s", sql, cond.Audited)
	}
	if cond.Phone != "" {
		sql = fmt.Sprintf("%s and Phone like '%s%%'", sql, cond.Phone)
	}
	if cond.Cityid != "" {
		sql = fmt.Sprintf("%s and Cityid = '%s'", sql, cond.Cityid)
	}
	if limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, limit)
	}
	if offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, offset)
	}

	beego.Info(sql)

	var Id []int
	var Idcard []string
	var Regionid []string
	var Headimageid []int
	var Name []string
	var CreateTime []time.Time
	var Audited []int16
	var Phone []string
	var Cityid []string
	var num int64
	num, err = o.Raw(sql).QueryRows(&Id, &Idcard, &Regionid, &Headimageid, &Name, &CreateTime, &Audited, &Phone, &Cityid)
	if err != nil {
		return nil, err
	}
	for i := int64(0); i < num; i++ {
		tmpm := make(map[string]interface{})
		tmpm["Id"] = Id[i]
		tmpm["Idcard"] = Idcard[i]
		tmpm["Regionid"] = Regionid[i]
		tmpm["Headimageid"] = Headimageid[i]
		tmpm["Name"] = Name[i]
		tmpm["CreateTime"] = CreateTime[i]
		tmpm["Audited"] = Audited[i]
		tmpm["Phone"] = Phone[i]
		tmpm["Cityid"] = Cityid[i]
		m = append(m, tmpm)
	}
	if m == nil {
		m = make([]interface{}, 0)
	}
	return
}
