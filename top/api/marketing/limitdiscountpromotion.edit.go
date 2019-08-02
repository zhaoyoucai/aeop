package marketing

import "github.com/zhaoyoucai/aeop/top"

type LimitDiscountPromotionEditAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *LimitDiscountPromotionEditAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.limitdiscountpromotion.edit")
	
	a.Ae = ae
	a.Api = api

}

func (a *LimitDiscountPromotionEditAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *LimitDiscountPromotionEditAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}