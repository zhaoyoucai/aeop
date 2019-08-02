package marketing

import "github.com/zhaoyoucai/aeop/top"

type QueryByProductAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QueryByProductAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.marketing.storepromotions.querybyproduct")
	
	a.Ae = ae
	a.Api = api

}

func (a *QueryByProductAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *QueryByProductAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}