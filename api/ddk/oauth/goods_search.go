package oauth

import (
	"github.com/xiangwork/openpdd/api/ddk"
	"github.com/xiangwork/openpdd/core"
)

// GoodsSearchRequest 多多进宝商品查询 API Request
type GoodsSearchRequest struct {
	ddk.GoodsSearchRequest
	// ForceAuthDuoID 是否使用工具商专属推广计划，默认为false
	ForceAuthDuoID bool `json:"force_auth_duo_id,omitempty"`
}

// GetType implement Request interface
func (r GoodsSearchRequest) GetType() string {
	return "pdd.ddk.oauth.goods.search"
}

// GoodsSearch 多多进宝商品查询
func GoodsSearch(clt *core.SDKClient, req *GoodsSearchRequest, accessToken string) (*ddk.GoodsSearchResult, error) {
	var resp ddk.GoodsSearchResponse
	if err := clt.Do(req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Response, nil
}
