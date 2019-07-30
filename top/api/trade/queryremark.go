package trade

import "github.com/zhaoyoucai/aeop/top"

type QueryRemarkAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryRemarkAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.redefining.queryremark")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryRemarkAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *QueryRemarkAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}