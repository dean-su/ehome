package main

import (
	_ "ehome/db"
	_ "ehome/routers"


	"ehome/controllers"
	_ "encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	/*
		data := make(map[string]interface{})
		data["content"] = "page not found"
		data["status"] = 1
		//b, _ := json.Marshal(data)
	*/
	beego.Info("here")

	//rw.Write(b)
	rw.Write(([]byte("Hello")))
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})

	var Filter = func(ctx *context.Context) {
		beego.Error("uri", ctx.Request.RequestURI)
		beego.Error("body", "[", string(ctx.Input.RequestBody), "]")
	}
	beego.InsertFilter("/", beego.BeforeRouter, Filter)
	beego.InsertFilter("/*", beego.BeforeRouter, Filter)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.SetStaticPath("/html", "html")
	beego.Info("before go run start")
	//beego.ErrorHandler("404", page_not_found)
	beego.Run()
}
