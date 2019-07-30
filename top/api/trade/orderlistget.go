package trade

import "github.com/zhaoyoucai/aeop/top"

type OrderListGetAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *OrderListGetAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.seller.orderlist.get")
	
	a.Ae = ae
	a.Api = api

}

func (a *OrderListGetAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *OrderListGetAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}