package freight

import "github.com/zhaoyoucai/aeop/top"

type ListFreightTemplateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ListFreightTemplateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.seller.order.acceptcancel")
	
	a.Ae = ae
	a.Api = api

}

func (a *ListFreightTemplateAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ListFreightTemplateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}