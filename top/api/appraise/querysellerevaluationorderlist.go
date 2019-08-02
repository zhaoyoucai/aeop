package appraise

import "github.com/zhaoyoucai/aeop/top"

type QuerySellerEvaluationOrderListAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *QuerySellerEvaluationOrderListAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.appraise.redefining.querysellerevaluationorderlist")
	
	a.Ae = ae
	a.Api = api

}

func (a *QuerySellerEvaluationOrderListAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *QuerySellerEvaluationOrderListAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}