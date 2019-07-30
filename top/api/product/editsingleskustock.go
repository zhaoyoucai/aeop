package product

import "github.com/zhaoyoucai/aeop/top"

type EditSingleSkuStockAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditSingleSkuStockAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editsingleskustock")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditSingleSkuStockAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditSingleSkuStockAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}