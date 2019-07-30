package logistics

import "github.com/zhaoyoucai/aeop/top"

type GetPrintInfosAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetPrintInfosAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.logistics.redefining.getprintinfos")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetPrintInfosAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetPrintInfosAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}