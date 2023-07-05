package oauth

import (
	"github.com/xiangwork/openpdd/api/ddk"
	"github.com/xiangwork/openpdd/core"
)

// AllOrderListIncrementGetRequest 查询所有授权的多多客订单API Request
type AllOrderListIncrementGetRequest struct {
	ddk.OrderListIncrementGetRequest
}

// GetType implement Request interface
func (r AllOrderListIncrementGetRequest) GetType() string {
	return "pdd.ddk.all.order.list.increment.get"
}

// AllOrderListIncrementGet 查询所有授权的多多客订单
func AllOrderListIncrementGet(clt *core.SDKClient, req *AllOrderListIncrementGetRequest, accessToken string) (int, []ddk.Order, error) {
	var resp ddk.OrderListIncrementGetResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
