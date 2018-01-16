package controllers

import (
	"ehome/models"

	"github.com/astaxie/beego"
)

// Appointment oprations for EhomeTopic
type Appointment struct {
	beego.Controller
}

// URLMapping ...
func (c *Appointment) URLMapping() {
	c.Mapping("Make", c.Make)
}

// Make appointment
// @Title Make appointment
// @Description Make appointment
// @Success 200 {object} models.EhomeTopic
// @Failure 403
// @router /make [get]
func (c *Appointment) Make() {

	var err error

	var v models.EhomeAppointment
	map2 := make(map[string]interface{})

	v.Name = c.GetString("Name")
	v.Phone = c.GetString("Phone")
	v.Address = c.GetString("Address")
	v.Content = c.GetString("Content")

	if v.Name == "" {
		SetError(map2, PARAM_ERR, "PARAM NAME is empty!")
		goto BOTTOM
	}

	if v.Phone == "" {
		SetError(map2, PARAM_ERR, "PARAM NAME is empty!")
		goto BOTTOM
	}

	if v.Address == "" {
		SetError(map2, PARAM_ERR, "PARAM NAME is empty!")
		goto BOTTOM
	}

	_, err = models.AddEhomeAppointment(&v)
	if err != nil {
		SetError(map2, DB_ERROR, "AddEhomeAppointment error! %v", err)
		goto BOTTOM
	}

	map2["status"] = 0

BOTTOM:
	c.Data["jsonp"] = map2
	c.ServeJSONP()
}
