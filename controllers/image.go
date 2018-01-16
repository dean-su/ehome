package controllers

import (
	"ehome/models"
	_ "errors"
	_ "strconv"
	"strings"

	"github.com/astaxie/beego"
)

// ImageController oprations for EhomeTopic
type ImageController struct {
	beego.Controller
}

// URLMapping ...
func (c *ImageController) URLMapping() {
	c.Mapping("Upload", c.Upload)
}

// Post ...
// @Title Post
// @Description create EhomeTopic
// @Param	body		body 	models.EhomeTopic	true		"body for EhomeTopic content"
// @Success 201 {int} models.EhomeTopic
// @Failure 403 body is empty
// @router /upload [post]
func (c *ImageController) Upload() {

	f, h, e := c.GetFile("File")
	if e == nil {
		f.Close()
	}

	var path string
	var ehomei models.EhomeImage
	var id int64

	map2 := make(map[string]interface{})

	sl := strings.Split(h.Filename, ".")

	/*
		mobile := c.GetString("Mobileno")
		token := c.GetString("Token")
		reqtime, err := c.GetInt64("Reqtime")

		if err != nil {
			map2["Status"] = 1
			goto BOTTOM
		}

		_, err = models.ValidateUser(mobile, token, reqtime)
		if err != nil {
			map2["Status"] = 2
			goto BOTTOM
		}

		if len(sl) < 2 {
			map2["Status"] = 1
			goto BOTTOM
		}
	*/

	ehomei.Filename = models.GetRandFileName(models.ImagePath, sl)
	path = models.ServerPath + ehomei.Filename
	beego.Info(path)

	e = c.SaveToFile("File", path)

	if e != nil {
		map2["Status"] = 2
		goto BOTTOM
	}
	id, e = models.AddEhomeImage(&ehomei)

	beego.Info("name:", h.Filename, id)

	map2["Status"] = 0
	map2["Id"] = id
	map2["Url"] = ehomei.Filename

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
