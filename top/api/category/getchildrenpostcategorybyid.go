package category

import "github.com/zhaoyoucai/aeop/top"

type GetChildrenPostCategoryByIdAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetChildrenPostCategoryByIdAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.category.redefining.getchildrenpostcategorybyid")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetChildrenPostCategoryByIdAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *GetChildrenPostCategoryByIdAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}