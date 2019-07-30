package product

import "github.com/zhaoyoucai/aeop/top"

type EditSimpleProductFiledAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditSimpleProductFiledAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editsimpleproductfiled")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditSimpleProductFiledAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditSimpleProductFiledAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}