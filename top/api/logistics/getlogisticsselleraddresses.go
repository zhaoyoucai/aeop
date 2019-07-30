package logistics

import "github.com/zhaoyoucai/aeop/top"

type GetSsellerAddressesAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetSsellerAddressesAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.getlogisticsselleraddresses")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetSsellerAddressesAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetSsellerAddressesAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}