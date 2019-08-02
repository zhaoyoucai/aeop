package issue

import "github.com/zhaoyoucai/aeop/top"

type UploadImageAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *UploadImageAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.issue.image.upload")
	
	a.Ae = ae
	a.Api = api

}

func (a *UploadImageAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *UploadImageAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}