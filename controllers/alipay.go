package controllers

import (
	"ehome/models"
	"ehome/trade"
	"github.com/astaxie/beego"
	"sort"
	"strconv"
	"strings"
)

type AlipayController struct {
	beego.Controller
}

// URLMapping ...
func (c *AlipayController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
	c.Mapping("Getparam", c.Getparam)
	c.Mapping("Notify", c.Notify)
}

// @Failure 403 body is empty
// @router /notify [post]
func (c *AlipayController) Notify() {
	var incomelog models.EhomeIncomeLog
	var notify models.EhomeNotifyLog
	var price float64
	var amount float64

	notify.Ipaddr = c.Ctx.Input.IP()
	notify.Uri = c.Ctx.Input.URI()
	notify.Body = string(c.Ctx.Input.RequestBody)

	params := c.Ctx.Request.Form
	pp := map[string]string{}
	for k, v := range params {
		if len(v) > 0 && k != "sign" && k != "sign_type" {
			pp[k] = v[0]
		}
	}

	keys := []string{}
	for k, _ := range pp {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	requestStr := []string{}
	for _, k := range keys {
		requestStr = append(requestStr, k+"="+pp[k])
	}

	success, err := trade.Verify(strings.Join(requestStr, "&"), params["sign"][0])
	if value, exsit := params["sign"]; !exsit || len(value) < 0 {
		beego.Error("notify no sign")
		goto BOTTOM
	}
	if err != nil || !success {
		beego.Error("Verify sign error!")
		c.Ctx.WriteString("fail")
		goto BOTTOM
	}
	notify.Verifystatus = 1

	if params["trade_status"][0] != "TRADE_SUCCESS" && params["trade_status"][0] != "TRADE_FINISHED" {
		beego.Error("trade fail", c.Ctx.Request.RequestURI, c.Ctx.Input.RequestBody)
		c.Ctx.WriteString("fail")
		goto BOTTOM
	}
	notify.Orderstatus = 1
	beego.Error("Verify sign success!")

	//check sellerid, check apiid, check out_trade_no, check total_amount
	if params["app_id"][0] != trade.GetEhomeAppid() {
		c.Ctx.WriteString("fail")
		goto BOTTOM
	}

	incomelog.Masterid, incomelog.Orderid, price, err = models.GetOrderPrice(params["out_trade_no"][0])
	if err != nil {
		beego.Error("out_trade_no not exists!", params["out_trade_no"][0])
		c.Ctx.WriteString("fail")
		goto BOTTOM
	}

	if incomelog.Masterid == 0 {
		beego.Error("Masterid not exists!")
		goto BOTTOM
	}

	amount, err = strconv.ParseFloat(params["total_amount"][0], 3)
	if err != nil {
		beego.Error("total_amount is not float!", params["total_amount"][0])
		c.Ctx.WriteString("total_amount is not float")
		goto BOTTOM
	}

	if !models.IsFloatEqual(price, amount) {
		beego.Error("order price not equal", price, params["total_amount"][0])
		c.Ctx.WriteString("fail")
		goto BOTTOM
	}
	incomelog.Orderno = params["out_trade_no"][0]
	incomelog.Price = price
	incomelog.Dividerate = models.GetSettlementrate(price)
	incomelog.Masteramount = incomelog.Dividerate * price
	incomelog.Platformamount = price - incomelog.Masteramount

	_, err = models.AddEhomeIncomeLog(&incomelog)
	if err != nil {
		beego.Error("AddEhomeIncomeLog error!")
	}

	err = models.AddIncome(&incomelog)
	if err != nil {
		beego.Error("AddIncome error!")
	}

	notify.Pricestatus = 1
	c.Ctx.WriteString("success")

BOTTOM:
	_, err = models.AddEhomeNotifyLog(&notify)
	if err != nil {
		beego.Error("AddEhomeNotifyLog error!")
	}

}

// @Failure 403 body is empty
// @router /getparam [get]
func (c *AlipayController) Getparam() {

	map2 := make(map[string]interface{})
	var signedstr string
	var err error
	var userid int
	tradeno := c.GetString("Orderno")

	mobile, token, reqtime, _, err := GetComUser(c, map2)
	if err != nil {
		goto BOTTOM
	}

	userid, err = models.ValidateUser(mobile, token, reqtime)
	if err != nil {
		SetError(map2, INVALID_USER, "user %s is not valid! [%v]", mobile, err)
		goto BOTTOM
	}
	beego.Error("userid = ", userid)

	if len(tradeno) == 0 {
		SetError(map2, PARAM_ERR, "tradeno error! %s", tradeno)
		goto BOTTOM
	}

	signedstr, err = trade.Alipay(tradeno)
	if err != nil {
		SetError(map2, PARAM_ERR, "trade.Alipay error! %v", err)
		goto BOTTOM
	}

	map2["Payparam"] = signedstr
	map2["status"] = 0
BOTTOM:
	c.Data["json"] = map2
	c.ServeJSON(true)
}

// @param   num   num false
// @Failure 403 body is empty
// @router / [get]
func (c *AlipayController) Get() {

	beego.Error("alipay Get Data")
	beego.Error(c.Input())

	beego.Error("alipay Get Body")
	beego.Error(string(c.Ctx.Input.RequestBody))
}

// @param   num   num false
// @Failure 403 body is empty
// @router / [post]
func (c *AlipayController) Post() {

	//c.Ctx.Request.Form
	beego.Error("alipay post Data")
	beego.Error(c.Ctx.Input.Data)

	beego.Error("alipay post Body")
	beego.Error(c.Ctx.Input.RequestBody)
}
