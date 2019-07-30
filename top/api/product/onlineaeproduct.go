package product

import "github.com/zhaoyoucai/aeop/top"

type OnlineAeProductAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *OnlineAeProductAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.onlineaeproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *OnlineAeProductAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *OnlineAeProductAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}