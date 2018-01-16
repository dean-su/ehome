package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EhomeUser struct {
	Id      int    `orm:"column(id);auto"`
	Userid  int    `orm:"column(userid)"`
	Image   string `orm:"column(image);size(255);null"`
	Balance int    `orm:"column(balance);null"`
	Points  int    `orm:"column(points);null"`
}

func (t *EhomeUser) TableName() string {
	return "ehome_user"
}

func init() {
	orm.RegisterModel(new(EhomeUser))
}

// AddEhomeUser insert a new EhomeUser into database and returns
// last inserted Id on success.
func AddEhomeUser(m *EhomeUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeUserById retrieves EhomeUser by Id. Returns error if
// Id doesn't exist
func GetEhomeUserById(id int) (v *EhomeUser, err error) {
	o := orm.NewOrm()
	v = &EhomeUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeUser retrieves all EhomeUser matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeUser))
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

	var l []EhomeUser
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

// UpdateEhomeUser updates EhomeUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeUserById(m *EhomeUser) (err error) {
	o := orm.NewOrm()
	v := EhomeUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeUser deletes EhomeUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeUser(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetEhomeUserByUid(uid int) (v []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	fields = append(fields, "Id", "Image", "Balance", "Points")
	query["Userid"] = strconv.Itoa(uid)

	v, err = GetAllEhomeUser(query, fields, sortby, order, offset, limit)

	/*
		if err != nil {
			v = make([]interface{}, 0)
		}

		if v == nil || len(v) == 0 {
			v = make([]interface{}, 0)
			err = fmt.Errorf("userid [%d] not exists", uid)
		}
	*/

	return
}
