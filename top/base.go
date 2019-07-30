package top

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Aliexpress struct {
	AppKey    string
	AppSecret    string
	RedirectUri    string
	ServerUrl    string

	//app_key    string
	V    string
	PartnerId    string
	TargetAppKey    string
	Format    string
	Simplify    string
	Session    string
	SignMethod    string
	Timestamp    string

}


func (ae *Aliexpress) InitAPI(session string)  {

	cfg, err := goconfig.LoadConfigFile("conf/config.conf")
	if err != nil {
		log.Fatal(err)
	}

	AppKey, _ := cfg.GetValue("app", "app_key")
	AppSecret, _ := cfg.GetValue("app", "app_secret")
	RedirectUri, _ := cfg.GetValue("app", "redirect_uri")
	ServerUrl, _ := cfg.GetValue("app", "server_url")


	ae.AppKey = AppKey
	ae.AppSecret = AppSecret
	ae.RedirectUri = RedirectUri
	ae.ServerUrl = ServerUrl

	ae.Format = "json"
	ae.Simplify = "true"
	ae.V = "2.0"
	ae.PartnerId = ""
	ae.SignMethod = "md5"
	ae.TargetAppKey = ""

	ae.Session = session

	ae.Timestamp = time.Now().Format("2006-01-02 15:04:05")

}

func Signature(ae Aliexpress,api API) (string) {

	var queryMap map[string]string

	queryMap = make(map[string]string)

	//参与签名的系统级参数
	queryMap["app_key"] = ae.AppKey
	queryMap["method"] = api.ApiName
	queryMap["format"] = ae.Format
	queryMap["v"] = ae.V
	queryMap["partner_id"] = ae.PartnerId
	queryMap["session"] = ae.Session
	queryMap["sign_method"] = ae.SignMethod
	queryMap["simplify"] = ae.Simplify
	queryMap["target_app_key"] = ae.TargetAppKey

	queryMap["timestamp"] = ae.Timestamp

	//应用级参数与系统级参数合并
	for k,v := range api.Params {
		queryMap[k] = v
	}

	//参数列表排序
	var keys []string
	for k := range queryMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//拼接字符串
	b := bytes.Buffer{}

	//开头插入密钥
	b.WriteString(ae.AppSecret)

	/* 使用 key 输出 map 值 */
	//拼接参数值
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString(queryMap[k])

	}

	//结尾插入密钥
	b.WriteString(ae.AppSecret)

	//加密前的字符串
	str := b.String()

	//fmt.Println(str)

	//使用MD5加密

	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	//fmt.Println(md5str1)

	//全部转大写
	return strings.ToUpper(md5str1)

}

func BuildRequest(ae Aliexpress,api API) (string) {
	sign := Signature(ae,api)
	//fmt.Println(sign)

	//组装请求参数

	requestUrl := ae.ServerUrl
	u, _ := url.Parse(requestUrl)
	q := u.Query()
	q.Set("app_key", ae.AppKey)
	q.Set("session", ae.Session)
	q.Set("method", api.ApiName)
	q.Set("timestamp", ae.Timestamp)
	q.Set("format", ae.Format)
	q.Set("v", ae.V)
	q.Set("partner_id", ae.PartnerId)
	q.Set("target_app_key", ae.TargetAppKey)
	q.Set("simplify", ae.Simplify)
	q.Set("sign_method", ae.SignMethod)
	q.Set("sign", sign)

	//应用级参数与系统级参数合并
	for k,v := range api.Params {
		q.Set(k,v)
	}

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String());

	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s", result)

	return string(result)


}





