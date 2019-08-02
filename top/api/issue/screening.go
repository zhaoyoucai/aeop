package issue

import "github.com/zhaoyoucai/aeop/top"

type ScreeningCreateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ScreeningCreateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.spain.issue.partner.rma.screening.create")
	
	a.Ae = ae
	a.Api = api

}

func (a *ScreeningCreateAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ScreeningCreateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}