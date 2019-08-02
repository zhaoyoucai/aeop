package marketing

import "github.com/zhaoyoucai/aeop/top"

type PromotionListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *PromotionListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.promotion.list")
	
	a.Ae = ae
	a.Api = api

}

func (a *PromotionListAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *PromotionListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}