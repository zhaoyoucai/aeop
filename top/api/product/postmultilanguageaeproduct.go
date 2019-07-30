package product

import "github.com/zhaoyoucai/aeop/top"

type PostMultiLanguageAeProductAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *PostMultiLanguageAeProductAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.postmultilanguageaeproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *PostMultiLanguageAeProductAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *PostMultiLanguageAeProductAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}