package image

import "github.com/zhaoyoucai/aeop/top"

type DelUnusePhotoAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *DelUnusePhotoAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.photobank.redefining.delunusephoto")
	
	a.Ae = ae
	a.Api = api

}

func (a *DelUnusePhotoAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *DelUnusePhotoAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}