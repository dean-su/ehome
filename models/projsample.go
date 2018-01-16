package models

import (
	_ "errors"
	"fmt"
	_ "strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//OrderList
func ProjSampleList(limit, offset int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	sql := "select thumbimg, intro, city, style, HousingLayout, HouseSize, visible from  ehome_projsample order by create_time desc"
	if limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, limit)
	}
	if offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, offset)
	}
	beego.Info("sql ", sql)

	var Imgs []string
	var Desc []string
	var City []string
	var Style []string
	var HousingLayout []string
	var HouseSize []string
	var visible []int

	var num int64

	num, err = o.Raw(sql).QueryRows(&Imgs, &Desc, &City, &Style, &HousingLayout, &HouseSize, &visible)
	if err != nil {
		return nil, err
	}

	for i := int64(0); i < num; i++ {
		m := make(map[string]interface{})
		m["Img"] = Imgs[i]
		m["Desc"] = Desc[i]
		if visible[i] == 1 {
			m["City"] = City[i]
			m["Style"] = Style[i]
			m["HousingLayout"] = HousingLayout[i]
			m["HouseSize"] = HouseSize[i]
		} else {
			m["City"] = ""
			m["Style"] = ""
			m["HousingLayout"] = ""
			m["HouseSize"] = ""
		}
		ml = append(ml, m)
	}
	if ml == nil {
		ml = make([]interface{}, 0)
	}

	return
}

func ProjSampleNum() (num int64, err error) {
	o := orm.NewOrm()
	sql := "select count(1) as num from ehome_projsample"
	var nums []int64
	num, err = o.Raw(sql).QueryRows(&nums)
	if err != nil {
		return
	}
	num = nums[0]
	return
}
