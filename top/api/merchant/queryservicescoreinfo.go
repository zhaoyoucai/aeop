package merchant

import "github.com/zhaoyoucai/aeop/top"

type QueryServicesCoreInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryServicesCoreInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.redefining.queryservicescoreinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryServicesCoreInfoAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryServicesCoreInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}