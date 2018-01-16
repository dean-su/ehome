package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type EhomeRequestOrder_Str struct {
	Id         string
	Orderid    string
	Masterid   string
	CreateTime string
}

func EhomeRequestOrder_list(limit, offset int64, cond *EhomeRequestOrder_Str) (m []interface{}, err error) {
	o := orm.NewOrm()
	var sql string
	var t EhomeRequestOrder
	fieldlist := "Id,Orderid,Masterid,create_time"

	sql = fmt.Sprintf("select distinct %s from %s where 1=1", fieldlist, t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and Id = %s", sql, cond.Id)
	}
	if cond.Orderid != "" {
		sql = fmt.Sprintf("%s and Orderid = %s", sql, cond.Orderid)
	}
	if cond.Masterid != "" {
		sql = fmt.Sprintf("%s and Masterid = %s", sql, cond.Masterid)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}
	if limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, limit)
	}
	if offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, offset)
	}

	beego.Info(sql)

	var Id []int
	var Orderid []int
	var Masterid []int
	var CreateTime []time.Time
	var num int64
	num, err = o.Raw(sql).QueryRows(&Id, &Orderid, &Masterid, &CreateTime)
	if err != nil {
		return nil, err
	}
	for i := int64(0); i < num; i++ {
		tmpm := make(map[string]interface{})
		tmpm["Id"] = Id[i]
		tmpm["Orderid"] = Orderid[i]
		tmpm["Masterid"] = Masterid[i]
		tmpm["CreateTime"] = CreateTime[i]
		m = append(m, tmpm)
	}
	if m == nil {
		m = make([]interface{}, 0)
	}
	return
}

func EhomeRequestOrder_num(cond *EhomeRequestOrder_Str) (num int64, err error) {
	o := orm.NewOrm()
	var sql string
	var t EhomeRequestOrder
	sql = fmt.Sprintf("select count(1) as num from %s where 1=1", t.TableName())
	if cond.Id != "" {
		sql = fmt.Sprintf("%s and Id = %s", sql, cond.Id)
	}
	if cond.Orderid != "" {
		sql = fmt.Sprintf("%s and Orderid = %s", sql, cond.Orderid)
	}
	if cond.Masterid != "" {
		sql = fmt.Sprintf("%s and Masterid = %s", sql, cond.Masterid)
	}
	if cond.CreateTime != "" {
		sql = fmt.Sprintf("%s and create_time = %s", sql, cond.CreateTime)
	}

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

func RequestOrder(v *EhomeRequestOrder) (err error) {
	o := orm.NewOrm()
	var sql string

	sql = fmt.Sprintf("insert into ehome_request_order(orderid, masterid) values(%d,%d)", v.Orderid, v.Masterid)

	_, err = o.Raw(sql).Exec()
	return
}
