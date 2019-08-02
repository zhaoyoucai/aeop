package image

import "github.com/zhaoyoucai/aeop/top"

type ListGroupAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ListGroupAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.photobank.redefining.listgroup")
	
	a.Ae = ae
	a.Api = api

}

func (a *ListGroupAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ListGroupAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}