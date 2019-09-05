package product

import "github.com/zhaoyoucai/aeop/top"

type GetProductGroupsAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetProductGroupsAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.product.productgroups.get")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetProductGroupsAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *GetProductGroupsAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}