package trade

import "github.com/zhaoyoucai/aeop/top"

type FindOrderReceiptInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindOrderReceiptInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.findorderreceiptinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindOrderReceiptInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindOrderReceiptInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}