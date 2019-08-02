package marketing

import "github.com/zhaoyoucai/aeop/top"

type LimitDiscountPromotionProductEditAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *LimitDiscountPromotionProductEditAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.limitdiscountpromotionproduct.edit")
	
	a.Ae = ae
	a.Api = api

}

func (a *LimitDiscountPromotionProductEditAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *LimitDiscountPromotionProductEditAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}