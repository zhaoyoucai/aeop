package product

import "github.com/zhaoyoucai/aeop/top"

type EditAeProductCidAttidSkuAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditAeProductCidAttidSkuAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editproductcidattidsku")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditAeProductCidAttidSkuAPI)SetParams(qs map[string]string) {
	//qs = {"product_id":1235464687,"category_id":"","product_skus":"","product_properties":""}
	a.Api.SetParams(qs)
}

func (a *EditAeProductCidAttidSkuAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}