package merchant

import "github.com/zhaoyoucai/aeop/top"

type GetProfileAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetProfileAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.profile.get")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetProfileAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *GetProfileAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}