package marketing

import "github.com/zhaoyoucai/aeop/top"

type AddPromotionProductAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *AddPromotionProductAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.limiteddiscountpromotion.addpromotionproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *AddPromotionProductAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *AddPromotionProductAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}