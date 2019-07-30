package logistics

import "github.com/zhaoyoucai/aeop/top"

type ListLogisticSserviceAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ListLogisticSserviceAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.listlogisticsservice")
	
	a.Ae = ae
	a.Api = api

}

func (a *ListLogisticSserviceAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ListLogisticSserviceAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}