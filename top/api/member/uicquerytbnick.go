package member

import "github.com/zhaoyoucai/aeop/top"

type UicQueryTbNnickAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *UicQueryTbNnickAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.member.redefining.uicquerytbnick")
	
	a.Ae = ae
	a.Api = api

}

func (a *UicQueryTbNnickAPI)SetParams(qs map[string]string) {
	//qs = {"aeop_a_e_product":{}}
	a.Api.SetParams(qs)
}

func (a *UicQueryTbNnickAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}