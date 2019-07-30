package trade

import "github.com/zhaoyoucai/aeop/top"

type FindOrderByIdAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindOrderByIdAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.new.redefining.findorderbyid")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindOrderByIdAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindOrderByIdAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}