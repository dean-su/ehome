package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EhomeFixAddress struct {
	Id            int    `orm:"column(id);auto"`
	Userid        int    `orm:"column(userid);null"`
	Region        string `orm:"column(region);size(20);null"`
	Cityid        string `orm:"column(cityid);size(20);null"`
	Contactname   string `orm:"column(contactname);size(20);null"`
	Phone         string `orm:"column(phone);size(20);null"`
	Contactaddr   string `orm:"column(contactaddress);size(100);null"`
	IsDefaultaddr int8   `orm:"column(isdefaultaddr);null"`
}

func (t *EhomeFixAddress) TableName() string {
	return "ehome_fix_address"
}

func init() {
	orm.RegisterModel(new(EhomeFixAddress))
}

// AddEhomeFixAddress insert a new EhomeFixAddress into database and returns
// last inserted Id on success.
func AddEhomeFixAddress(m *EhomeFixAddress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeFixAddressById retrieves EhomeFixAddress by Id. Returns error if
// Id doesn't exist
func GetEhomeFixAddressById(id int) (v *EhomeFixAddress, err error) {
	o := orm.NewOrm()
	v = &EhomeFixAddress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeFixAddress retrieves all EhomeFixAddress matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeFixAddress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeFixAddress))
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

	var l []EhomeFixAddress
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

// UpdateEhomeFixAddress updates EhomeFixAddress by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeFixAddressById(m *EhomeFixAddress) (err error) {
	o := orm.NewOrm()
	v := EhomeFixAddress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeFixAddress deletes EhomeFixAddress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeFixAddress(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeFixAddress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeFixAddress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func UnsetDefaultAddress(id int64, userid int) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("update ehome_fix_address set isdefaultaddr=0 where userid = %d and id!=%d", userid, id)
	_, err = o.Raw(sql).Exec()
	return
}
