package logistics

import "github.com/zhaoyoucai/aeop/top"

type GetOnlineLogisticsServiceListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetOnlineLogisticsServiceListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.getonlinelogisticsservicelistbyorderid")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetOnlineLogisticsServiceListAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetOnlineLogisticsServiceListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}