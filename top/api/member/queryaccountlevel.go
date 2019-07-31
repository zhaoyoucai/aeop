package member

import "github.com/zhaoyoucai/aeop/top"

type QueryAccountLevelAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryAccountLevelAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.member.redefining.queryaccountlevel")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryAccountLevelAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryAccountLevelAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}