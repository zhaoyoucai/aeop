package logistics

import "github.com/zhaoyoucai/aeop/top"

type GetPrintInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetPrintInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.getprintinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetPrintInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetPrintInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}