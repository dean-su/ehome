package admins

import (
	"fmt"
)

const (
	_start = iota
	MOBILE_LEN_ERR
	TOKEN_LEN_ERR
	REQTIME_LEN_ERR
	PARAM_ERR
	INVALID_USER
	INVALID_FILENAME
	SAVE_FILE_ERROR
	GET_EHOME_USER_ERROR
	ADD_EHOME_USER_ERROR
	UPDATE_EHOME_USER_ERROR
	USER_TYPE_LEN_ERR
	USER_EXIST
	DB_ERROR
	SEND_SMS_ERR
	REDIS_DO_ERR
	ID_CODE_ERR
	INVALID_ORDER
	INVALID_PASS
	SYS_ERROR
	MASTER_HAVE_ORDER
)

type Err struct {
	Errno  int
	Errmsg string
}

func SetError(m map[string]interface{}, errno int, format string, a ...interface{}) {
	m["status"] = errno
	m["errmsg"] = fmt.Sprintf(format, a...)
	return
}
