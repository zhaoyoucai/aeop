package category

import "github.com/zhaoyoucai/aeop/top"

type GetChildAttributesAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetChildAttributesAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.category.redefining.getchildattributesresultbypostcateidandpath")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetChildAttributesAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetChildAttributesAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}