package controllers

import (
	"ehome/db"
	"ehome/models"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	REDPACKAGE_BONUSPOINT = iota
	REDPACKAGE_BONUSMONEY
	REDPACKAGE_BONUSGIFT
	REDPACKAGE_BONUSDISCOUNT
)

type RedPackageController struct {
	beego.Controller
}

func (c *RedPackageController) URLMapping() {
	c.Mapping("Share", c.Share)
}

func get_gradrate() (rate float64, lower, upper int) {
	v, err := models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error! %v", err)
		return 0.1, 1, 1000
	}
	sp := strings.Split(v.Bonuspointrange, "-")
	if len(sp) != 2 {
		return float64(v.Bonuspointdifficulty) / float64(100), 1, 1000
	}

	lower, _ = strconv.Atoi(sp[0])
	upper, _ = strconv.Atoi(sp[1])
	if upper < lower {
		beego.Error("Bonuspointrange error! %v", err)
		upper = 1000
		lower = 1
	}

	return float64(v.Bonuspointdifficulty) / float64(100), lower, upper

}

func get_inittimes() (num int) {
	v, err := models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error! %v", err)
		return 3
	}

	return v.Redpackagetimesperday
}

func redkey(mobile string, usertype int) (key string) {
	key = fmt.Sprintf("%d%s%s", usertype, mobile, time.Now().Format("20060102"))
	beego.Error("key ", key)
	return
}

func redget(mobile string, usertype int) (num int, err error) {
	redis := db.RedisClient.Get()
	defer redis.Close()
	key := redkey(mobile, usertype)

	rv, err := redis.Do("GET", key)
	if rv == nil || err != nil {
		init_times := get_inittimes()
		rv, err = redis.Do("SET", key, init_times)
		if err != nil {
			beego.Error("redis SET error!")
			return
		}
		num = init_times
	} else {
		beego.Error("test %v", string(rv.([]byte)))
		num, err = strconv.Atoi(string(rv.([]byte)))
	}

	return
}

func redshare(mobile string, usertype int, n int) (num int, err error) {
	redis := db.RedisClient.Get()
	defer redis.Close()
	key := redkey(mobile, usertype)

	_, err = redis.Do("SET", key, n+1)
	if err != nil {
		return
	}
	/*
		num = int(rv.(int64))
		beego.Error("incrby ", key, num)
		redis.Do("EXEC")
	*/

	/*
		beego.Error("incrby ", string(rv.([]byte)))
		num, err = strconv.Atoi(string(rv.([]byte)))
	*/
	return
}

func redsub(mobile string, usertype int) (err error) {
	redis := db.RedisClient.Get()
	defer redis.Close()
	key := redkey(mobile, usertype)

	_, err = redis.Do("DECRBY", key, 1)

	return

}

// Share ...
// @Title share
// @Description share
// @Success 200 {object}
// @Failure 403 body is empty
// @router /share [get]
func (c *RedPackageController) Share() {
	map2 := make(map[string]interface{})
	var num int
	var v models.EhomeShareLog

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, PARAM_ERR, "user %s is invalid! %v", err)
		goto BOTTOM
	}

	v.Channel = c.GetString("Channel")
	if len(v.Channel) == 0 {
		SetError(map2, PARAM_ERR, "Channel is empty")
		goto BOTTOM
	}
	v.Usertype = int8(usertype)

	v.Userid, err = models.GetIdByNo(mobile, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "GetIdByNo error! %v", err)
		goto BOTTOM
	}

	num, err = redget(mobile, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "redget error! %v", err)
		goto BOTTOM
	}

	_, err = redshare(mobile, usertype, num)
	if err != nil {
		SetError(map2, DB_ERROR, "redshare error! %v", err)
		goto BOTTOM
	}

	models.AddEhomeShareLog(&v)

	/*
		err = redsub(mobile, usertype)
		if err != nil {
			SetError(map2, DB_ERROR, "redsub error!")
			goto BOTTOM
		}
	*/

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// grab redpackage
// @Title grad redpackage
// @Description grad redpackage
// @Success 200 {object}
// @Failure 403 body is empty
// @router /grad [get]
func (c *RedPackageController) Grad() {
	map2 := make(map[string]interface{})
	var num int
	var point int
	var v models.EhomeRedpackage

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, PARAM_ERR, "user %s is invalid! %v", err)
		goto BOTTOM
	}

	num, err = redget(mobile, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "redget error!%v", err)
		goto BOTTOM
	}

	if num <= 0 {
		SetError(map2, GRAB_TIME_EXHUAST, "grab time exhuast!")
		goto BOTTOM
	}

	//grab red package logic
	v.Usertype = int8(usertype)
	v.Userid, err = models.GetIdByNo(mobile, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "GetIdByNo error! %v", err)
		goto BOTTOM
	}

	v.Type = REDPACKAGE_BONUSPOINT

	point, err = GradLogic()
	if err != nil {
		SetError(map2, DB_ERROR, "GradLogic error!%v", err)
		goto BOTTOM
	}

	v.Bonuspoint = point

	err = redsub(mobile, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "redsub error!%v", err)
		goto BOTTOM
	}

	if v.Bonuspoint > 0 {
		_, err = models.AddEhomeRedpackage(&v)
		if err != nil {
			SetError(map2, DB_ERROR, "AddEhomeRedpackage error! %v", err)
			goto BOTTOM
		}

		if usertype == 2 {
			models.Updatemasterbonuspoint(v.Userid, v.Bonuspoint)
		}
	}

	map2["Type"] = REDPACKAGE_BONUSPOINT
	map2["Value"] = point
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// grab Chance
// @Title chance
// @Description chance
// @Success 200 {object}
// @Failure 403 body is empty
// @router /chance [get]
func (c *RedPackageController) Chance() {
	map2 := make(map[string]interface{})
	var num int

	mobile, token, reqtime, usertype, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = ChecComUser(mobile, token, reqtime, usertype)
	if err != nil {
		SetError(map2, PARAM_ERR, "user %s is invalid! %v", err)
		goto BOTTOM
	}

	num, err = redget(mobile, usertype)
	if err != nil {
		SetError(map2, DB_ERROR, "redget error!%v", err)
		goto BOTTOM
	}

	map2["Chance"] = num
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GradLogic() (point int, err error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	grad_rate, lower, upper := get_gradrate()

	beego.Error("error", rnd.Uint32()%1000, uint32(float64(1000)*grad_rate))
	if grad_rate < float64(1) && rnd.Uint32()%1000 > uint32(float64(1000)*float64(grad_rate)) {
		r := upper - lower
		point = int(rnd.Uint32() % uint32(r))
		point += lower
	} else {
		point = 0
	}
	return
}
