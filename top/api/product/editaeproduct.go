package product

import "github.com/zhaoyoucai/aeop/top"

type EditAeProductAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditAeProductAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editaeproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditAeProductAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditAeProductAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}