package image

import "github.com/zhaoyoucai/aeop/top"

type UploadTempImageAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *UploadTempImageAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.image.redefining.uploadtempimage")
	
	a.Ae = ae
	a.Api = api

}

func (a *UploadTempImageAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *UploadTempImageAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}