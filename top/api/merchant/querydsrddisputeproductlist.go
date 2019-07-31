package merchant

import "github.com/zhaoyoucai/aeop/top"

type QueryDSRddisputeProductListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryDSRddisputeProductListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.merchant.redefining.querydsrddisputeproductlist")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryDSRddisputeProductListAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryDSRddisputeProductListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}