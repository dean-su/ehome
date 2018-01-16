package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ehome/controllers/admins:AdminLoginController"] = append(beego.GlobalControllerRouter["ehome/controllers/admins:AdminLoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
