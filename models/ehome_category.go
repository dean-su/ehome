package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EhomeCategory struct {
	Id      int     `orm:"column(id);auto"`
	Title   string  `orm:"column(title);size(100);null"`
	Picture string  `orm:"column(picture);size(255);null"`
	Thumb   string  `orm:"column(thumb);size(255);null"`
	Price   float64 `orm:"column(price);null;digits(12);decimals(2)"`
	Time    string  `orm:"column(time);size(100);null"`
	Attach  string  `orm:"column(attach);size(255);null"`
	Detail  string  `orm:"column(detail);null"`
	Sort    int16   `orm:"column(sort);null"`
	Addtime int     `orm:"column(addtime);null"`
}

func (t *EhomeCategory) TableName() string {
	return "ehome_category"
}

func init() {
	orm.RegisterModel(new(EhomeCategory))
}

// AddEhomeCategory insert a new EhomeCategory into database and returns
// last inserted Id on success.
func AddEhomeCategory(m *EhomeCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeCategoryById retrieves EhomeCategory by Id. Returns error if
// Id doesn't exist
func GetEhomeCategoryById(id int) (v *EhomeCategory, err error) {
	o := orm.NewOrm()
	v = &EhomeCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeCategory retrieves all EhomeCategory matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeCategory))
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

	var l []EhomeCategory
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

// UpdateEhomeCategory updates EhomeCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeCategoryById(m *EhomeCategory) (err error) {
	o := orm.NewOrm()
	v := EhomeCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeCategory deletes EhomeCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeCategory(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
