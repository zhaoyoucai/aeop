package product

import "github.com/zhaoyoucai/aeop/top"

type EditSingleSkuPriceAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditSingleSkuPriceAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editsingleskuprice")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditSingleSkuPriceAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditSingleSkuPriceAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}