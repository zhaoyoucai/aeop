package logistics

import "github.com/zhaoyoucai/aeop/top"

type CreateWarehouseOrderAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *CreateWarehouseOrderAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.createwarehouseorder")
	
	a.Ae = ae
	a.Api = api

}

func (a *CreateWarehouseOrderAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *CreateWarehouseOrderAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}