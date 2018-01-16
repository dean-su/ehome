package admins

import (
	"ehome/models"
	"github.com/astaxie/beego"
	"strconv"
)

type FrontpageController struct {
	beego.Controller
}

// URLMapping ...
func (c *FrontpageController) URLMapping() {
	c.Mapping("GetBanner", c.GetBanner)
	c.Mapping("AddBanner", c.AddBanner)
	c.Mapping("ModifyBanner", c.ModifyBanner)
	c.Mapping("DeleteBanner", c.DeleteBanner)
	c.Mapping("BannerNum", c.BannerNum)

	c.Mapping("GetTopicCategory", c.GetTopicCategory)
	c.Mapping("AddTopicCategory", c.AddTopicCategory)
	c.Mapping("ModifyTopicCategory", c.ModifyTopicCategory)
	c.Mapping("DeleteTopicCategory", c.DeleteTopicCategory)

	c.Mapping("GetTopic", c.GetTopic)
	c.Mapping("AddTopic", c.AddTopic)
	c.Mapping("ModifyTopic", c.ModifyTopic)
	c.Mapping("DeleteTopic", c.DeleteTopic)
	c.Mapping("TopicNum", c.TopicNum)

	c.Mapping("GetExperience", c.GetExperience)
	c.Mapping("AddExperience", c.AddExperience)
	c.Mapping("ModifyExperience", c.ModifyExperience)
	c.Mapping("DeleteExperience", c.DeleteExperience)
	c.Mapping("ExperienceNum", c.ExperienceNum)

}

// @param
// @Failure 403 body is empty
// @router /bannernum [post]
func (c *FrontpageController) BannerNum() {
	map2 := make(map[string]interface{})

	var m int64
	Type, _ := c.GetInt("Type")

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	m, err = models.GetBannerNum(Type)
	if err != nil {
		SetError(map2, DB_ERROR, "GetBannerNum error!")
		goto BOTTOM
	} else {
		map2["status"] = 0
		map2["Total"] = m
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /topicnum [post]
func (c *FrontpageController) TopicNum() {
	map2 := make(map[string]interface{})

	var m int64

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	m, err = models.GetTopicNum(1, "")

	if err != nil {
		SetError(map2, DB_ERROR, "GetTopicNum error!")
		goto BOTTOM
	} else {
		map2["status"] = 0
		map2["Total"] = m
	}

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /experiencenum [post]
func (c *FrontpageController) ExperienceNum() {
	map2 := make(map[string]interface{})
	var m int64

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	m, err = models.GetExperienceNum()

	if err != nil {
		SetError(map2, DB_ERROR, "GetExperienceNum error!")
		goto BOTTOM
	}
	map2["Total"] = m

	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /getbanner [post]
func (c *FrontpageController) GetBanner() {
	map2 := make(map[string]interface{})

	var m int64
	var Type int
	reqtype := c.GetString("Type")
	limit, offset := GetPage(c)

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if reqtype != "0" && reqtype != "1" && reqtype != "3" && reqtype != "4" {
		SetError(map2, PARAM_ERR, "param reqtype error! %s", reqtype)
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllBanner(reqtype, limit, offset)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllBanner error! %v", err)
		delete(map2, "records")
		goto BOTTOM
	}

	Type, _ = strconv.Atoi(reqtype)
	m, err = models.GetBannerNum(Type)
	if err != nil {
		SetError(map2, DB_ERROR, "GetBannerNum error!")
		goto BOTTOM
	}
	map2["Total"] = m

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAllBanner(reqtype string, limit, offset int64) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	query["Type"] = reqtype

	ml, err = models.GetAllEhomeBanner(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) <= 0 {
		ml = make([]interface{}, 0)
	}

	return
}

// @param
// @Failure 403 body is empty
// @router /addbanner [post]
func (c *FrontpageController) AddBanner() {
	beego.Error(c.Input())
	beego.Error(c.Ctx.Input.Header("Content-type"))
	map2 := make(map[string]interface{})

	var v models.EhomeBanner
	reqtype := c.GetString("Type")
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if reqtype != "0" && reqtype != "1" && reqtype != "3" && reqtype != "4" {
		SetError(map2, PARAM_ERR, "param reqtype error! %s", reqtype)
		goto BOTTOM
	}

	v.Href = c.GetString("Href")
	v.Type, _ = strconv.Atoi(reqtype)
	v.Img = c.GetString("Img")
	v.Status, err = c.GetInt16("Status")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Status invalid! [%s]", c.GetString("Status"))
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["Id"], err = models.AddEhomeBanner(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeBanner error! %v", err)
		delete(map2, "records")
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modifybanner [post]
func (c *FrontpageController) ModifyBanner() {
	map2 := make(map[string]interface{})

	var v models.EhomeBanner
	reqtype := c.GetString("Type")
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if reqtype != "0" && reqtype != "1" && reqtype != "3" && reqtype != "4" {
		SetError(map2, PARAM_ERR, "param reqtype error! %s", reqtype)
		goto BOTTOM
	}

	v.Href = c.GetString("Href")
	v.Type, _ = strconv.Atoi(reqtype)
	v.Img = c.GetString("Img")
	v.Status, err = c.GetInt16("Status")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Status invalid! [%s]", c.GetString("Status"))
		goto BOTTOM
	}
	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id invalid! [%s]", c.GetString("Id"))
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	err = models.UpdateEhomeBannerById(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeBannerById error! %v", err)
		delete(map2, "records")
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /deletebanner [delete]
func (c *FrontpageController) DeleteBanner() {
	map2 := make(map[string]interface{})

	var v models.EhomeBanner
	reqtype := c.GetString("Type")
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	if reqtype != "0" && reqtype != "1" && reqtype != "3" && reqtype != "4" {
		SetError(map2, PARAM_ERR, "param reqtype error! %s", reqtype)
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id invalid! [%s]", c.GetString("Id"))
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	err = models.DeleteEhomeBanner(v.Id)

	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeBann error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /gettopiccategory [post]
func (c *FrontpageController) GetTopicCategory() {
	map2 := make(map[string]interface{})
	limit, offset := GetPage(c)
	Type := c.GetString("Type")

	if Type == "" {
		Type = "0"
	}

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllTopicCategory(limit, offset, Type)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllTopicCategory error!")
		goto BOTTOM
	}

	map2["Total"], err = models.GetTopicCatNum(1, Type)

	if err != nil {
		SetError(map2, DB_ERROR, "GetTopicCatNum error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAllTopicCategory(limit, offset int64, Type string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	query["Type"] = Type

	ml, err = models.GetAllEhomeTopicCategory(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) == 0 {
		ml = make([]interface{}, 0)
	}

	return
}

// @param
// @Failure 403 body is empty
// @router /addtopiccategory [post]
func (c *FrontpageController) AddTopicCategory() {
	var v models.EhomeTopicCategory

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	v.Type, err = c.GetInt8("Type")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Type error! [%s]", err)
		goto BOTTOM
	}

	v.Title = c.GetString("Title")
	if len(v.Title) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Title)
		goto BOTTOM
	}

	v.CatImg = c.GetString("CatImg")
	if len(v.CatImg) == 0 {
		SetError(map2, PARAM_ERR, "param Img error! [%s]", v.CatImg)
		goto BOTTOM
	}

	v.Status, err = c.GetInt8("Status")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Status"))
		goto BOTTOM
	}

	map2["Id"], err = models.AddEhomeTopicCategory(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeTopicCategory error! %s", err)
		goto BOTTOM
	}
	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modifytopiccategory [post]
func (c *FrontpageController) ModifyTopicCategory() {
	var v models.EhomeTopicCategory

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	v.Type, err = c.GetInt8("Type")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Type error! [%v]", err)
		goto BOTTOM
	}

	v.Title = c.GetString("Title")
	if len(v.Title) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Title)
		goto BOTTOM
	}

	v.CatImg = c.GetString("CatImg")
	if len(v.CatImg) == 0 {
		SetError(map2, PARAM_ERR, "param Img error! [%s]", v.CatImg)
		goto BOTTOM
	}

	v.Status, err = c.GetInt8("Status")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Status"))
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.UpdateEhomeTopicCategoryById(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeTopicCategoryById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /deletetopiccategory [post]
func (c *FrontpageController) DeleteTopicCategory() {
	var Id int
	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.DeleteEhomeTopicCategory(Id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeTopicCategory error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /gettopic [post]
func (c *FrontpageController) GetTopic() {
	map2 := make(map[string]interface{})
	limit, offset := GetPage(c)

	Catid := c.GetString("Catid")
	if Catid == "" {
		Catid = "0"
	}

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllTopic(limit, offset, Catid)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllTopic error!")
		goto BOTTOM
	}

	map2["Total"], err = models.GetTopicNum(1, Catid)

	if err != nil {
		SetError(map2, DB_ERROR, "GetTopicNum error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAllTopic(limit, offset int64, Catid string) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	sortby = append(sortby, "CreateTime")
	order = append(order, "desc")

	if Catid != "" && Catid != "0" {
		query["TopicCatid"] = Catid
	}

	ml, err = models.GetAllEhomeTopic(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) == 0 {
		ml = make([]interface{}, 0)
	}

	return
}

// @param
// @Failure 403 body is empty
// @router /addtopic [post]
func (c *FrontpageController) AddTopic() {
	var v models.EhomeTopic

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	v.TopicCatid, err = c.GetInt("TopicCatid")
	if err != nil {
		SetError(map2, PARAM_ERR, "param TopicCatid error! [%v]", err)
		goto BOTTOM
	}

	v.Source = c.GetString("Source")
	if len(v.Source) == 0 {
		SetError(map2, PARAM_ERR, "param Source error! [%s]", v.Source)
		goto BOTTOM
	}

	v.Title = c.GetString("Title")
	if len(v.Title) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Title)
		goto BOTTOM
	}
	v.Desc = c.GetString("Desc")
	if len(v.Desc) == 0 {
		SetError(map2, PARAM_ERR, "param Desc error! [%s]", v.Desc)
		goto BOTTOM
	}
	v.Data = c.GetString("Data")
	if len(v.Data) == 0 {
		SetError(map2, PARAM_ERR, "param Data error! [%s]", v.Data)
		goto BOTTOM
	}

	v.Img = c.GetString("Img")
	if len(v.Img) == 0 {
		SetError(map2, PARAM_ERR, "param Img error! [%s]", v.Img)
		goto BOTTOM
	}

	v.Clicks, err = c.GetInt("Clicks")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Clicks"))
		goto BOTTOM
	}

	v.Status, err = c.GetInt16("Status")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Status error! [%v]", err)
		goto BOTTOM
	}

	map2["Id"], err = models.AddEhomeTopic(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeTopic error! %s", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modifytopic [post]
func (c *FrontpageController) ModifyTopic() {
	var v models.EhomeTopic

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	v.TopicCatid, err = c.GetInt("TopicCatid")
	if err != nil {
		SetError(map2, PARAM_ERR, "param TopicCatid error! [%v]", err)
		goto BOTTOM
	}

	v.Source = c.GetString("Source")
	if len(v.Source) == 0 {
		SetError(map2, PARAM_ERR, "param Source error! [%s]", v.Source)
		goto BOTTOM
	}

	v.Title = c.GetString("Title")
	if len(v.Title) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Title)
		goto BOTTOM
	}
	v.Desc = c.GetString("Desc")
	if len(v.Desc) == 0 {
		SetError(map2, PARAM_ERR, "param Desc error! [%s]", v.Desc)
		goto BOTTOM
	}
	v.Data = c.GetString("Data")
	if len(v.Data) == 0 {
		SetError(map2, PARAM_ERR, "param Data error! [%s]", v.Data)
		goto BOTTOM
	}

	v.Img = c.GetString("Img")
	if len(v.Img) == 0 {
		SetError(map2, PARAM_ERR, "param Img error! [%s]", v.Img)
		goto BOTTOM
	}

	v.Clicks, err = c.GetInt("Clicks")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Clicks"))
		goto BOTTOM
	}

	v.Status, err = c.GetInt16("Status")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Status error! [%v]", err)
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.UpdateEhomeTopicById(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeTopicById error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /deletetopic [post]
func (c *FrontpageController) DeleteTopic() {
	var Id int
	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.DeleteEhomeTopic(Id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeTopic error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /getexperience [post]
func (c *FrontpageController) GetExperience() {
	map2 := make(map[string]interface{})
	limit, offset := GetPage(c)
	var m int64

	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	map2["records"], err = GetAllExperience(limit, offset)
	if err != nil {
		SetError(map2, DB_ERROR, "GetAllExperience error!")
		goto BOTTOM
	}

	m, err = models.GetExperienceNum()

	if err != nil {
		SetError(map2, DB_ERROR, "GetExperienceNum error!")
		goto BOTTOM
	}
	map2["Total"] = m

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()

}

// @param
// @Failure 403 body is empty
// @router /addexperience [post]
func (c *FrontpageController) AddExperience() {
	var v models.EhomeExperience

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	v.Title = c.GetString("Title")
	if len(v.Title) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Title)
		goto BOTTOM
	}
	v.Image = c.GetString("Image")
	if len(v.Image) == 0 {
		SetError(map2, PARAM_ERR, "param Image error! [%s]", v.Image)
		goto BOTTOM
	}

	v.Body = c.GetString("Body")
	if len(v.Body) == 0 {
		SetError(map2, PARAM_ERR, "param Body error! [%s]", v.Body)
		goto BOTTOM
	}

	v.Clicks, err = c.GetInt("Clicks")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Clicks"))
		goto BOTTOM
	}

	map2["Id"], err = models.AddEhomeExperience(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeTopic error! %s", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /modifyexperience [post]
func (c *FrontpageController) ModifyExperience() {
	var v models.EhomeExperience

	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}
	v.Title = c.GetString("Title")
	if len(v.Title) == 0 {
		SetError(map2, PARAM_ERR, "param Title error! [%s]", v.Title)
		goto BOTTOM
	}
	v.Image = c.GetString("Image")
	if len(v.Image) == 0 {
		SetError(map2, PARAM_ERR, "param Image error! [%s]", v.Image)
		goto BOTTOM
	}

	v.Body = c.GetString("Body")
	if len(v.Body) == 0 {
		SetError(map2, PARAM_ERR, "param Body error! [%s]", v.Body)
		goto BOTTOM
	}

	v.Clicks, err = c.GetInt("Clicks")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Clicks error! [%s]", c.GetString("Clicks"))
		goto BOTTOM
	}

	v.Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.UpdateEhomeExperienceById(&v)

	if err != nil {
		SetError(map2, DB_ERROR, "UpdateEhomeExperienceById error! %s", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

// @param
// @Failure 403 body is empty
// @router /deleteexperience [delete]
func (c *FrontpageController) DeleteExperience() {
	var Id int
	map2 := make(map[string]interface{})
	name, token, reqtime, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	Id, err = c.GetInt("Id")
	if err != nil {
		SetError(map2, PARAM_ERR, "param Id error! [%s]", c.GetString("Id"))
		goto BOTTOM

	}

	err = models.DeleteEhomeExperience(Id)
	if err != nil {
		SetError(map2, DB_ERROR, "DeleteEhomeExperienc error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}

func GetAllExperience(limit, offset int64) (ml []interface{}, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)

	sortby = append(sortby, "Clicks")
	order = append(order, "desc")

	ml, err = models.GetAllEhomeExperience(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(ml) == 0 {
		ml = make([]interface{}, 0)
	}

	return
}
