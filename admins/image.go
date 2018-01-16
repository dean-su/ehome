package admins

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
// @Description create
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
	var ehomei models.EhomeAdminImage
	var id int64

	map2 := make(map[string]interface{})
	var sl []string

	name, token, reqtime, err := GetComUser(c, map2)
	if e != nil {
		SetError(map2, PARAM_ERR, "Param File is invalid!%s", e)
		goto BOTTOM
	}
	sl = strings.Split(h.Filename, ".")

	if err != nil {
		goto BOTTOM
	}

	err = models.CheckAdminUser(name, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "admin user %s is invalid", name)
		goto BOTTOM
	}

	ehomei.Filename = models.GetRandFileName(models.ImagePath, sl)
	path = models.ServerPath + ehomei.Filename
	beego.Info(path)

	e = c.SaveToFile("File", path)

	if e != nil {
		SetError(map2, SYS_ERROR, "SaveToFile error! [%s]", path)
		goto BOTTOM
	}
	id, e = models.AddEhomeAdminImage(&ehomei)

	beego.Info("name:", h.Filename, id)

	map2["status"] = 0
	map2["Id"] = id
	map2["Url"] = ehomei.Filename
	map2["link"] = ehomei.Filename

BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON()
}
