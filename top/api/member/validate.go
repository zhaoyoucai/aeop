package member

import "github.com/zhaoyoucai/aeop/top"

type ValidateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ValidateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.member.trust.token.validate")
	
	a.Ae = ae
	a.Api = api

}

func (a *ValidateAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *ValidateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}