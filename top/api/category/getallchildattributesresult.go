package category

import "github.com/zhaoyoucai/aeop/top"

type GetAllChildAttributesResultAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetAllChildAttributesResultAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.category.redefining.getallchildattributesresult")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetAllChildAttributesResultAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetAllChildAttributesResultAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}