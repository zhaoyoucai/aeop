package category

import "github.com/zhaoyoucai/aeop/top"

type GetPostCategoryByIdAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetPostCategoryByIdAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.category.redefining.getpostcategorybyid")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetPostCategoryByIdAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetPostCategoryByIdAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}