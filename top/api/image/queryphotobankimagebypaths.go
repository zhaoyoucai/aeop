package image

import "github.com/zhaoyoucai/aeop/top"

type QueryPhotoBankImageByPathsAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryPhotoBankImageByPathsAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.photobank.redefining.queryphotobankimagebypaths")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryPhotoBankImageByPathsAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *QueryPhotoBankImageByPathsAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}