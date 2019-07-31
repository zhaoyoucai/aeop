package merchant

import "github.com/zhaoyoucai/aeop/top"

type QueryLevelInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryLevelInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.redefining.querylevelinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryLevelInfoAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryLevelInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}