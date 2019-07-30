package logistics

import "github.com/zhaoyoucai/aeop/top"

type SellerModifiedShipmentForTopAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SellerModifiedShipmentForTopAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.sellermodifiedshipmentfortop")
	
	a.Ae = ae
	a.Api = api

}

func (a *SellerModifiedShipmentForTopAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SellerModifiedShipmentForTopAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}