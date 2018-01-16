package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ehome/admins:AdminLoginController"] = append(beego.GlobalControllerRouter["ehome/admins:AdminLoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:AdminLogoutController"] = append(beego.GlobalControllerRouter["ehome/admins:AdminLogoutController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "BannerNum",
			Router: `/bannernum`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "TopicNum",
			Router: `/topicnum`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "ExperienceNum",
			Router: `/experiencenum`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "GetBanner",
			Router: `/getbanner`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "AddBanner",
			Router: `/addbanner`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "ModifyBanner",
			Router: `/modifybanner`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "DeleteBanner",
			Router: `/deletebanner`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "GetTopicCategory",
			Router: `/gettopiccategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "AddTopicCategory",
			Router: `/addtopiccategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "ModifyTopicCategory",
			Router: `/modifytopiccategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "DeleteTopicCategory",
			Router: `/deletetopiccategory`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "GetTopic",
			Router: `/gettopic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "AddTopic",
			Router: `/addtopic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "ModifyTopic",
			Router: `/modifytopic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "DeleteTopic",
			Router: `/deletetopic`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "GetExperience",
			Router: `/getexperience`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "AddExperience",
			Router: `/addexperience`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "ModifyExperience",
			Router: `/modifyexperience`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:FrontpageController"] = append(beego.GlobalControllerRouter["ehome/admins:FrontpageController"],
		beego.ControllerComments{
			Method: "DeleteExperience",
			Router: `/deleteexperience`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:ImageController"] = append(beego.GlobalControllerRouter["ehome/admins:ImageController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "GetMaster",
			Router: `/get`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/detail`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "GetRegion",
			Router: `/region`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "GetAuditPending",
			Router: `/auditpending`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "Audit",
			Router: `/audit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "GetCashPending",
			Router: `/cashpending`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "DealCash",
			Router: `/dealcash`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:MasterController"] = append(beego.GlobalControllerRouter["ehome/admins:MasterController"],
		beego.ControllerComments{
			Method: "Modify",
			Router: `/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "GetRegion",
			Router: `/region`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "GetStatusList",
			Router: `/statuslist`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "OrderList",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/detail`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Audit",
			Router: `/audit`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "DealCash",
			Router: `/dealcash`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "PlaceOrder",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Modify",
			Router: `/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Requestmasterlist",
			Router: `/requestmasterlist`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Assignmaster",
			Router: `/assignmaster`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:OrderController"] = append(beego.GlobalControllerRouter["ehome/admins:OrderController"],
		beego.ControllerComments{
			Method: "Modifystatusandprice",
			Router: `/modifystatusandprice`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:ProjsampleController"] = append(beego.GlobalControllerRouter["ehome/admins:ProjsampleController"],
		beego.ControllerComments{
			Method: "Num",
			Router: `/num`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:ProjsampleController"] = append(beego.GlobalControllerRouter["ehome/admins:ProjsampleController"],
		beego.ControllerComments{
			Method: "GetProjsample",
			Router: `/get`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:ProjsampleController"] = append(beego.GlobalControllerRouter["ehome/admins:ProjsampleController"],
		beego.ControllerComments{
			Method: "AddProjsample",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:ProjsampleController"] = append(beego.GlobalControllerRouter["ehome/admins:ProjsampleController"],
		beego.ControllerComments{
			Method: "ModifyProjsample",
			Router: `/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:ProjsampleController"] = append(beego.GlobalControllerRouter["ehome/admins:ProjsampleController"],
		beego.ControllerComments{
			Method: "DeleteProjsample",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "GetAdminlist",
			Router: `/getadminlist`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "ModifyAdmin",
			Router: `/modifyadmin`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "AddAdmin",
			Router: `/addadmin`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "DelAdmin",
			Router: `/deladmin`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "GetSetting",
			Router: `/getsetting`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "ModifySetting",
			Router: `/modifysetting`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:SettingController"] = append(beego.GlobalControllerRouter["ehome/admins:SettingController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:UserController"] = append(beego.GlobalControllerRouter["ehome/admins:UserController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:UserController"] = append(beego.GlobalControllerRouter["ehome/admins:UserController"],
		beego.ControllerComments{
			Method: "Infobyno",
			Router: `/infobyno`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:UserController"] = append(beego.GlobalControllerRouter["ehome/admins:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/admins:UserController"] = append(beego.GlobalControllerRouter["ehome/admins:UserController"],
		beego.ControllerComments{
			Method: "Modify",
			Router: `/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
