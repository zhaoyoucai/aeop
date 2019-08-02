package marketing

import "github.com/zhaoyoucai/aeop/top"

type LimitDiscountPromotionProductDelAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *LimitDiscountPromotionProductDelAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.limitdiscountpromotionproduct.del")
	
	a.Ae = ae
	a.Api = api

}

func (a *LimitDiscountPromotionProductDelAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *LimitDiscountPromotionProductDelAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}