package issue

import "github.com/zhaoyoucai/aeop/top"

type SaveSolutionAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SaveSolutionAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.issue.solution.save")
	
	a.Ae = ae
	a.Api = api

}

func (a *SaveSolutionAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SaveSolutionAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}