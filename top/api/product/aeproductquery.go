package product

import "github.com/zhaoyoucai/aeop/top"

type AeProductQueryAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *AeProductQueryAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.offer.product.query")
	
	a.Ae = ae
	a.Api = api

}

func (a *AeProductQueryAPI)SetParams(qs map[string]string) {
	//qs = {"product_id":1235464687}
	a.Api.SetParams(qs)
}

func (a *AeProductQueryAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}