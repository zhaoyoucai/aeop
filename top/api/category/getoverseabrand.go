package category

import "github.com/zhaoyoucai/aeop/top"

type GetOverseaBrandAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetOverseaBrandAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.oversea.brand.get")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetOverseaBrandAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetOverseaBrandAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}