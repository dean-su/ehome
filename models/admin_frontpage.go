package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func GetBannerNum(Type int) (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeBanner
	sql := fmt.Sprintf("select count(1) as num from  %s where type=%d", v.TableName(), Type)
	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}

/*
func GetTopicNum() (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeTopic
	sql := fmt.Sprintf("select count(1) as num from  %s", v.TableName())
	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}
*/

func GetExperienceNum() (num int64, err error) {
	o := orm.NewOrm()
	var v EhomeExperience
	sql := fmt.Sprintf("select count(1) as num from  %s", v.TableName())
	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}
