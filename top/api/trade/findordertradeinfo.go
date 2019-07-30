package trade

import "github.com/zhaoyoucai/aeop/top"

type FindOrderTradeInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindOrderTradeInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.findordertradeinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindOrderTradeInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindOrderTradeInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}