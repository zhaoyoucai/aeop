package issue

import "github.com/zhaoyoucai/aeop/top"

type AgreeSolutionAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *AgreeSolutionAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.issue.solution.agree")
	
	a.Ae = ae
	a.Api = api

}

func (a *AgreeSolutionAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *AgreeSolutionAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}