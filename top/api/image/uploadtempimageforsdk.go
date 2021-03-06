package image

import "github.com/zhaoyoucai/aeop/top"

type UploadTempImageForSdkAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *UploadTempImageForSdkAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.image.redefining.uploadtempimageforsdk")
	
	a.Ae = ae
	a.Api = api

}

func (a *UploadTempImageForSdkAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *UploadTempImageForSdkAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}