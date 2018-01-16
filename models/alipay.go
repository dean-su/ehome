package models

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"sort"
	"time"
)

var (
	PRIVATEKEY = `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDbNkIhr0wbK/dmqiPvc/wNcNZ5kxzwZRGo5c5V+zV8/ikU7cy/
1tqzzlYtECniKJWBWlF4ySCRUL3J6AZEc9EbxI6Nd5q+zZOfayO7rBh+MohZSQuM
CwRV2dzww370oSCllz8o0j7boixEPIEhFrYERrBl5kFN8vmDiPqxytKKgwIDAQAB
AoGAKzt3LWeCDfXM1A2ICsiIFCqF0fJGic6U/gdcey6Q7Pa/KWaAs/Durvlpm3eN
mxG/8oPaf4TDDIWs+G3vFn705TRsY/gsazM4/o7X8c4fyh94jyN/9EIu3iB/o9qJ
5ZtcFiOYtjjmjSJ1kCyJDvoXoqcfAhiW1+8yuF/2TgbSonkCQQD8exTsNl8Vuz5W
1T4cM3DPrqSMDnkWHBR2gKJ24cBPq1Wraat01sQWLuz+eZEFUhlswzU9uKQZqiJo
plOoD5cHAkEA3kR27dK1fNt+QfNoMH6WhY7OJh667vuKd5Zzc+9hBD0NOA13eF4r
DlzLwy8+uCZIzHb66UhheY1Xw5n88UB1pQJAMgFen3oVUwSG0EPjzUoS6c85H9Qt
/0cIdn/7rXgA0saobQ68uqNeqtYYcj45jsg36OawwMy1E7AyqG9o2jxcOQJAc9bi
1Nf4VnZeuyiMqJmRAVKIBj4F1v+qDuGOkmG0Am1/MjNyjH0nm3ipibRZz8fEMSvh
YSukAaG0l/Dtbx7VMQJBAMZ+CHVsQnNFBOm8ZZbj01bkJVGc/mbQaEz9MYa1BdDi
F8jmL1eyA+Svx1U4qLIMt+bUNFT8U9B/1v9ghQGwlkc=
-----END RSA PRIVATE KEY-----
`
)

func aliPaySign(mReq map[string]interface{}) string {

	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for i, k := range sorted_keys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + "=" + value + "&"
			} else {
				signStrings = signStrings + k + "=" + value //最后一个不加此符号
			}
		}
	}
	fmt.Println("生成的待签名---->", signStrings)

	//============================= 开始签名 ==================================
	block, _ := pem.Decode([]byte(PRIVATEKEY))
	if block == nil {
		fmt.Println("rsaSign private_key error")
		return ""
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey-------privateKey----- error : %v\n", err)
		return ""
	} else {
		//fmt.Println("x509.ParsePKCS1PrivateKey-------privateKey-----", privateKey)
	}

	result, err := RsaSign(signStrings, privateKey)
	fmt.Println("alipay.RsaSign=========", result, err)
	return result
}

/**
 * RSA签名
 * @param $data 待签名数据
 * @param $private_key_path 商户私钥文件路径
 * return 签名结果
 */
func RsaSign(origData string, privateKey *rsa.PrivateKey) (string, error) {

	h := sha1.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}

func GetOrderInfo() string {
	var m map[string]interface{}
	m = make(map[string]interface{})

	m["app_id"] = "2016101000651592"
	m["method"] = "alipay.trade.app.pay"
	m["format"] = "JSON"
	m["charset"] = "utf-8"
	m["sign_type"] = "RSA"
	m["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	m["version"] = "1.0"
	m["notify_url"] = "http://120.25.74.193/v1/alipay/notify"
	m["biz_content"] = fmt.Sprintf("body:test,subject:testsub,out_trade_no:%s,total_amount:0.01,product_code:QUICK_MSECURITY_PAY", "123")

	m["sign"] = aliPaySign(m)
	return ""
}

/*
参数 类型 是否必填 最大长度 描述 示例值
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
