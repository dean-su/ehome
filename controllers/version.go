package controllers

import (
	"ehome/models"
	"fmt"
	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

func (c *VersionController) URLMapping() {
	c.Mapping("Version", c.Version)
	c.Mapping("IosVersion", c.IosVersion)
}

// GetAllBank ...
// @Title Get all bank
// @Description get all bank
// @Success 200 {object}
// @Failure 403 body is empty
// @router / [get]
func (c *VersionController) Version() {
	Type := c.GetString("Type")
	map2 := make(map[string]interface{})

	if Type == "1" {
		map2["Chksum"] = "56754579"
		map2["Download"] = fmt.Sprintf("http://%s/static/download/ehome.apk", models.Serverip)
		map2["Version"] = "1.0.5"
	} else {
		map2["Chksum"] = "56754579"
		map2["Download"] = fmt.Sprintf("http://%s/static/download/emaster.apk", models.Serverip)
		map2["Version"] = "1.0.5"
	}

	map2["Notice"] = "为了能够更好使用App请立即更新!"
	map2["status"] = 0

	c.Data["json"] = map2
	c.ServeJSON()
}

// @Success 200 {object}
// @Failure 403 body is empty
// @router /ios [get]
func (c *VersionController) IosVersion() {
	Type := c.GetString("Type")

	map2 := make(map[string]interface{})

	if Type == "1" {
		map2["Chksum"] = "56754579"
		map2["Download"] = fmt.Sprintf("http://%s/static/download/ios", models.Serverip)
		map2["Version"] = "1.2"
	} else {
		map2["Chksum"] = "56754579"
		map2["Download"] = fmt.Sprintf("http://%s/static/download/ios", models.Serverip)
		map2["Version"] = "1.2"
	}
	map2["Notice"] = "为了能够更好使用App请立即更新!"
	map2["status"] = 0

	c.Data["json"] = map2
	c.ServeJSON()
}
