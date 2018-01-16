package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func GetMasterNum(cityid, regionid, name, phone string) (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeMaster
	var sql string

	sql = fmt.Sprintf("select count(1) as num from  %s where 1=1", v.TableName())

	if len(regionid) > 0 {
		sql = fmt.Sprintf("%s and regionid=%s", sql, regionid)
	}

	if len(cityid) > 0 {
		sql = fmt.Sprintf("%s and cityid=%s", sql, cityid)
	}

	if len(name) > 0 {
		sql = fmt.Sprintf("%s and name='%s'", sql, name)
	}

	if len(phone) > 0 {
		sql = fmt.Sprintf("%s and phone like '%s%%'", sql, phone)
	}

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

func OrderListNum(cityid, regionid string, statusid string, ordertype string, fixtypeid, mastername, month string) (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeOrder
	var sql string

	sql = fmt.Sprintf("select count(1) as num from  %s where 1=1 ", v.TableName())
	if len(cityid) > 0 {
		sql = fmt.Sprintf("%s and cityid='%s'", sql, cityid)
	}

	if len(regionid) > 0 {
		sql = fmt.Sprintf("%s and region='%s'", sql, regionid)
	}

	if len(statusid) > 0 {
		sql = fmt.Sprintf("%s and status=%s", sql, statusid)
	}

	if len(ordertype) > 0 {
		sql = fmt.Sprintf("%s and type=%s", sql, ordertype)
	}

	if ordertype == "1" {
	} else {
		if len(fixtypeid) > 0 {
			sql = fmt.Sprintf("%s and (catidlist='%s' || catidlist like '%s,%%' || catidlist like '%%,%s,%%' || catidlist like '%%,%s')", sql, fixtypeid, fixtypeid, fixtypeid, fixtypeid)
		}
	}

	if len(mastername) > 0 {
		sql = fmt.Sprintf("%s and masterid in (select id from ehome_master where name='%s')", sql, mastername)
	}

	if len(month) > 0 {
		sql = fmt.Sprintf("%s and date_format(create_time, '%%Y-%%m')=%s", sql, month)
	}

	beego.Info(sql)

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

func GetAuditPendingNum() (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeMasterAuditPending
	var sql string

	sql = fmt.Sprintf("select count(1) as num from  %s where deal_flag=0", v.TableName())

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

func GetCashPendingNum() (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeMasterCashLog
	var sql string

	sql = fmt.Sprintf("select count(1) as num from  %s where deal_flag=0", v.TableName())

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

func GetTopicNum(isadmin int, Id string) (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeTopic
	var sql string

	if isadmin == 0 {
		sql = fmt.Sprintf("select count(1) as num from  %s where status=1", v.TableName())
	} else {
		sql = fmt.Sprintf("select count(1) as num from  %s ", v.TableName())
	}

	if Id != "" && Id != "0" {
		if isadmin != 0 {
			sql = fmt.Sprintf("%s %s", sql, "where")
		} else {
			sql = fmt.Sprintf("%s %s", sql, "and")
		}
		sql = fmt.Sprintf("%s %s='%s'", sql, " topic_catid", Id)
	}
	beego.Error(sql)

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

func GetTopicCatNum(isadmin int, Type string) (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeTopicCategory
	var sql string

	if isadmin == 0 {
		sql = fmt.Sprintf("select count(1) as num from  %s where status=1 and type =%s", v.TableName(), Type)
	} else {
		sql = fmt.Sprintf("select count(1) as num from  %s where type=%s", v.TableName(), Type)
	}

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}
