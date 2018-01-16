package trade

import (
	"ehome/models"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"time"
)

/*
app_id String 是 32 支付宝分配给开发者的应用ID 2014072300007148
method String 是 128 接口名称 alipay.trade.app.pay
format String 否 40 仅支持JSON JSON
charset String 是 10 请求使用的编码格式，如utf-8,gbk,gb2312等 utf-8
sign_type String 是 10 商户生成签名字符串所使用的签名算法类型，目前支持RSA RSA
sign String 是 256 商户请求参数的签名串，详见签名 详见示例
timestamp String 是 19 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss" 2014-07-24 03:07:50
version String 是 3 调用的接口版本，固定为：1.0 1.0
notify_url String 是 256 支付宝服务器主动通知商户服务器里指定的页面http/https路径。建议商户使用https https://api.xx.com/receive_notify.htm
biz_content String 是 - 业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
*/

/*
body String 否 128 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。 Iphone6 16G
subject String 是 256 商品的标题/交易标题/订单标题/订单关键字等。 大乐透
out_trade_no String 是 64 商户网站唯一订单号 70501111111S001111119
timeout_express String 否 6 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。 90m
total_amount String 是 9 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 9.00
seller_id String 否 16 收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID 2088102147948060
product_code String 是 64 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY QUICK_MSECURITY_PAY
*/

func Alipay(trade_no string) (signedstr string, err error) {
	var str string
	var m, j map[string]interface{}

	m = make(map[string]interface{})
	j = make(map[string]interface{})

	m["app_id"] = GetEhomeAppid()
	m["method"] = "alipay.trade.app.pay"
	m["format"] = "JSON"
	m["charset"] = "utf-8"
	m["sign_type"] = "RSA"
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["version"] = "1.0"
	m["notify_url"] = fmt.Sprintf("http://%s/v1/alipay/notify", models.Serverip)

	j["body"] = ""
	j["subject"] = fmt.Sprintf("order %s", trade_no)
	j["out_trade_no"] = trade_no
	j["timeout_express"] = "90m"
	var amount float64
	_, _, amount, err = models.GetOrderPrice(trade_no)
	if err != nil {
		return
	}
	j["total_amount"] = fmt.Sprintf("%f", amount)
	j["seller_id"] = ""
	j["product_code"] = "QUICK_MSECURITY_PAY"

	b, e := json.Marshal(j)
	if e != nil {
		err = fmt.Errorf("Marshal error! %v", e)
		return
	}

	m["biz_content"] = string(b)

	str, err = GetAlipaySignString(m)
	if err != nil {
		fmt.Printf("GetAlipaySignString error! %v", err)
		return
	}

	signedstr, err = Sign(str)
	if err != nil {
		fmt.Printf("Sign error! %v", e)
		return
	}
	fmt.Printf("url[%s] \n", str)
	fmt.Printf("sign[%s] \n", signedstr)

	str, err = GetEscapeString(m)
	signedstr = str + "&sign=" + url.QueryEscape(signedstr)
	fmt.Printf("signedstr:[%s] \n", signedstr)
	return
}

func GetAlipaySignString(mapBody map[string]interface{}) (sign string, err error) {
	sorted_keys := make([]string, 0)
	for k, _ := range mapBody {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	index := 0
	for _, k := range sorted_keys {
		//fmt.Println("k=", k, "v =", mapBody[k])
		value := fmt.Sprintf("%v", mapBody[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value
		} //最后一项后面不要&
		if index < len(sorted_keys)-1 {
			signStrings = signStrings + "&"
		}
		index++
	}
	return signStrings, nil
}

func GetEscapeString(mapBody map[string]interface{}) (sign string, err error) {
	sorted_keys := make([]string, 0)
	for k, _ := range mapBody {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	index := 0
	for _, k := range sorted_keys {
		//fmt.Println("k=", k, "v =", mapBody[k])
		value := fmt.Sprintf("%v", mapBody[k])
		if value != "" {
			signStrings = signStrings + k + "=" + url.QueryEscape(value)
		}
		if index < len(sorted_keys)-1 {
			signStrings = signStrings + "&"
		}
		index++
	}
	return signStrings, nil
}
