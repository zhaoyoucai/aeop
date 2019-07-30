package trade

import "github.com/zhaoyoucai/aeop/top"

type RefuseCancelAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *RefuseCancelAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.seller.order.refusecancel")
	
	a.Ae = ae
	a.Api = api

}

func (a *RefuseCancelAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *RefuseCancelAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}