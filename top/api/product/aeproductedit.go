package product

import "github.com/zhaoyoucai/aeop/top"

type AeProductEditAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *AeProductEditAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.offer.product.edit")
	
	a.Ae = ae
	a.Api = api

}

func (a *AeProductEditAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *AeProductEditAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}