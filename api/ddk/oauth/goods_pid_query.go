package oauth

import (
	"github.com/xiangwork/openpdd/api/ddk"
	"github.com/xiangwork/openpdd/core"
)

// GoodsPidQueryRequest 查询已经生成的推广位信息 API Request
type GoodsPidQueryRequest struct {
	ddk.GoodsPidQueryRequest
}

// GetType implement Request interface
func (r GoodsPidQueryRequest) GetType() string {
	return "pdd.ddk.oauth.goods.pid.query"
}

// GoodsPidQuery 查询已经生成的推广位信息
func GoodsPidQuery(clt *core.SDKClient, req *GoodsPidQueryRequest, accessToken string) (int, []ddk.Pid, error) {
	var resp ddk.GoodsPidQueryResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return 0, nil, err
	}
	return resp.Response.TotalCount, resp.Response.List, nil
}
