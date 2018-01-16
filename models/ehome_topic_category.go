package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EhomeTopicCategory struct {
	Id         int       `orm:"column(id);auto"`
	Type       int8      `orm:"column(type)"`
	Title      string    `orm:"column(title);size(100);null"`
	CatImg     string    `orm:"column(topic_img);size(128);null"`
	Status     int8      `orm:"column(status)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now"`
}

func (t *EhomeTopicCategory) TableName() string {
	return "ehome_topic_category"
}

func init() {
	orm.RegisterModel(new(EhomeTopicCategory))
}

// AddEhomeTopicCategory insert a new EhomeTopicCategory into database and returns
// last inserted Id on success.
func AddEhomeTopicCategory(m *EhomeTopicCategory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeTopicCategoryById retrieves EhomeTopicCategory by Id. Returns error if
// Id doesn't exist
func GetEhomeTopicCategoryById(id int) (v *EhomeTopicCategory, err error) {
	o := orm.NewOrm()
	v = &EhomeTopicCategory{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeTopicCategory retrieves all EhomeTopicCategory matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeTopicCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeTopicCategory))
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

	var l []EhomeTopicCategory
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

// UpdateEhomeTopicCategory updates EhomeTopicCategory by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeTopicCategoryById(m *EhomeTopicCategory) (err error) {
	o := orm.NewOrm()
	v := EhomeTopicCategory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeTopicCategory deletes EhomeTopicCategory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeTopicCategory(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeTopicCategory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeTopicCategory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
