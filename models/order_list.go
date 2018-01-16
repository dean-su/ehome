package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type EhomeOrder_Str struct {
	Id              string
	Orderno         string
	Status          string
	Regionid        string
	Contactphone    string
	Appointmenttime string
	Contactname     string
	Catidlist       string
	Price           string
	Masterprice     string
	Cityid          string
	CreateTime      string
	Masterid        string
	Type            string
	Stylepriceid    string
	Mastername      string
	Month           string
}

func EhomeOrder_list(limit, offset int64, cond *EhomeOrder_Str) (m []interface{}, err error) {
	o := orm.NewOrm()
	var sql string
	var t EhomeOrder
	fieldlist := "orderid,Orderno,Status,Region,Contactphone,Appointmenttime,Contactname,Catidlist,Price,Masterprice,Cityid,create_time,Masterid,Type,Stylepriceid"

	sql = fmt.Sprintf("select %s from %s where 1=1", fieldlist, t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and orderid = %s", sql, cond.Id)
	}
	if cond.Orderno != "" {
		sql = fmt.Sprintf("%s and Orderno = '%s'", sql, cond.Orderno)
	}
	if cond.Status != "" {
		sql = fmt.Sprintf("%s and Status = %s", sql, cond.Status)
	}
	if cond.Regionid != "" {
		sql = fmt.Sprintf("%s and Region = '%s'", sql, cond.Regionid)
	}
	if cond.Contactphone != "" {
		sql = fmt.Sprintf("%s and Contactphone = '%s'", sql, cond.Contactphone)
	}
	if cond.Appointmenttime != "" {
		sql = fmt.Sprintf("%s and Appointmenttime = %s", sql, cond.Appointmenttime)
	}
	if cond.Contactname != "" {
		sql = fmt.Sprintf("%s and Contactname = '%s'", sql, cond.Contactname)
	}
	if cond.Catidlist != "" {
		sql = fmt.Sprintf("%s and Catidlist = '%s'", sql, cond.Catidlist)
	}
	if cond.Price != "" {
		sql = fmt.Sprintf("%s and Price = %s", sql, cond.Price)
	}
	if cond.Masterprice != "" {
		sql = fmt.Sprintf("%s and Masterprice = %s", sql, cond.Masterprice)
	}
	if cond.Cityid != "" {
		sql = fmt.Sprintf("%s and Cityid = '%s'", sql, cond.Cityid)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}
	if cond.Masterid != "" {
		sql = fmt.Sprintf("%s and Masterid = %s", sql, cond.Masterid)
	}
	if cond.Type != "" {
		sql = fmt.Sprintf("%s and Type = %s", sql, cond.Type)
	}
	if cond.Stylepriceid != "" {
		sql = fmt.Sprintf("%s and Stylepriceid = %s", sql, cond.Stylepriceid)
	}

	if len(cond.Mastername) > 0 {
		sql = fmt.Sprintf("%s and masterid in (select id from ehome_master where name='%s')", sql, cond.Mastername)
	}

	if len(cond.Month) > 0 {
		sql = fmt.Sprintf("%s and date_format(create_time, '%%Y-%%m')=%s", sql, cond.Month)
	}

	sql = fmt.Sprintf("%s order by create_time desc", sql)

	if limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, limit)
	}
	if offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, offset)
	}

	beego.Info(sql)

	var Id []int
	var Orderno []string
	var Status []int
	var Regionid []string
	var Contactphone []string
	var Appointmenttime []time.Time
	var Contactname []string
	var Catidlist []string
	var Price []float64
	var Masterprice []float64
	var Cityid []string
	var CreateTime []time.Time
	var Masterid []int
	var Type []int8
	var Stylepriceid []int
	var num int64
	num, err = o.Raw(sql).QueryRows(&Id, &Orderno, &Status, &Regionid, &Contactphone, &Appointmenttime, &Contactname, &Catidlist, &Price, &Masterprice, &Cityid, &CreateTime, &Masterid, &Type, &Stylepriceid)
	if err != nil {
		return nil, err
	}
	for i := int64(0); i < num; i++ {
		tmpm := make(map[string]interface{})
		tmpm["Id"] = Id[i]
		tmpm["Orderno"] = Orderno[i]
		tmpm["Status"] = Status[i]
		tmpm["Regionid"] = Regionid[i]
		tmpm["Contactphone"] = Contactphone[i]
		tmpm["Appointmenttime"] = Appointmenttime[i]
		tmpm["Contactname"] = Contactname[i]
		tmpm["Catidlist"] = Catidlist[i]
		tmpm["Price"] = Price[i]
		tmpm["Masterprice"] = Masterprice[i]
		tmpm["Cityid"] = Cityid[i]
		tmpm["CreateTime"] = CreateTime[i]
		tmpm["Masterid"] = Masterid[i]
		tmpm["Type"] = Type[i]
		tmpm["Stylepriceid"] = Stylepriceid[i]
		m = append(m, tmpm)
	}
	if m == nil {
		m = make([]interface{}, 0)
	}
	return
}

func EhomeOrder_num(cond *EhomeOrder_Str) (num int64, err error) {
	o := orm.NewOrm()
	var sql string
	var t EhomeOrder
	sql = fmt.Sprintf("select count(1) as num from %s where 1=1", t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and orderid = %s", sql, cond.Id)
	}
	if cond.Orderno != "" {
		sql = fmt.Sprintf("%s and Orderno = '%s'", sql, cond.Orderno)
	}
	if cond.Status != "" {
		sql = fmt.Sprintf("%s and Status = %s", sql, cond.Status)
	}
	if cond.Regionid != "" {
		sql = fmt.Sprintf("%s and Region = '%s'", sql, cond.Regionid)
	}
	if cond.Contactphone != "" {
		sql = fmt.Sprintf("%s and Contactphone = '%s'", sql, cond.Contactphone)
	}
	if cond.Appointmenttime != "" {
		sql = fmt.Sprintf("%s and Appointmenttime = %s", sql, cond.Appointmenttime)
	}
	if cond.Contactname != "" {
		sql = fmt.Sprintf("%s and Contactname = '%s'", sql, cond.Contactname)
	}
	if cond.Catidlist != "" {
		sql = fmt.Sprintf("%s and Catidlist = '%s'", sql, cond.Catidlist)
	}
	if cond.Price != "" {
		sql = fmt.Sprintf("%s and Price = %s", sql, cond.Price)
	}
	if cond.Masterprice != "" {
		sql = fmt.Sprintf("%s and Masterprice = %s", sql, cond.Masterprice)
	}
	if cond.Cityid != "" {
		sql = fmt.Sprintf("%s and Cityid = '%s'", sql, cond.Cityid)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}
	if cond.Masterid != "" {
		sql = fmt.Sprintf("%s and Masterid = %s", sql, cond.Masterid)
	}
	if cond.Type != "" {
		sql = fmt.Sprintf("%s and Type = %s", sql, cond.Type)
	}
	if cond.Stylepriceid != "" {
		sql = fmt.Sprintf("%s and Stylepriceid = %s", sql, cond.Stylepriceid)
	}

	if len(cond.Mastername) > 0 {
		sql = fmt.Sprintf("%s and masterid in (select id from ehome_master where name='%s')", sql, cond.Mastername)
	}

	if len(cond.Month) > 0 {
		sql = fmt.Sprintf("%s and date_format(create_time, '%%Y-%%m')=%s", sql, cond.Month)
	}

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}
