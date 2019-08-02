package issue

import "github.com/zhaoyoucai/aeop/top"

type GetIssueListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetIssueListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.issue.issuelist.get")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetIssueListAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetIssueListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}