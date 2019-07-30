package trade

import "github.com/zhaoyoucai/aeop/top"

type FindOrderBaseInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindOrderBaseInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.findorderbaseinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindOrderBaseInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindOrderBaseInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}