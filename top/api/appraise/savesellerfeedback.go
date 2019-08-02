package appraise

import "github.com/zhaoyoucai/aeop/top"

type SaveSellerFeedbackAPI struct {
	Ae top.Aliexpress
	Api top.API
}

func  (a *SaveSellerFeedbackAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.appraise.redefining.savesellerfeedback")
	
	a.Ae = ae
	a.Api = api

}

func (a *SaveSellerFeedbackAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *SaveSellerFeedbackAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}