package marketing

import "github.com/zhaoyoucai/aeop/top"

type FindSellerCouponActivityAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindSellerCouponActivityAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.redefining.findsellercouponactivity")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindSellerCouponActivityAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *FindSellerCouponActivityAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}