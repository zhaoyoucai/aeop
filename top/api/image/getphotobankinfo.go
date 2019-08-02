package image

import "github.com/zhaoyoucai/aeop/top"

type GetPhotoBankInfoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetPhotoBankInfoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.photobank.redefining.getphotobankinfo")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetPhotoBankInfoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetPhotoBankInfoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}