package marketing

import "github.com/zhaoyoucai/aeop/top"

type QueryProductPromotionlistAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryProductPromotionlistAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.redefining.querypromotionlistresultforproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryProductPromotionlistAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryProductPromotionlistAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}