package trade

import "github.com/zhaoyoucai/aeop/top"

type QueryRemarksAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryRemarksAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.redefining.queryremarks")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryRemarksAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *QueryRemarksAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}