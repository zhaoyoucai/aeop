package issue

import "github.com/zhaoyoucai/aeop/top"

type StateUpdateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *StateUpdateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.spain.issue.partner.rma.state.update")
	
	a.Ae = ae
	a.Api = api

}

func (a *StateUpdateAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *StateUpdateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}