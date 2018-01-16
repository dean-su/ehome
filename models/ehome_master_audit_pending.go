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

type EhomeMasterAuditPending struct {
	Id         int       `orm:"column(id);auto"`
	Masterid   int       `orm:"column(masterid)"`
	Adminid    int       `orm:"column(adminid)"`
	Commented  string    `orm:"column(commented);size(255);null"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	DealFlag   int8      `orm:"column(deal_flag)"`
	DealTime   time.Time `orm:"column(deal_time);type(timestamp)"`
}

func (t *EhomeMasterAuditPending) TableName() string {
	return "ehome_master_audit_pending"
}

func init() {
	orm.RegisterModel(new(EhomeMasterAuditPending))
}

// AddEhomeMasterAuditPending insert a new EhomeMasterAuditPending into database and returns
// last inserted Id on success.
func AddEhomeMasterAuditPending(m *EhomeMasterAuditPending) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeMasterAuditPendingById retrieves EhomeMasterAuditPending by Id. Returns error if
// Id doesn't exist
func GetEhomeMasterAuditPendingById(id int) (v *EhomeMasterAuditPending, err error) {
	o := orm.NewOrm()
	v = &EhomeMasterAuditPending{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeMasterAuditPending retrieves all EhomeMasterAuditPending matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeMasterAuditPending(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeMasterAuditPending))
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

	var l []EhomeMasterAuditPending
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

// UpdateEhomeMasterAuditPending updates EhomeMasterAuditPending by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeMasterAuditPendingById(m *EhomeMasterAuditPending) (err error) {
	o := orm.NewOrm()
	v := EhomeMasterAuditPending{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeMasterAuditPending deletes EhomeMasterAuditPending by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeMasterAuditPending(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeMasterAuditPending{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeMasterAuditPending{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func DeleteEhomeMasterAuditPendingbymasterid(id int) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("delete from  ehome_master_audit_pending where masterid=%d", id)

	beego.Info(sql)
	_, err = o.Raw(sql).Exec()
	return

}
