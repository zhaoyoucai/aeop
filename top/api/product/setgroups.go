package product

import "github.com/zhaoyoucai/aeop/top"

type AeProductSetGroupsAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *AeProductSetGroupsAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.setgroups")
	
	a.Ae = ae
	a.Api = api

}

func (a *AeProductSetGroupsAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *AeProductSetGroupsAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}