package product

import "github.com/zhaoyoucai/aeop/top"

type EditMutilpleSkuStocksAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditMutilpleSkuStocksAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editmutilpleskustocks")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditMutilpleSkuStocksAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditMutilpleSkuStocksAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}