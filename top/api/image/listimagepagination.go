package image

import "github.com/zhaoyoucai/aeop/top"

type ListImagePaginationAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ListImagePaginationAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.photobank.redefining.listimagepagination")
	
	a.Ae = ae
	a.Api = api

}

func (a *ListImagePaginationAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ListImagePaginationAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}