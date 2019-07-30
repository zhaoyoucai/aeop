package logistics

import "github.com/zhaoyoucai/aeop/top"

type GetOnlineLogisticsInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetOnlineLogisticsInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.getonlinelogisticsinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetOnlineLogisticsInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetOnlineLogisticsInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}