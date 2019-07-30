package logistics

import "github.com/zhaoyoucai/aeop/top"

type QueryLogisticsOrderDetailAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryLogisticsOrderDetailAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.querylogisticsorderdetail")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryLogisticsOrderDetailAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *QueryLogisticsOrderDetailAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}