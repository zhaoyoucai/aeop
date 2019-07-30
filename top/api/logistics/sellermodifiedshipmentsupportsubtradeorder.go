package logistics

import "github.com/zhaoyoucai/aeop/top"

type SellerModifiedShipmentSubTradeOrderAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SellerModifiedShipmentSubTradeOrderAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.sellermodifiedshipmentsupportsubtradeorder")
	
	a.Ae = ae
	a.Api = api

}

func (a *SellerModifiedShipmentSubTradeOrderAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SellerModifiedShipmentSubTradeOrderAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}