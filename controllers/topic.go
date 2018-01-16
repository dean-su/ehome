package controllers

import (
	"ehome/models"
	"strconv"

	"github.com/astaxie/beego"
)

// TopicController oprations for EhomeTopic
type TopicController struct {
	beego.Controller
}

// URLMapping ...
func (c *TopicController) URLMapping() {
	c.Mapping("HotTopic", c.HotTopic)
	c.Mapping("Category", c.Category)
	c.Mapping("Click", c.Click)
}

func AllCategory(limit, offset int64, Type string) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	fields = append(fields, "Id", "Title", "CatImg", "Status")
	query["Status"] = "1"
	query["Type"] = Type
	m, err = models.GetAllEhomeTopicCategory(query, fields, sortby, order, offset, limit)
	if m == nil {
		m = make([]interface{}, 0)
	}

	return
}

// Category...
// @Title Category
// @Description get Category
// @Param	num  num string	false	"the max topics numbers"
// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router /category [get]
func (c *TopicController) Category() {
	var err error
	limit, offset := GetPage(c)
	map2 := make(map[string]interface{})
	Type := c.GetString("Type")
	if Type == "" {
		Type = "0"
	}

	map2["records"], err = AllCategory(limit, offset, Type)
	if err != nil {
		SetError(map2, DB_ERROR, "AllCategory error! %v", err)
		goto BOTTOM
	}
	map2["Total"], err = models.GetTopicCatNum(0, Type)
	if err != nil {
		SetError(map2, DB_ERROR, "GetTopicCatNum error! %v", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()

}

func AllTopic(limit, offset int64, Id string) (m []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	sortby = append(sortby, "CreateTime")
	order = append(order, "desc")

	fields = append(fields, "Id", "Title", "Img", "Desc", "Clicks", "Source")
	query["Status"] = "1"
	if Id != "" && Id != "0" {
		query["TopicCatid"] = Id
	}

	m, err = models.GetAllEhomeTopic(query, fields, sortby, order, offset, limit)
	if m == nil {
		m = make([]interface{}, 0)
	}

	return
}

// HotTopic ...
// @Title HotTopic
// @Description get HotTopic
// @Param	num  num string	false	"the max topics numbers"
// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router / [get]
func (c *TopicController) HotTopic() {
	var err error
	map2 := make(map[string]interface{})
	limit, offset := GetPage(c)
	Id := c.GetString("Catid")

	map2["records"], err = AllTopic(limit, offset, Id)
	if err != nil {
		SetError(map2, DB_ERROR, "AllTopic error! %v", err)
		goto BOTTOM
	}

	map2["Total"], err = models.GetTopicNum(0, Id)
	if err != nil {
		SetError(map2, DB_ERROR, "GetTopicNum error! %v", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// Clicks ...
// @Title Clicks
// @Description get Clicks
// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router /click [get]
func (c *TopicController) Click() {
	var err error
	map2 := make(map[string]interface{})
	Id := c.GetString("Id")
	var i int
	var v *models.EhomeTopic

	if Id == "" {
		SetError(map2, PARAM_ERR, "Id is empty")
		goto BOTTOM
	}

	err = models.UpdateTopicClick(Id)
	if err != nil {
		SetError(map2, DB_ERROR, "Update Click error! %v", err)
		goto BOTTOM
	}
	i, _ = strconv.Atoi(Id)
	v, err = models.GetEhomeTopicById(i)
	if err != nil {
		map2["Title"] = "出错了"
		map2["Content"] = "数据暂时没法访问"
		map2["status"] = 1
		goto BOTTOM
	}

	map2["Title"] = v.Title
	map2["Content"] = v.Data

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
