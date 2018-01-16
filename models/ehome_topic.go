package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EhomeTopic struct {
	Id         int       `orm:"column(topic_id);auto"`
	TopicCatid int       `orm:"column(topic_catid)"`
	Source     string    `orm:"column(source);size(50)"`
	Title      string    `orm:"column(title);size(255)"`
	Desc       string    `orm:"column(intro);size(255)"`
	Data       string    `orm:"column(data)"`
	Img        string    `orm:"column(topic_img);size(128)"`
	Clicks     int       `orm:"column(clicks);null"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	Status     int16     `orm:"column(status)"`
}

func (t *EhomeTopic) TableName() string {
	return "ehome_topic"
}

func init() {
	//beego.Info("topic register init start")

	orm.RegisterModel(new(EhomeTopic))
	//beego.Info("topic register init end")
}

// AddEhomeTopic insert a new EhomeTopic into database and returns
// last inserted Id on success.
func AddEhomeTopic(m *EhomeTopic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeTopicById retrieves EhomeTopic by Id. Returns error if
// Id doesn't exist
func GetEhomeTopicById(id int) (v *EhomeTopic, err error) {
	o := orm.NewOrm()
	v = &EhomeTopic{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeTopic retrieves all EhomeTopic matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeTopic(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeTopic))
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

	var l []EhomeTopic
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

// UpdateEhomeTopic updates EhomeTopic by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeTopicById(m *EhomeTopic) (err error) {
	o := orm.NewOrm()
	v := EhomeTopic{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeTopic deletes EhomeTopic by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeTopic(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeTopic{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeTopic{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
