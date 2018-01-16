package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EhomeMaster struct {
	Id               int       `orm:"column(id);auto"`
	Phone            string    `orm:"column(phone);size(20);null"`
	Password         string    `orm:"column(password);size(20);null"`
	Name             string    `orm:"column(name);size(20);null"`
	Idcard           string    `orm:"column(idcard);size(30);null"`
	Tos              string    `orm:"column(tos);size(255);null"`
	Regionid         string    `orm:"column(regionid);size(100);null"`
	Address          string    `orm:"column(address);size(255);null"`
	Idcardimage      string    `orm:"column(idcardimage);size(50);null"`
	Certificateimage string    `orm:"column(certificateimage);size(50);null"`
	Headimageid      int       `orm:"column(headimageid);null`
	Balance          float64   `orm:"column(balance);null;digits(12);decimals(2)"`
	CreateTime       time.Time `orm:"column(create_time);type(timestamp);auto_now"`
	Auditadminid     int       `orm:"column(auditadminid);null`
	AuditTime        time.Time `orm:"column(create_time);type(timestamp);null`
	Audited          int16     `orm:"column(audited);null`
	Cityid           string    `orm:"column(cityid);size(20);null"`
	Bonuspoint       int       `orm:"column(bonuspoint);null`
}

func (t *EhomeMaster) TableName() string {
	return "ehome_master"
}

func init() {
	orm.RegisterModel(new(EhomeMaster))
}

// AddEhomeMaster insert a new EhomeMaster into database and returns
// last inserted Id on success.
func AddEhomeMaster(m *EhomeMaster) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeMasterById retrieves EhomeMaster by Id. Returns error if
// Id doesn't exist
func GetEhomeMasterById(id int) (v *EhomeMaster, err error) {
	o := orm.NewOrm()
	v = &EhomeMaster{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeMaster retrieves all EhomeMaster matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeMaster(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeMaster))
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

	var l []EhomeMaster
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

// UpdateEhomeMaster updates EhomeMaster by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeMasterById(m *EhomeMaster) (err error) {
	o := orm.NewOrm()
	v := EhomeMaster{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeMaster deletes EhomeMaster by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeMaster(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeMaster{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeMaster{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetEhomeMasterByPhone(mobile string) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	query["Phone"] = mobile

	m, err = GetAllEhomeMaster(query, fields, sortby, order, offset, limit)

	return
}

func UpdateEhomeMasterPasswd(m *EhomeMaster) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("update ehome_master set password='%s' where id=%d", m.Password, m.Id)
	_, err = o.Raw(sql).Exec()
	return
}

func UpdateEhomeMasterHeadimage(mobile string, id int) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("update ehome_master set headimageid=%d where Phone='%s'", id, mobile)
	_, err = o.Raw(sql).Exec()
	return
}

func BakupEhomeMaster(id int) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("insert into ehome_master_bak select * from ehome_master where id=%d", id)
	_, err = o.Raw(sql).Exec()
	return
}

func Updatemasterbonuspoint(id int, addpoint int) (err error) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("update ehome_master set bonuspoint=bonuspoint+%d  where id=%d", addpoint, id)
	_, err = o.Raw(sql).Exec()
	return
}
