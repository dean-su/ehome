package controllers

import (
	"ehome/models"
	_ "errors"
	_ "strconv"
	"strings"

	"github.com/astaxie/beego"
)

// VoiceController oprations for EhomeTopic
type VoiceController struct {
	beego.Controller
}

// URLMapping ...
func (c *VoiceController) URLMapping() {
	c.Mapping("Upload", c.Upload)
}

// Post ...
// @Title Post
// @Description create EhomeTopic
// @Param	body		body 	models.EhomeTopic	true		"body for EhomeTopic content"
// @Success 201 {int} models.EhomeTopic
// @Failure 403 body is empty
// @router /upload [post]
func (c *VoiceController) Upload() {

	f, h, e := c.GetFile("File")
	if e == nil {
		f.Close()
	}

	var path string
	var ehomev models.EhomeVoice
	var id int64

	map2 := make(map[string]interface{})

	sl := strings.Split(h.Filename, ".")

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
		map2["Status"] = 3
		goto BOTTOM
	}

	ehomev.Filename = models.GetRandFileName(models.VoicePath, sl)
	path = models.ServerPath + ehomev.Filename
	beego.Info(path)

	e = c.SaveToFile("File", path)

	if e != nil {
		map2["Status"] = 4
		goto BOTTOM
	}
	id, e = models.AddEhomeVoice(&ehomev)

	beego.Info("name:", h.Filename, id)

	map2["Status"] = 0
	map2["Id"] = id

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
