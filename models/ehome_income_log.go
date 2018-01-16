package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EhomeIncomeLog struct {
	Id             int       `orm:"column(id);auto"`
	Orderid        int       `orm:"column(orderid);null"`
	Orderno        string    `orm:"column(orderno);size(25);null"`
	Masterid       int       `orm:"column(masterid);null"`
	Price          float64   `orm:"column(price);null;digits(14);decimals(4)"`
	Dividerate     float64   `orm:"column(dividerate);null;digits(14);decimals(4)"`
	Platformamount float64   `orm:"column(platformamount);null;digits(14);decimals(4)"`
	Masteramount   float64   `orm:"column(masteramount);null;digits(14);decimals(4)"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp);auto_now"`
}

func (t *EhomeIncomeLog) TableName() string {
	return "ehome_income_log"
}

func init() {
	orm.RegisterModel(new(EhomeIncomeLog))
}

// AddEhomeIncomeLog insert a new EhomeIncomeLog into database and returns
// last inserted Id on success.
func AddEhomeIncomeLog(m *EhomeIncomeLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeIncomeLogById retrieves EhomeIncomeLog by Id. Returns error if
// Id doesn't exist
func GetEhomeIncomeLogById(id int) (v *EhomeIncomeLog, err error) {
	o := orm.NewOrm()
	v = &EhomeIncomeLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeIncomeLog retrieves all EhomeIncomeLog matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeIncomeLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeIncomeLog))
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

	var l []EhomeIncomeLog
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

// UpdateEhomeIncomeLog updates EhomeIncomeLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeIncomeLogById(m *EhomeIncomeLog) (err error) {
	o := orm.NewOrm()
	v := EhomeIncomeLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeIncomeLog deletes EhomeIncomeLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeIncomeLog(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeIncomeLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeIncomeLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
