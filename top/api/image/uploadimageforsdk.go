package image

import "github.com/zhaoyoucai/aeop/top"

type UploadImageForSdkAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *UploadImageForSdkAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.photobank.redefining.uploadimageforsdk")
	
	a.Ae = ae
	a.Api = api

}

func (a *UploadImageForSdkAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *UploadImageForSdkAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}