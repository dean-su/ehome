package admins

import (
	"fmt"
	"reflect"
)

func GetComUser(c interface{}, m map[string]interface{}) (name string, token string, reqtime int64, err error) {
	args := []reflect.Value{reflect.ValueOf("Name")}
	name = reflect.ValueOf(c).MethodByName("GetString").Call(args)[0].Interface().(string)
	args = []reflect.Value{reflect.ValueOf("Token")}
	token = reflect.ValueOf(c).MethodByName("GetString").Call(args)[0].Interface().(string)
	args = []reflect.Value{reflect.ValueOf("Reqtime")}
	result := reflect.ValueOf(c).MethodByName("GetInt64").Call(args)
	reqtime = result[0].Interface().(int64)

	if len(name) == 0 {
		SetError(m, MOBILE_LEN_ERR, "param name not exist !")
		err = fmt.Errorf("Name not exist or len error!")
		return
	}

	if len(token) == 0 {
		SetError(m, TOKEN_LEN_ERR, "Token not exist!")
		err = fmt.Errorf("Token not exist!")
		return
	}

	if reqtime == 0 {
		SetError(m, REQTIME_LEN_ERR, "Reqtime error!")
		err = fmt.Errorf("Reqtime error")
		return
	}

	return
}

func GetPage(c interface{}) (limit, offset int64) {
	args := []reflect.Value{reflect.ValueOf("Pagenum")}
	result := reflect.ValueOf(c).MethodByName("GetInt64").Call(args)
	n := result[0].Interface().(int64)

	args = []reflect.Value{reflect.ValueOf("Recperpage")}
	result = reflect.ValueOf(c).MethodByName("GetInt64").Call(args)
	m := result[0].Interface().(int64)

	if n == 0 || m == 0 {
		limit = 0
		offset = 0
	} else {
		n = n - 1
		limit = m
		offset = m * n
	}

	return
}
