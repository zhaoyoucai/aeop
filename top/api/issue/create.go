package issue

import "github.com/zhaoyoucai/aeop/top"

type ReverseLogisticCreateAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ReverseLogisticCreateAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.spain.issue.partner.rma.reverselogistic.trackinginfo.create")
	
	a.Ae = ae
	a.Api = api

}

func (a *ReverseLogisticCreateAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ReverseLogisticCreateAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}