package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetAdminRegion() (m [](map[string]interface{}), err error) {
	o := orm.NewOrm()
	sql := "select distinct areaid, area from ehome_area where areaid in (select regionid from ehome_master)"

	var Regionid []int
	var Region []string
	var num int64

	num, err = o.Raw(sql).QueryRows(&Regionid, &Region)

	m = make([](map[string]interface{}), num)
	for i := int64(0); i < num; i++ {
		m[i] = make(map[string]interface{})
		m[i]["Regionid"] = Regionid[i]
		m[i]["Region"] = Region[i]
	}

	return
}

func GetOrderRegion() (m [](map[string]interface{}), err error) {
	o := orm.NewOrm()
	sql := "select distinct areaid, area from ehome_area where areaid in (select region from ehome_order)"

	var Regionid []int
	var Region []string
	var num int64

	num, err = o.Raw(sql).QueryRows(&Regionid, &Region)

	m = make([](map[string]interface{}), num)
	for i := int64(0); i < num; i++ {
		m[i] = make(map[string]interface{})
		m[i]["Regionid"] = Regionid[i]
		m[i]["Region"] = Region[i]
	}

	return
}

func AuditEhomeMaster(id, masterid, adminid int, value int, commented string) (err error) {
	o := orm.NewOrm()

	sql := fmt.Sprintf("update ehome_master set audited=%d, auditadminid=%d, audit_time=now() where id=%d",
		value, adminid, masterid)

	_, err = o.Raw(sql).Exec()
	if err != nil {
		return
	}

	sql = fmt.Sprintf("update ehome_master_audit_pending set adminid=%d, commented='%s', deal_flag=1, deal_time=now() where id=%d",
		adminid, commented, id)
	_, err = o.Raw(sql).Exec()

	return
}

func DealCash(id, masterid, adminid int) (err error) {
	o := orm.NewOrm()

	sql := fmt.Sprintf("update ehome_master_cash_log set deal_flag=1, adminid=%d, deal_time=now() where id=%d",
		adminid, id)

	_, err = o.Raw(sql).Exec()
	return
}

func GetOrderStatusList(gettotal, ordertype string) (m []interface{}, err error) {
	var limit int64
	var offset int64
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	fields = []string{"Id", "Introduction"}

	m, err = GetAllEhomeOrderStatus(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(m) == 0 {
		m = make([]interface{}, 0)
	} else {
		tmp := make(map[string]interface{})
		tmp["Id"] = 0
		tmp["Introduction"] = "已取消"
		m = append(m, tmp)
	}

	if gettotal == "1" {
		var tm map[int]int64
		tm, err = GetOrderNumByStatus(ordertype)
		for i := 0; i < len(m); i++ {
			m[i].(map[string]interface{})["Total"] = tm[m[i].(map[string]interface{})["Id"].(int)]
		}
	}

	return

}

func GetOrderNumByStatus(ordertype string) (m map[int]int64, err error) {
	o := orm.NewOrm()

	var sql string
	sql = fmt.Sprintf("select status, count(1) as num from ehome_order where 1=1")

	sql = fmt.Sprintf("%s and type=%s", sql, ordertype)

	sql = fmt.Sprintf("%s group by  status", sql)

	var status []int
	var total []int64

	var num int64
	num, err = o.Raw(sql).QueryRows(&status, &total)
	if err != nil {
		return nil, err
	}

	m = make(map[int]int64)
	for i := int64(0); i < num; i++ {
		m[status[i]] = total[i]
	}

	return
}

func GetEhomeMasterOrderNum(id int) (num int64, err error) {
	o := orm.NewOrm()

	sql := fmt.Sprintf("select count(1) as num from ehome_order where masterid=%d", id)

	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]

	return
}
