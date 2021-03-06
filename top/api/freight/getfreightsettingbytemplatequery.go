package freight

import "github.com/zhaoyoucai/aeop/top"

type GetFreightSettingAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetFreightSettingAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.freight.redefining.getfreightsettingbytemplatequery")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetFreightSettingAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetFreightSettingAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}