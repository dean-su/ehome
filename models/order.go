package models

import (
	_ "errors"
	"fmt"
	_ "strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	ORDER_CANCEL                = 0
	ORDER_PLACE_SUCCESS         = 1
	ORDER_WAITING_ACCEPT        = 2
	ORDER_MASTER_ACCEPTED       = 2
	ORDER_MASTER_SETOUT         = 3
	ORDER_MASTER_ARRIVED        = 4
	ORDER_MASTER_PRICING        = 5
	ORDER_CLIENT_ACCEPT_PRICING = 6
	ORDER_FINISH_CONSTRUCTION   = 8
	ORDER_CHECK_AND_ACCEPT      = 9
	ORDER_CLIENT_PAY            = 10
	ORDER_MASTER_CONFIRM_PAY    = 11
)

//OrderList
func OrderList(reqtype string, usertype int, userid int, limit, offset int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var sql string

	fieldlist := "orderid, orderno, imageid, create_time, catidlist, status, region, type, stylepriceid"

	if usertype == 1 {
		sql = fmt.Sprintf("select %s from ehome_order where userid=%d", fieldlist, userid)
	} else {
		if reqtype == "0" {
			sql = fmt.Sprintf(`select %s from ehome_order where 
			status = 1 and orderid not in (select orderid from ehome_request_order where masterid=%d)`, fieldlist, userid)
		} else if reqtype == "1" {
			sql = fmt.Sprintf(`select %s from ehome_order where 
			masterid=%d`, fieldlist, userid)
		} else {
			err = fmt.Errorf("Reqtype error! [%s]", reqtype)
			return
		}
	}

	sql = fmt.Sprintf("%s order by create_time  desc", sql)

	beego.Info(sql)

	if limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, limit)
	}
	if offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, offset)
	}

	var Orderid []int
	var Orderno []string
	var Imgs []int
	var Create_times []time.Time
	var Catidlist []string
	var Status []int
	var Region []string
	var Type []int8
	var Stylepriceid []int
	var num int64

	num, err = o.Raw(sql).QueryRows(&Orderid, &Orderno, &Imgs, &Create_times, &Catidlist, &Status, &Region, &Type, &Stylepriceid)
	if err != nil {
		return nil, err
	}
	for i := int64(0); i < num; i++ {
		m := make(map[string]interface{})
		m["Orderid"] = Orderid[i]
		m["Orderno"] = Orderno[i]
		m["Img"], err = GetImageById(Imgs[i])
		if err != nil {
			beego.Error("GetImageById error", err)
			return
		}

		now := time.Now().AddDate(0, 0, -1)
		if Create_times[i].After(now) {
			if time.Now().Hour() == Create_times[i].Hour() {
				m["Ordertime"] = fmt.Sprintf("%d分钟前", time.Now().Minute()-Create_times[i].Minute())
			} else {
				m["Ordertime"] = fmt.Sprintf("%d小时前", time.Now().Hour()-Create_times[i].Hour())
			}
		} else {
			m["Ordertime"] = Create_times[i].Format("2006年01月02日")
		}
		if Type[i] == 0 {
			m["Fixtype"], err = GetFixtypelistbyIdlist(Catidlist[i])
			if err != nil {
				beego.Error("GetFixtypelistbyIdlist error! %v", err)
				return
			}
		} else {
			m["Fixtype"], err = GetWholeFixType(Catidlist[i], Stylepriceid[i])
			if err != nil {
				beego.Error("GetWholeFixType error! %v", err)
				return
			}
		}

		m["Status"], err = GetStatusById(Status[i], usertype)
		if err != nil {
			beego.Error("GetStatusById error", err)
			return
		}
		_, m["City"], m["Region"], err = GetRegionDetailById(Region[i])
		m["Regionid"] = Region[i]
		ml = append(ml, m)
	}
	if ml == nil {
		ml = make([]interface{}, 0)
	}

	return
}

func OrderNum() (num int64, err error) {
	o := orm.NewOrm()
	sql := "select count(1) as num from ehome_order o"
	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

type OrderEvalue struct {
	Orderid    string
	Orderno    string
	Appearance string
	Punctual   string
	Service    string
	Quality    string
	Feeback    string
	Shareimage string
	Userid     int
}

func EvaluteOrder(v *OrderEvalue) (err error) {
	o := orm.NewOrm()
	//var num int
	sql := fmt.Sprintf(`update ehome_order set Appearance=%s, Punctual=%s, Service=%s, Quality=%s, Feeback='%s', 
	Shareimage='%s' where orderid=%s and orderno='%s'`,
		v.Appearance, v.Punctual, v.Service, v.Quality, v.Feeback, v.Shareimage, v.Orderid, v.Orderno)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()
	/*
		if num == 0 {
			err = fmt.Errorf("UpdateOrderStatus error! orderid [%d] orderno[%d] not exists!", v.Orderid, v.Orderno)
		}
	*/
	return
}

func UpdateClientConfirmPriceStatus(orderid int, orderno string, status int, reason string, image string) (err error) {
	o := orm.NewOrm()
	//var num int

	sql := fmt.Sprintf(`update ehome_order set Status=%d, rejectpricereason='%s', rejectpriceimage='%s'  where orderid=%d and orderno='%s'`, status, reason, image, orderid, orderno)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	sql = fmt.Sprintf(`update ehome_order_path set finish_time=now()  where orderno='%s' and status=%d`, orderno, status)
	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	/*
		if num == 0 {
			err = fmt.Errorf("UpdateOrderStatus error! orderid [%d] orderno[%d] not exists!", orderid, orderno)
		}
	*/
	return
}

func UpdateOrderStatus(orderid int, orderno string, status int, reason string, image string) (err error) {
	o := orm.NewOrm()
	//var num int

	sql := fmt.Sprintf(`update ehome_order set Status=%d  where orderid=%d and orderno='%s'`, status, orderid, orderno)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	sql = fmt.Sprintf(`update ehome_order_path set finish_time=now()  where orderno='%s' and status=%d`, orderno, status)
	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	/*
		if num == 0 {
			err = fmt.Errorf("UpdateOrderStatus error! orderid [%d] orderno[%d] not exists!", orderid, orderno)
		}
	*/
	return
}

func UpdateOrderStatusLess(orderid int, orderno string, status int) (err error) {
	o := orm.NewOrm()
	//var num int

	sql := fmt.Sprintf(`update ehome_order set Status=%d  where orderid=%d and orderno='%s' and status < %d`, status, orderid, orderno, status)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	sql = fmt.Sprintf(`update ehome_order_path set finish_time=now()  where orderno='%s' and status=%d`, orderno, status)
	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	/*
		if num == 0 {
			err = fmt.Errorf("UpdateOrderStatus error! orderid [%d] orderno[%d] not exists!", orderid, orderno)
		}
	*/
	return
}

func UpdateOrderMaster(orderid string, masterid int) (err error) {
	o := orm.NewOrm()
	//var num int
	sql := fmt.Sprintf(`update ehome_order set masterid=%d, modify_time=now() where orderid=%s`, masterid, orderid)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()

	/*
		if num == 0 {
			err = fmt.Errorf("UpdateOrderStatus error! orderid [%d] orderno[%d] not exists!", orderid, orderno)
		}
	*/
	return
}

func InsertOrderPath(orderno string) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("insert into ehome_order_path (orderno, status, userintro, masterintro) select '%s',id, userintro, masterintro from ehome_order_status order by id;", orderno)
	_, err = o.Raw(sql).Exec()
	return
}

func UpdateOrderMasterPrice(orderid int, orderno string, price float64, reason string, image string) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf(`update ehome_order set MasterPrice=%f, modify_time=now(), pricereason='%s', priceimage='%s'  where orderid=%d and orderno='%s'`, price, reason, image, orderid, orderno)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()
	return
}

func CheckOrder(orderid int, orderno string, userid int, usertype int) (err error) {

	var v *EhomeOrder
	v, err = GetEhomeOrderById(orderid)
	if err != nil {
		return
	}

	if orderno != v.Orderno {
		err = fmt.Errorf("orderid or orderno invalid orderid[%d] orderno[%s] ", orderid, orderno)
		return
	}

	if usertype == 1 {
		if userid != 0 && v.Userid != userid {
			err = fmt.Errorf("user[%d] do not have order orderid[%d] orderno[%s] ", userid, orderid, orderno)
		}
	} else {
		if userid != 0 && v.Masterid != userid {
			err = fmt.Errorf("master[%d] do not have order orderid[%d] orderno[%s] ", userid, orderid, orderno)
		}
	}

	return
}

func GetOrderStatus(orderno string, status int, modifyTime time.Time, usertype int) (ml []interface{}, err error) {

	if status == ORDER_CANCEL {
		ml = make([]interface{}, 1)
		ml[0] = make(map[string]interface{})
		(ml[0].(map[string]interface{}))["Id"] = 0
		(ml[0].(map[string]interface{}))["Index"] = 1
		(ml[0].(map[string]interface{}))["Introduction"] = "已取消"
		(ml[0].(map[string]interface{}))["FinishTime"] = modifyTime.Format("2006-01-02 15:04:05")
		return
	}

	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Id", "Index", "Introduction", "FinishTime", "Masterintro")
	query["Orderno"] = orderno
	var i int

	sortby = append(sortby, "Index")
	order = append(order, "asc")

	ml, err = GetAllEhomeOrderPath(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No this  order %s", orderno)
		return
	}

	for i = 0; i < len(ml); i++ {
		if (ml[i].(map[string]interface{})["Index"]).(int) > status {
			ml[i].(map[string]interface{})["FinishTime"] = ""
		} else {
			ml[i].(map[string]interface{})["FinishTime"] = (ml[i].(map[string]interface{})["FinishTime"].(time.Time)).Format("2006-01-02 15:04:05")

		}
		if status+1 == (ml[i].(map[string]interface{})["Index"]).(int) {
			ml[i].(map[string]interface{})["Next"] = 1
		} else {
			ml[i].(map[string]interface{})["Next"] = 0
		}

		if usertype == 2 {
			ml[i].(map[string]interface{})["Introduction"] = ml[i].(map[string]interface{})["Masterintro"]
		}
		delete(ml[i].(map[string]interface{}), "Masterintro")
	}
	if usertype == 2 {
		ml = ml[1:len(ml)]
	}

	return
}

func GetOrderNoByid(orderid int) (orderno string, err error) {
	var v *EhomeOrder

	v, err = GetEhomeOrderById(orderid)
	if err != nil {
		return
	}

	orderno = v.Orderno
	return
}

func GetOrderPrice(orderno string) (masterid, orderid int, price float64, err error) {
	o := orm.NewOrm()
	var sql string
	var t EhomeOrder
	fieldlist := "Masterprice, orderid, masterid"

	sql = fmt.Sprintf("select %s from %s where 1=1", fieldlist, t.TableName())

	sql = fmt.Sprintf("%s and orderno = '%s'", sql, orderno)

	beego.Info(sql)

	var Masterprice []float64
	var Orderid []int
	var Masterid []int
	var num int64
	num, err = o.Raw(sql).QueryRows(&Masterprice, &Orderid, &Masterid)
	if err != nil {
		return
	}
	for i := int64(0); i < num; i++ {
		price = Masterprice[i]
		masterid = Masterid[i]
		orderid = Orderid[i]

		return
	}

	err = fmt.Errorf("can't find orderno %s", orderno)
	return
}
