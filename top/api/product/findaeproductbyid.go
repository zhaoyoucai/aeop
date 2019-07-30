package product

import "github.com/zhaoyoucai/aeop/top"

type FindAeProductByIdAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindAeProductByIdAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.findaeproductbyid")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindAeProductByIdAPI)SetParams(qs map[string]string) {
	//qs = {"product_id":1235464687}
	a.Api.SetParams(qs)
}

func (a *FindAeProductByIdAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}