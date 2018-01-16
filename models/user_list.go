package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type EcsUsers_Str struct {
	Id          string
	UserName    string
	MobilePhone string
	CreateTime  string
	Selfreg     string
}

func EcsUsers_list(limit, offset int64, cond *EcsUsers_Str) (m []interface{}, err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	if err != nil {
		beego.Error("Using error", err)
		return
	}
	var sql string
	var t EcsUsers
	fieldlist := "user_id,user_name,mobile_phone,create_time,Selfreg"

	sql = fmt.Sprintf("select %s from %s where 1=1", fieldlist, t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and user_id = %s", sql, cond.Id)
	}
	if cond.UserName != "" {
		sql = fmt.Sprintf("%s and  user_name like '%s%%'", sql, cond.UserName)
	}
	if cond.MobilePhone != "" {
		sql = fmt.Sprintf("%s and  mobile_phone like '%s%%'", sql, cond.MobilePhone)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}
	if cond.Selfreg != "" {
		sql = fmt.Sprintf("%s and Selfreg = %s", sql, cond.Selfreg)
	}
	if limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, limit)
	}
	if offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, offset)
	}

	beego.Info(sql)

	var Id []int
	var UserName []string
	var MobilePhone []string
	var CreateTime []time.Time
	var Selfreg []int8
	var num int64
	num, err = o.Raw(sql).QueryRows(&Id, &UserName, &MobilePhone, &CreateTime, &Selfreg)
	if err != nil {
		return nil, err
	}
	for i := int64(0); i < num; i++ {
		tmpm := make(map[string]interface{})
		tmpm["Id"] = Id[i]
		tmpm["UserName"] = UserName[i]
		tmpm["MobilePhone"] = MobilePhone[i]
		tmpm["CreateTime"] = CreateTime[i]
		tmpm["Selfreg"] = Selfreg[i]
		m = append(m, tmpm)
	}
	if m == nil {
		m = make([]interface{}, 0)
	}
	return
}

func EcsUsers_num(cond *EcsUsers_Str) (num int64, err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	if err != nil {
		beego.Error("Using error", err)
		return
	}
	var sql string
	var t EcsUsers
	sql = fmt.Sprintf("select count(1) as num from %s where 1=1", t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and user_id = %s", sql, cond.Id)
	}
	if cond.UserName != "" {
		sql = fmt.Sprintf("%s and user_name like '%s%%'", sql, cond.UserName)
	}
	if cond.MobilePhone != "" {
		sql = fmt.Sprintf("%s and mobile_phone like '%s%%'", sql, cond.MobilePhone)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}
	if cond.Selfreg != "" {
		sql = fmt.Sprintf("%s and Selfreg = %s", sql, cond.Selfreg)
	}

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}
