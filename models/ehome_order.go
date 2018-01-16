package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EhomeOrder struct {
	Id                int       `orm:"column(orderid);auto"`
	Catidlist         string    `orm:"column(catidlist);size(50);null"`
	Imageid           int       `orm:"column(imageid);null"`
	Status            int       `orm:"column(status);null"`
	Userid            int       `orm:"column(userid)"`
	Price             float64   `orm:"column(price);null;digits(12);decimals(2)"`
	Orderno           string    `orm:"column(orderno);size(25);null"`
	Voiceidlist       string    `orm:"column(voiceidlist);size(255);null"`
	Imageidlist       string    `orm:"column(imageidlist);size(255);null"`
	Labouridlist      string    `orm:"column(labouridlist);size(255);null"`
	Materialidlist    string    `orm:"column(materialidlist);size(255);null"`
	CreateTime        time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	Region            string    `orm:"column(region);size(20);null"`
	Cityid            string    `orm:"column(cityid);size(20);null"`
	Contactaddr       string    `orm:"column(contactaddr);size(255);null"`
	Contactname       string    `orm:"column(contactname);size(255);null"`
	Contactphone      string    `orm:"column(contactphone);size(25);null"`
	Appointmenttime   time.Time `orm:"column(appointmenttime);type(timestamp)"`
	Attact            string    `orm:"column(attact);null"`
	Masterid          int       `orm:"column(masterid);null"`
	Appearance        int8      `orm:"column(appearance);null"`
	Punctual          int8      `orm:"column(punctual);null"`
	Service           int8      `orm:"column(service);null"`
	Quality           int8      `orm:"column(quality);null"`
	Feeback           string    `orm:"column(feeback);null"`
	Shareimage        string    `orm:"column(shareimage);size(100);null"`
	Room              int8      `orm:"column(room);null"`
	Hall              int8      `orm:"column(hall);null"`
	Kitchen           int8      `orm:"column(kitchen);null"`
	Toilet            int8      `orm:"column(toilet);null"`
	Size              float64   `orm:"column(size);null;digits(12);decimals(2)"`
	Stylepriceid      int       `orm:"column(stylepriceid);null"`
	Type              int8      `orm:"column(type);null"`
	ModifyTime        time.Time `orm:"column(modify_time);type(timestamp)"`
	Masterprice       float64   `orm:"column(masterprice);null;digits(12);decimals(2)"`
	Pricereason       string    `orm:"column(pricereason);size(255);null"`
	Priceimage        string    `orm:"column(priceimage);size(100);null"`
	Rejectpricereason string    `orm:"column(rejectpricereason);size(255);null"`
	Rejectpriceimage  string    `orm:"column(rejectpriceimage);size(100);null"`
	Remark            string    `orm:"column(remark);null"`
}

func (t *EhomeOrder) TableName() string {
	return "ehome_order"
}

func init() {
	orm.RegisterModel(new(EhomeOrder))
}

// AddEhomeOrder insert a new EhomeOrder into database and returns
// last inserted Id on success.
func AddEhomeOrder(m *EhomeOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeOrderById retrieves EhomeOrder by Id. Returns error if
// Id doesn't exist
func GetEhomeOrderById(id int) (v *EhomeOrder, err error) {
	o := orm.NewOrm()
	v = &EhomeOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeOrder retrieves all EhomeOrder matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeOrder))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EhomeOrder
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEhomeOrder updates EhomeOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeOrderById(m *EhomeOrder) (err error) {
	o := orm.NewOrm()
	v := EhomeOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeOrder deletes EhomeOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeOrder(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
