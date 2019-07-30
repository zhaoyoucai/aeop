package main

import (
	"fmt"
	"github.com/zhaoyoucai/aeop/top/api/product"  //引入API所在的包
)



func main() {

	//设置session
	var session = "50002900324b1rpabpSdtFhFPVQuZC7I0EMR170c63d5AVodZeDxgnuVdNne984m5Gi"

	//创建API
	aeApi := product.AeProductQueryAPI{}
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
