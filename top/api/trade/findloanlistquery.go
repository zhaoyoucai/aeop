package trade

import "github.com/zhaoyoucai/aeop/top"

type FindLoanListQueryAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *FindLoanListQueryAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.findloanlistquery")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindLoanListQueryAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindLoanListQueryAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}