package models

import (
	"crypto/md5"

	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"reflect"
	"strings"
	"time"
)

var (
	ServerPath string
	VoicePath  string
	ImagePath  string
	Serverip   string
)

func init() {
	ServerPath = beego.AppPath
	VoicePath = beego.AppConfig.String("common.voicepath")
	ImagePath = beego.AppConfig.String("common.imagepath")
	Serverip = beego.AppConfig.String("Serverip")
}

/*
type ReturnSms struct {
	Message string `xml:"message"`
}
*/
type ReturnSms struct {
	Code string `xml:"code"`
	Msg  string `xml:"msg"`
}

func SendSms(mobile string, content string) (err error) {

	/*
		curtime := time.Now().Format("20060102150405")

		str := string("易居乐筑yjlz789") + curtime
		sign := md5.Sum([]byte(str))
		signstr := hex.EncodeToString(sign[0:16])

		resp, e := http.PostForm("http://211.147.242.161:8888/v2sms.aspx",
			url.Values{"action": {"send"}, "userid": {"5824"}, "timestamp": {curtime},
				"sign": {signstr}, "mobile": {mobile}, "content": {content}, "sendTime": {""}, "extno": {""}})
		if e != nil {
			return e
		}

		body, err := ioutil.ReadAll(resp.Body)
		beego.Info("body:" + string(body[0:len(body)]))

		var result ReturnSms
		err = xml.Unmarshal(body, &result)
		if err != nil || result.Message != "ok" {
			return fmt.Errorf("send sms error!%s", result.Message)
		}
	*/

	curtime := fmt.Sprintf("%d", time.Now().Unix())
	beego.Info("curtime :", curtime)
	account := "cf_gzyjl"

	str := account + "d365bcb28a19f1215f13449fb466cb44" + mobile +
		content + curtime
	sign := md5.Sum([]byte(str))
	signstr := hex.EncodeToString(sign[0:16])

	//method=Submit&content=&account=cf_gzyjl&password=0584af602d5114c96b53305f14316837&mobile=13450249489&time=1478612988
	resp, e :=
		http.PostForm("http://106.ihuyi.com/webservice/sms.php?method=Submit",
			url.Values{"method": {"Submit"}, "content": {content},
				"time": {curtime}, "account": {account},
				"password": {signstr}, "mobile": {mobile}})
	if e != nil {
		return e
	}

	body, err := ioutil.ReadAll(resp.Body)
	beego.Info("body:" + string(body[0:len(body)]))

	var result ReturnSms
	err = xml.Unmarshal(body, &result)
	if err != nil || result.Code != "2" {
		return fmt.Errorf("send sms error!%s", result.Msg)
	}

	return nil
}

func GetRandFileName(path string, sl []string) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := fmt.Sprintf("%s/%d_%d.%s", path, time.Now().UnixNano(), rnd.Int()%10000, sl[len(sl)-1])
	return p

}

func GetOrderNo() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), rnd.Int()%1000000)
}

func IsFloatEqual(left float64, right float64) bool {
	if (left-right < 0.01 && left-right >= 0) || (right-left < 0.01 && right-left >= 0) {
		return true
	}
	return false
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func GetMailList() (mail string, err error) {
	var fields []string
	var sortby []string

	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	fields = append(fields, "Email")

	var m []interface{}

	m, err = GetAllEhomeAdmin(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}

	if len(m) <= 0 {
		err = fmt.Errorf("No admins")
	}
	for i := 0; i < len(m); i++ {
		if m[i].(map[string]interface{})["Email"].(string) == "" {
			continue
		}

		mail = mail + m[i].(map[string]interface{})["Email"].(string)
		if i != len(m)-1 {
			mail = mail + ";"
		}
	}
	beego.Error(mail)

	return
}

func SendMail(ordertime, orderno, appointtime, fixtype string, price float64, attact, name, phone, region, address string) error {
	user := "service@yijulezhu.com"
	passwd := "Jan591458"
	host := "smtp.mxhichina.com:25"
	to, err := GetMailList()
	if err != nil {
		return err
	}
	subject := fmt.Sprintf("新订单通知:%s", orderno)
	body := fmt.Sprintf("下单时间:%s<br>订单编号:%s<br>预约时间:%s<br>维修类型:%s<br>起步价格:%f<br>维修说明:%s<br>用户姓名:%s<br>手机号码:%s<br>所在地区:%s<br>联系地址:%s", ordertime, orderno, appointtime, fixtype, price, attact, name, phone, region, address)
	return SendToMail(user, passwd, host, to, subject, body)
}

func SendToMail(user, password, host, to, subject, body string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	content_type = "Content-Type: text/html" + "; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")

	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func GetSettlementrate(amount float64) (rate float64) {
	if amount <= 2000 {
		return 0.8
	} else if amount >= 5000 {
		return 0.9
	} else {
		return 0.85
	}
}

func GetMasterSettlement(amount float64) (set float64) {
	return amount * GetSettlementrate(amount)
}

func GetPlatformSettlement(amount float64) (set float64) {
	return amount * (1 - GetSettlementrate(amount))
}
