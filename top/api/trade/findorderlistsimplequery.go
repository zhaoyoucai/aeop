package trade

import "github.com/zhaoyoucai/aeop/top"

type FindOrderListSimpleQueryAPI struct {
	Ae top.Aliexpress
	Api top.API
}

type FindOrderListSimpleQueryResponse struct {
	Result struct {
		OrderList []struct {
			GmtModified string `json:"gmt_modified"`
			OrderID int64 `json:"order_id"`
			Memo string `json:"memo,omitempty"`
			GmtCreate string `json:"gmt_create"`
			BizType string `json:"biz_type"`
			ProductList []struct {
				GoodsPrepareTime int `json:"goods_prepare_time"`
				ProductImgURL string `json:"product_img_url"`
				ProductID int64 `json:"product_id"`
				ProductSnapURL string `json:"product_snap_url"`
				ProductUnitPrice string `json:"product_unit_price"`
				MoneyBack3X bool `json:"money_back3x"`
				ProductUnit string `json:"product_unit"`
				ProductCount int `json:"product_count"`
				ChildID int64 `json:"child_id"`
				SonOrderStatus string `json:"son_order_status"`
				ProductName string `json:"product_name"`
				ProductUnitPriceCur string `json:"product_unit_price_cur"`
			} `json:"product_list"`
			OrderStatus string `json:"order_status"`
		} `json:"order_list"`
		TotalItem int `json:"total_item"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

func  (a *FindOrderListSimpleQueryAPI)Init(session string)  {

	ae := top.Aliexpress{}
	ae.InitAPI(session)

	api := top.API{}

	api.SetApiName("aliexpress.trade.redefining.findorderlistsimplequery")
	
	a.Ae = ae
	a.Api = api

}

func (a *FindOrderListSimpleQueryAPI)SetParams(qs map[string]string) {
	//qs = {"param1":{}}
	a.Api.SetParams(qs)
}

func (a *FindOrderListSimpleQueryAPI)GetResponse()  (string){
	return top.BuildRequest(a.Ae,a.Api)
}