package trade

import "github.com/zhaoyoucai/aeop/top"

type SaveRemarkAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SaveRemarkAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.redefining.saveremark")
	
	a.Ae = ae
	a.Api = api

}

func (a *SaveRemarkAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SaveRemarkAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}