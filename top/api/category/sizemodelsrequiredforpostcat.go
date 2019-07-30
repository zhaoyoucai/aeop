package category

import "github.com/zhaoyoucai/aeop/top"

type SizeModelsRequiredAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SizeModelsRequiredAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.category.redefining.sizemodelsrequiredforpostcat")
	
	a.Ae = ae
	a.Api = api

}

func (a *SizeModelsRequiredAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SizeModelsRequiredAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}