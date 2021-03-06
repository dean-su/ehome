// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ehome/admins"
	"ehome/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/hottopics",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/projectsample",
			beego.NSInclude(
				&controllers.ProjSampleController{},
			),
		),

		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.EcsUsersController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),
		beego.NSNamespace("/version",
			beego.NSInclude(
				&controllers.VersionController{},
			),
		),
		beego.NSNamespace("/price",
			beego.NSInclude(
				&controllers.PriceController{},
			),
		),
		beego.NSNamespace("/voice",
			beego.NSInclude(
				&controllers.VoiceController{},
			),
		),
		beego.NSNamespace("/image",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
		beego.NSNamespace("/fixaddress",
			beego.NSInclude(
				&controllers.FixAddressController{},
			),
		),
		beego.NSNamespace("/upload",
			beego.NSInclude(
				&controllers.UploadController{},
			),
		),
		beego.NSNamespace("/masterreg",
			beego.NSInclude(
				&controllers.MasterController{},
			),
		),
		beego.NSNamespace("/tos",
			beego.NSInclude(
				&controllers.TosController{},
			),
		),
		beego.NSNamespace("/region",
			beego.NSInclude(
				&controllers.RegionController{},
			),
		),
		beego.NSNamespace("/banner",
			beego.NSInclude(
				&controllers.BannerController{},
			),
		),
		beego.NSNamespace("/experience",
			beego.NSInclude(
				&controllers.ExperienceController{},
			),
		),
		beego.NSNamespace("/alipay",
			beego.NSInclude(
				&controllers.AlipayController{},
			),
		),
		beego.NSNamespace("/redirect",
			beego.NSInclude(
				&controllers.RedirectController{},
			),
		),
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),
		beego.NSNamespace("/bank",
			beego.NSInclude(
				&controllers.BankController{},
			),
		),
		beego.NSNamespace("/style",
			beego.NSInclude(
				&controllers.StyleController{},
			),
		),
		beego.NSNamespace("/feedback",
			beego.NSInclude(
				&controllers.FeedbackController{},
			),
		),
		beego.NSNamespace("/appointment",
			beego.NSInclude(
				&controllers.Appointment{},
			),
		),
		beego.NSNamespace("/master",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.MasteruserController{},
				),
			),
			beego.NSNamespace("/order",
				beego.NSInclude(
					&controllers.MasterorderController{},
				),
			),
			beego.NSNamespace("/bank",
				beego.NSInclude(
					&controllers.MasterbankController{},
				),
			),
		),
		beego.NSNamespace("/admin",
			beego.NSNamespace("/login",
				beego.NSInclude(
					&admins.AdminLoginController{},
				),
			),
			beego.NSNamespace("/logout",
				beego.NSInclude(
					&admins.AdminLogoutController{},
				),
			),
			beego.NSNamespace("/image",
				beego.NSInclude(
					&admins.ImageController{},
				),
			),
			beego.NSNamespace("/frontpage",
				beego.NSInclude(
					&admins.FrontpageController{},
				),
			),
			beego.NSNamespace("/projsample",
				beego.NSInclude(
					&admins.ProjsampleController{},
				),
			),
			beego.NSNamespace("/master",
				beego.NSInclude(
					&admins.MasterController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&admins.UserController{},
				),
			),
			beego.NSNamespace("/order",
				beego.NSInclude(
					&admins.OrderController{},
				),
			),
			beego.NSNamespace("/setting",
				beego.NSInclude(
					&admins.SettingController{},
				),
			),
		),
		beego.NSNamespace("/redpackage",
			beego.NSInclude(
				&controllers.RedPackageController{},
			),
		),
		beego.NSNamespace("/web",
			beego.NSInclude(
				&controllers.WebController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
