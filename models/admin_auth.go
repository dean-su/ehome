package models

import (
	"fmt"
	_ "github.com/astaxie/beego"
)

func CheckAdminUser(name string, token string, reqtime int64) (err error) {
	_, err = ValidAdmin(name)
	if err != nil {
		return
	}

	err = CheckAdminToken(name, token, reqtime)

	return
}

func GetAdminId(name string) (id int, err error) {
	var m []interface{}
	m, err = ValidAdmin(name)
	if err != nil {
		return
	}
	if len(m) != 1 {
		err = fmt.Errorf("Can't find admin by name [%s]", name)
		return
	}
	id = (m[0].(EhomeAdmin)).Id
	return
}

func ValidAdmin(name string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	query["Name"] = name

	limit = 1

	ml, err = GetAllEhomeAdmin(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		err = fmt.Errorf("No  admin for %s", name)
	}

	return
}
