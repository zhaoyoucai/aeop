package trade

import "github.com/zhaoyoucai/aeop/top"

type FindOrderListSimpleQueryAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindOrderListSimpleQueryAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.findorderlistsimplequery")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindOrderListSimpleQueryAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindOrderListSimpleQueryAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}