package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ehome/controllers:AlipayController"] = append(beego.GlobalControllerRouter["ehome/controllers:AlipayController"],
		beego.ControllerComments{
			Method: "Notify",
			Router: `/notify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:AlipayController"] = append(beego.GlobalControllerRouter["ehome/controllers:AlipayController"],
		beego.ControllerComments{
			Method: "Getparam",
			Router: `/getparam`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:AlipayController"] = append(beego.GlobalControllerRouter["ehome/controllers:AlipayController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:AlipayController"] = append(beego.GlobalControllerRouter["ehome/controllers:AlipayController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:Appointment"] = append(beego.GlobalControllerRouter["ehome/controllers:Appointment"],
		beego.ControllerComments{
			Method: "Make",
			Router: `/make`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:BankController"] = append(beego.GlobalControllerRouter["ehome/controllers:BankController"],
		beego.ControllerComments{
			Method: "GetBankList",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:BannerController"] = append(beego.GlobalControllerRouter["ehome/controllers:BannerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"] = append(beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"],
		beego.ControllerComments{
			Method: "SendSms",
			Router: `/sendsms`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"] = append(beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"],
		beego.ControllerComments{
			Method: "RegisterInit",
			Router: `/init`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"] = append(beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"],
		beego.ControllerComments{
			Method: "RegisterRequest",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"] = append(beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"],
		beego.ControllerComments{
			Method: "ForgetInit",
			Router: `/forgetinit`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"] = append(beego.GlobalControllerRouter["ehome/controllers:EcsUsersController"],
		beego.ControllerComments{
			Method: "ForgetRequest",
			Router: `/forgetrequest`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ExperienceController"] = append(beego.GlobalControllerRouter["ehome/controllers:ExperienceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["ehome/controllers:FeedbackController"],
		beego.ControllerComments{
			Method: "Feedback",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:FixAddressController"] = append(beego.GlobalControllerRouter["ehome/controllers:FixAddressController"],
		beego.ControllerComments{
			Method: "Request",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:FixAddressController"] = append(beego.GlobalControllerRouter["ehome/controllers:FixAddressController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:FixAddressController"] = append(beego.GlobalControllerRouter["ehome/controllers:FixAddressController"],
		beego.ControllerComments{
			Method: "Modify",
			Router: `/modify`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:FixAddressController"] = append(beego.GlobalControllerRouter["ehome/controllers:FixAddressController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:FixAddressController"] = append(beego.GlobalControllerRouter["ehome/controllers:FixAddressController"],
		beego.ControllerComments{
			Method: "Default",
			Router: `/default`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ImageController"] = append(beego.GlobalControllerRouter["ehome/controllers:ImageController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:LoginController"] = append(beego.GlobalControllerRouter["ehome/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:LogoutController"] = append(beego.GlobalControllerRouter["ehome/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterController"],
		beego.ControllerComments{
			Method: "RegisterInit",
			Router: `/init`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterController"],
		beego.ControllerComments{
			Method: "RegisterRequest",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterbankController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterbankController"],
		beego.ControllerComments{
			Method: "Binding",
			Router: `/binding`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterbankController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterbankController"],
		beego.ControllerComments{
			Method: "ModifyBinding",
			Router: `/modifybinding`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterbankController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterbankController"],
		beego.ControllerComments{
			Method: "DeleteBinding",
			Router: `/deletebinding`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterbankController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterbankController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterbankController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterbankController"],
		beego.ControllerComments{
			Method: "Cash",
			Router: `/cash`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "RequestOrder",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "ConfirmOrder",
			Router: `/confirm`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "PricingOrder",
			Router: `/pricing`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "Setout",
			Router: `/setout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "Arrived",
			Router: `/arrived`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "FinishConstruction",
			Router: `/finishconstruction`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasterorderController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasterorderController"],
		beego.ControllerComments{
			Method: "ConfirmPaid",
			Router: `/confirmpaid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:MasteruserController"] = append(beego.GlobalControllerRouter["ehome/controllers:MasteruserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ehome/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ehome/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ehome/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ehome/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ObjectController"] = append(beego.GlobalControllerRouter["ehome/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetAllOrder",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOrderNum",
			Router: `/number`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "GetOrderPage",
			Router: `/page`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "PlaceOrder",
			Router: `/place`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Evaluate",
			Router: `/evaluate`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Paycash",
			Router: `/paycash`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "Cancel",
			Router: `/cancel`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "ConfirmPricing",
			Router: `/confirmpricing`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:OrderController"] = append(beego.GlobalControllerRouter["ehome/controllers:OrderController"],
		beego.ControllerComments{
			Method: "CheckAndAccept",
			Router: `/checkandaccept`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:PriceController"] = append(beego.GlobalControllerRouter["ehome/controllers:PriceController"],
		beego.ControllerComments{
			Method: "Request",
			Router: `/request`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:PriceController"] = append(beego.GlobalControllerRouter["ehome/controllers:PriceController"],
		beego.ControllerComments{
			Method: "Fixtype",
			Router: `/fixtype`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ProjSampleController"] = append(beego.GlobalControllerRouter["ehome/controllers:ProjSampleController"],
		beego.ControllerComments{
			Method: "GetAllProjSample",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ProjSampleController"] = append(beego.GlobalControllerRouter["ehome/controllers:ProjSampleController"],
		beego.ControllerComments{
			Method: "GetProjSampleNum",
			Router: `/number`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:ProjSampleController"] = append(beego.GlobalControllerRouter["ehome/controllers:ProjSampleController"],
		beego.ControllerComments{
			Method: "GetProjSamplePage",
			Router: `/page`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RedPackageController"] = append(beego.GlobalControllerRouter["ehome/controllers:RedPackageController"],
		beego.ControllerComments{
			Method: "Share",
			Router: `/share`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RedPackageController"] = append(beego.GlobalControllerRouter["ehome/controllers:RedPackageController"],
		beego.ControllerComments{
			Method: "Grad",
			Router: `/grad`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RedPackageController"] = append(beego.GlobalControllerRouter["ehome/controllers:RedPackageController"],
		beego.ControllerComments{
			Method: "Chance",
			Router: `/chance`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RedirectController"] = append(beego.GlobalControllerRouter["ehome/controllers:RedirectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RedirectController"] = append(beego.GlobalControllerRouter["ehome/controllers:RedirectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RegionController"] = append(beego.GlobalControllerRouter["ehome/controllers:RegionController"],
		beego.ControllerComments{
			Method: "GetProvince",
			Router: `/province`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RegionController"] = append(beego.GlobalControllerRouter["ehome/controllers:RegionController"],
		beego.ControllerComments{
			Method: "GetCity",
			Router: `/city`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:RegionController"] = append(beego.GlobalControllerRouter["ehome/controllers:RegionController"],
		beego.ControllerComments{
			Method: "GetRegion",
			Router: `/region`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:StyleController"] = append(beego.GlobalControllerRouter["ehome/controllers:StyleController"],
		beego.ControllerComments{
			Method: "GetStyle",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:StyleController"] = append(beego.GlobalControllerRouter["ehome/controllers:StyleController"],
		beego.ControllerComments{
			Method: "GetStylePage",
			Router: `/page`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:StyleController"] = append(beego.GlobalControllerRouter["ehome/controllers:StyleController"],
		beego.ControllerComments{
			Method: "GetStylePrice",
			Router: `/price`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:TestController"] = append(beego.GlobalControllerRouter["ehome/controllers:TestController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:TestController"] = append(beego.GlobalControllerRouter["ehome/controllers:TestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:TopicController"] = append(beego.GlobalControllerRouter["ehome/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Category",
			Router: `/category`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:TopicController"] = append(beego.GlobalControllerRouter["ehome/controllers:TopicController"],
		beego.ControllerComments{
			Method: "HotTopic",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:TopicController"] = append(beego.GlobalControllerRouter["ehome/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Click",
			Router: `/click`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:TosController"] = append(beego.GlobalControllerRouter["ehome/controllers:TosController"],
		beego.ControllerComments{
			Method: "GetTos",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:UploadController"] = append(beego.GlobalControllerRouter["ehome/controllers:UploadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:UserController"] = append(beego.GlobalControllerRouter["ehome/controllers:UserController"],
		beego.ControllerComments{
			Method: "UploadImage",
			Router: `/uploadimage`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:UserController"] = append(beego.GlobalControllerRouter["ehome/controllers:UserController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:UserController"] = append(beego.GlobalControllerRouter["ehome/controllers:UserController"],
		beego.ControllerComments{
			Method: "Changepass",
			Router: `/changepasswd`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:VersionController"] = append(beego.GlobalControllerRouter["ehome/controllers:VersionController"],
		beego.ControllerComments{
			Method: "Version",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:VersionController"] = append(beego.GlobalControllerRouter["ehome/controllers:VersionController"],
		beego.ControllerComments{
			Method: "IosVersion",
			Router: `/ios`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:VoiceController"] = append(beego.GlobalControllerRouter["ehome/controllers:VoiceController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["ehome/controllers:WebController"] = append(beego.GlobalControllerRouter["ehome/controllers:WebController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
