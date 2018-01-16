package models

import (
	"crypto/md5"
	"ehome/db"
	"fmt"
	"github.com/astaxie/beego"

	"io"
	"strconv"
	"time"
)

const (
	user   = "0"
	master = "1"
	admin  = "2"
)

func GetToken() (token string) {
	crutime := time.Now().Unix()

	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token = fmt.Sprintf("%x", h.Sum(nil))

	return
}

func UserToken(mobile string) (token string) {
	token = GetToken()
	redis := db.RedisClient.Get()
	defer redis.Close()

	_, err := redis.Do("SET", user+token+mobile, user)
	if err != nil {
		beego.Error("UserToken redis Do error! %v", err)
	}

	return
}

func MasterToken(mobile string) (token string) {
	token = GetToken()
	redis := db.RedisClient.Get()
	defer redis.Close()

	_, err := redis.Do("SET", master+token+mobile, master)
	if err != nil {
		beego.Error("MasterToken redis Do error! %v", err)
	}

	return
}

func AdminToken(name string) (token string) {
	token = GetToken()
	redis := db.RedisClient.Get()
	defer redis.Close()

	_, err := redis.Do("SET", admin+token+name, admin)
	if err != nil {
		beego.Error("AdminToken redis Do error! %v", err)
	}

	return
}

func CheckUserToken(mobile string, token string, reqtime int64) error {
	redis := db.RedisClient.Get()
	defer redis.Close()

	r, err := redis.Do("GET", user+token+mobile)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("CheckUserToken error ! [%s] [%s]", mobile, token)
	}

	if string(r.([]byte)) != user {
		return fmt.Errorf("CheckUserToken error! impossible!")
	}

	return nil
}

func CheckMasterToken(mobile string, token string, reqtime int64) error {
	redis := db.RedisClient.Get()
	defer redis.Close()

	r, err := redis.Do("GET", master+token+mobile)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("CheckMasterToken error ! [%s] [%s]", mobile, token)
	}

	if string(r.([]byte)) != master {
		return fmt.Errorf("CheckMasterToken error! impossible!")
	}

	return nil
}

func CheckAdminToken(name string, token string, reqtime int64) error {

	redis := db.RedisClient.Get()
	defer redis.Close()
	r, err := redis.Do("GET", admin+token+name)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("CheckAdminToken error ! [%s] [%s]", name, token)
	}

	if string(r.([]byte)) != admin {
		return fmt.Errorf("CheckAdminToken error! impossible!")
	}

	return nil
}

func UserLogout(mobile, token string) error {
	return RemoveRedisKey(user + token + mobile)
}

func MasterLogout(mobile, token string) error {
	return RemoveRedisKey(master + token + mobile)
}

func AdminLogout(name, token string) error {
	return RemoveRedisKey(admin + token + name)
}

func RemoveRedisKey(key string) error {
	redis := db.RedisClient.Get()
	defer redis.Close()
	_, err := redis.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}
