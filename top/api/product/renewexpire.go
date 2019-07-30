package product

import "github.com/zhaoyoucai/aeop/top"

type RenewExpireAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *RenewExpireAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.postproduct.redefining.renewexpire")
	
	a.Ae = ae
	a.Api = api

}

func (a *RenewExpireAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *RenewExpireAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}