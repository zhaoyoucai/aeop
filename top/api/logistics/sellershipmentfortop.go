package logistics

import "github.com/zhaoyoucai/aeop/top"

type SellerShipmentForTopAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SellerShipmentForTopAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.sellershipmentfortop")
	
	a.Ae = ae
	a.Api = api

}

func (a *SellerShipmentForTopAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SellerShipmentForTopAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}