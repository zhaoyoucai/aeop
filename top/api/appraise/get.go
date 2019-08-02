package appraise

import "github.com/zhaoyoucai/aeop/top"

type ListOrderEvaluationGetAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *ListOrderEvaluationGetAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.evaluation.listorderevaluation.get")
	
	a.Ae = ae
	a.Api = api

}

func (a *ListOrderEvaluationGetAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *ListOrderEvaluationGetAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}