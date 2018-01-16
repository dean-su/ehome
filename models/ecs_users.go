package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EcsUsers struct {
	Id          int       `orm:"column(user_id);auto"`
	UserName    string    `orm:"column(user_name);size(60)"`
	Password    string    `orm:"column(password);size(32)"`
	Alias       string    `orm:"column(alias);size(60)"`
	Msn         string    `orm:"column(msn);size(60)"`
	Qq          string    `orm:"column(qq);size(20)"`
	OfficePhone string    `orm:"column(office_phone);size(20)"`
	HomePhone   string    `orm:"column(home_phone);size(20)"`
	MobilePhone string    `orm:"column(mobile_phone);size(20)"`
	CreditLine  float64   `orm:"column(credit_line);digits(10);decimals(2)"`
	CreateTime  time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	Selfreg     int8      `orm:"column(selfreg);null"`
}

func (t *EcsUsers) TableName() string {
	return "ecs_users"
}

func init() {
	orm.RegisterModel(new(EcsUsers))
}

// AddEcsUsers insert a new EcsUsers into database and returns
// last inserted Id on success.
func AddEcsUsers(m *EcsUsers) (id int64, err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	if err != nil {
		beego.Info("Using error", err)
	}
	id, err = o.Insert(m)
	return
}

// GetEcsUsersById retrieves EcsUsers by Id. Returns error if
// Id doesn't exist
func GetEcsUsersByPhone(id string) (num int64, err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	var v []string
	sql := fmt.Sprintf("select 1 from ecs.ecs_users where mobile_phone = '%s'", id)

	num, err = o.Raw(sql).QueryRows(&v)
	beego.Info("m", err)

	return num, err
}

func GetEcsUsersById(id int) (v *EcsUsers, err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	v = &EcsUsers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEcsUsers retrieves all EcsUsers matches certain condition. Returns empty list if
// no records exist
func GetAllEcsUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	qs := o.QueryTable(new(EcsUsers))
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

	var l []EcsUsers
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

// UpdateEcsUsers updates EcsUsers by Id and returns error if
// the record to be updated doesn't exist
func UpdateEcsUsersById(m *EcsUsers) (err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	if err != nil {
		beego.Error("Using error", err)
		return
	}
	v := EcsUsers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEcsUsers deletes EcsUsers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEcsUsers(id int) (err error) {
	o := orm.NewOrm()
	err = o.Using("ecs")
	v := EcsUsers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EcsUsers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func UpdateEcsUsersPasswd(m *EcsUsers) (err error) {
	o := orm.NewOrm()
	o.Using("ecs")
	sql := fmt.Sprintf("update ecs_users set password='%s' where user_id=%d", m.Password, m.Id)
	_, err = o.Raw(sql).Exec()
	return
}
