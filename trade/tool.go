package trade

import (
	"ehome/models"
	"github.com/astaxie/beego"
)

func GetEhomeAppid() (appid string) {
	v, err := models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error!")
		return
	}
	appid = v.Clientappid
	return
}

func GetEmasterAppid() (appid string) {
	v, err := models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error!")
		return
	}
	appid = v.Masterappid
	return
}

func GetClientPrivatekey() (key []byte, err error) {
	var v *models.EhomeSetting
	v, err = models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error!")
		return
	}
	key = []byte(v.Clientprivatekey)

	return
}

func GetClientPublickey() (key []byte, err error) {
	var v *models.EhomeSetting
	v, err = models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error!")
		return
	}

	key = []byte(v.Clientpublickey)

	return

}

func GetMasterPrivatekey() (key []byte, err error) {
	var v *models.EhomeSetting
	v, err = models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error!")
		return
	}
	key = []byte(v.Masterprivatekey)

	return
}

func GetMasterPublickey() (key []byte, err error) {
	var v *models.EhomeSetting
	v, err = models.GetEhomeSetting()
	if err != nil {
		beego.Error("GetEhomeSetting error!")
		return
	}

	key = []byte(v.Masterpublickey)

	return

}
