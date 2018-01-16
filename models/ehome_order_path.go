package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EhomeOrderPath struct {
	Id           int       `orm:"column(id);auto"`
	Orderno      string    `orm:"column(orderno);size(25)"`
	Index        int       `orm:"column(status)"`
	Introduction string    `orm:"column(userintro);size(20)"`
	Masterintro  string    `orm:"column(masterintro);size(20)"`
	FinishTime   time.Time `orm:"column(finish_time);type(timestamp);auto_now"`
}

func (t *EhomeOrderPath) TableName() string {
	return "ehome_order_path"
}

func init() {
	orm.RegisterModel(new(EhomeOrderPath))
}

// AddEhomeOrderPath insert a new EhomeOrderPath into database and returns
// last inserted Id on success.
func AddEhomeOrderPath(m *EhomeOrderPath) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeOrderPathById retrieves EhomeOrderPath by Id. Returns error if
// Id doesn't exist
func GetEhomeOrderPathById(id int) (v *EhomeOrderPath, err error) {
	o := orm.NewOrm()
	v = &EhomeOrderPath{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeOrderPath retrieves all EhomeOrderPath matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeOrderPath(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeOrderPath))
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

	var l []EhomeOrderPath
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

// UpdateEhomeOrderPath updates EhomeOrderPath by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeOrderPathById(m *EhomeOrderPath) (err error) {
	o := orm.NewOrm()
	v := EhomeOrderPath{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeOrderPath deletes EhomeOrderPath by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeOrderPath(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeOrderPath{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeOrderPath{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func DeleteEhomeOrderPathByno(orderno string) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("delete from ehome_order_path  where orderno='%s'", orderno)

	_, err = o.Raw(sql).Exec()

	return
}
