package product

import "github.com/zhaoyoucai/aeop/top"

type EditAeProductMultiLanguageAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditAeProductMultiLanguageAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editmultilanguageproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditAeProductMultiLanguageAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditAeProductMultiLanguageAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}