package marketing

import "github.com/zhaoyoucai/aeop/top"

type ProductsQueryAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ProductsQueryAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.storepromotion.products.query")
	
	a.Ae = ae
	a.Api = api

}

func (a *ProductsQueryAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *ProductsQueryAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}