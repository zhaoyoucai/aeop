package marketing

import "github.com/zhaoyoucai/aeop/top"

type StorePromotionsListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *StorePromotionsListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.storepromotions.list")
	
	a.Ae = ae
	a.Api = api

}

func (a *StorePromotionsListAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *StorePromotionsListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}