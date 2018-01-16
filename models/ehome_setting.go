package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type EhomeSetting struct {
	Id                    int     `orm:"column(id);auto"`
	Emailcontent          string  `orm:"column(emailcontent);null"`
	Bonuspoint2moneyrate  float32 `orm:"column(bonuspoint2moneyrate);null"`
	Bonuspointdifficulty  int     `orm:"column(bonuspointdifficulty);null"`
	Redpackagetimesperday int     `orm:"column(redpackagetimesperday);null"`
	Bonuspointrange       string  `orm:"column(bonuspointrange);size(30);null"`
	Totalamount           float64 `orm:"column(totalamount);null;digits(14);decimals(4)"`
	Platformamount        float64 `orm:"column(platformamount);null;digits(14);decimals(4)"`
	Masteramount          float64 `orm:"column(masteramount);null;digits(14);decimals(4)"`
	Clientappid           string  `orm:"column(clientappid);size(20);null"`
	Masterappid           string  `orm:"column(masterappid);size(20);null"`
	Clientprivatekey      string  `orm:"column(clientprivatekey);null"`
	Clientpublickey       string  `orm:"column(clientpublickey);null"`
	Masterprivatekey      string  `orm:"column(masterprivatekey);null"`
	Masterpublickey       string  `orm:"column(masterpublickey);null"`
}

func (t *EhomeSetting) TableName() string {
	return "ehome_setting"
}

func init() {
	orm.RegisterModel(new(EhomeSetting))
}

// AddEhomeSetting insert a new EhomeSetting into database and returns
// last inserted Id on success.
func AddEhomeSetting(m *EhomeSetting) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEhomeSettingById retrieves EhomeSetting by Id. Returns error if
// Id doesn't exist
func GetEhomeSettingById(id int) (v *EhomeSetting, err error) {
	o := orm.NewOrm()
	v = &EhomeSetting{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEhomeSetting retrieves all EhomeSetting matches certain condition. Returns empty list if
// no records exist
func GetAllEhomeSetting(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EhomeSetting))
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

	var l []EhomeSetting
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

// UpdateEhomeSetting updates EhomeSetting by Id and returns error if
// the record to be updated doesn't exist
func UpdateEhomeSettingById(m *EhomeSetting) (err error) {
	o := orm.NewOrm()
	v := EhomeSetting{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEhomeSetting deletes EhomeSetting by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEhomeSetting(id int) (err error) {
	o := orm.NewOrm()
	v := EhomeSetting{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EhomeSetting{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetEhomeSetting() (v *EhomeSetting, err error) {
	return GetEhomeSettingById(1)
}
