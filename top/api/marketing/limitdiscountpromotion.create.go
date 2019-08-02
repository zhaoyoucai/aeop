package marketing

import "github.com/zhaoyoucai/aeop/top"

type LimitDiscountPromotionCreateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *LimitDiscountPromotionCreateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.limitdiscountpromotion.create")
	
	a.Ae = ae
	a.Api = api

}

func (a *LimitDiscountPromotionCreateAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *LimitDiscountPromotionCreateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}