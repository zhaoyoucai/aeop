package marketing

import "github.com/zhaoyoucai/aeop/top"

type GetActListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *GetActListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.redefining.getactlist")
	
	a.Ae = ae
	a.Api = api

}

func (a *GetActListAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *GetActListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}