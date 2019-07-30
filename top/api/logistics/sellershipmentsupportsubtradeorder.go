package logistics

import "github.com/zhaoyoucai/aeop/top"

type SellerShipmentSubTradeOrderAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SellerShipmentSubTradeOrderAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.sellershipmentsupportsubtradeorder")
	
	a.Ae = ae
	a.Api = api

}

func (a *SellerShipmentSubTradeOrderAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SellerShipmentSubTradeOrderAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}