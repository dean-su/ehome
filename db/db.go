package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	RedisClient *redis.Pool
)

func init() {
	var host, user, passwd, dbname string
	var port int
	var redisdb int

	beego.Info("start to init databases")
	host = beego.AppConfig.String("mysql.host")
	port, _ = beego.AppConfig.Int("mysql.port")
	user = beego.AppConfig.String("mysql.user")
	passwd = beego.AppConfig.String("mysql.passwd")
	dbname = beego.AppConfig.String("mysql.db")

	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", user, passwd, host, port, dbname)

	orm.RegisterDataBase("default", "mysql", connstr)

	connstr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", user, passwd, host, port, "ecs")
	orm.RegisterDataBase("ecs", "mysql", connstr)
	beego.Info("finish RegisterDataBase mysql")

	host = beego.AppConfig.String("redis.host")
	redisdb, _ = beego.AppConfig.Int("redis.db")

	RedisClient = &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			// ѡÔdb
			c.Do("SELECT", redisdb)
			return c, nil
		},
	}
	beego.Info("finish prepare redis pool")

	beego.Info("end init databases")
}
