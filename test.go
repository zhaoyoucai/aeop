package main

import (
	"fmt"
	"github.com/zhaoyoucai/aeop/top/api/product"  //引入API所在的包
)



func main() {

	//设置session
	var session = "5459519804624b1rpabpSdtFhdsQuZC7I0EMR170c63d5AVodZehwauVdNne1254m5Gi"

	//创建API
	aeApi := product.OnlineAeProductAPI{}
	//初始化
	aeApi.Init(session)
	//构造请求参数
	p := make(map[string]string)
	p["product_id"] = "1232546546"
	aeApi.SetParams(p)

	//打印API名称
	//fmt.Println(aeApi.Api.GetApiName())

	//执行请求
	result := aeApi.GetResponse()

	fmt.Println(result)
}
