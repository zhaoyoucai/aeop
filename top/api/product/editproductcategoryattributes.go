package product

import "github.com/zhaoyoucai/aeop/top"

type EditAeProductCategoryAttributesAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EditAeProductCategoryAttributesAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.editproductcategoryattributes")
	
	a.Ae = ae
	a.Api = api

}

func (a *EditAeProductCategoryAttributesAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *EditAeProductCategoryAttributesAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}