package appraise

import "github.com/zhaoyoucai/aeop/top"

type EvaluationReplyAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *EvaluationReplyAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.evaluation.evaluation.reply")
	
	a.Ae = ae
	a.Api = api

}

func (a *EvaluationReplyAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *EvaluationReplyAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}