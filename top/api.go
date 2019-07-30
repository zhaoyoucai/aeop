package top

type API struct {
	ApiName string
	Params map[string]string
}

func (api *API)SetApiName(apiname string)  {
	api.ApiName = apiname
}

func (api *API)SetParams(parms map[string]string)  {
	api.Params = parms
}

func (api *API)GetApiName() string  {
	return api.ApiName
}
func (api *API)GetParams() map[string]string  {
	return api.Params
}

