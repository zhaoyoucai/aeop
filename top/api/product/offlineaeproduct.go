package product

import "github.com/zhaoyoucai/aeop/top"

type OfflineAeProductAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *OfflineAeProductAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.offlineaeproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *OfflineAeProductAPI)SetParams(qs map[string]string) {
	//qs = {"product_ids":"123134645;16546978;165467987"}
	a.Api.SetParams(qs)
}

func (a *OfflineAeProductAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}