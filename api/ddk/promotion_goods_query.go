package ddk

import (
	"github.com/xiangwork/openpdd/core"
	"github.com/xiangwork/openpdd/model"
)

// PromotionGoodsQueryRequest 多多进宝信息流渠道备案查询接口 API Request
type PromotionGoodsQueryRequest struct {
	// DemoURL 推广商品视频素材url
	GoodsID    int64 `json:"goods_id,omitempty"`
	MallId     int64 `json:"mall_id,omitempty"`
	PageNumber int   `json:"page_number,omitempty"`
	PageSize   int   `json:"page_size,omitempty"`
}

// GetType implement Request interface
func (r PromotionGoodsQueryRequest) GetType() string {
	return "pdd.ddk.promotion.goods.query"
}

// PromotionGoodsQueryResponse 多多进宝信息流渠道备案授查询接口 API Response
type PromotionGoodsQueryResponse struct {
	model.CommonResponse
	Response struct {
		ApplicationList []ApplicationList `json:"application_list,omitempty"`
		Total           int               `json:"total,omitempty"`
	} `json:"response"`
}
type ApplicationList struct {
	Comment string `json:"comment,omitempty"`
	GoodsID int64  `json:"goods_id,omitempty"`
	MallId  int64  `json:"mall_id,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// GoodsPromotionRightAuth 多多进宝信息流渠道备案授权素材上传接口
func PromotionGoodsQuery(clt *core.SDKClient, req *PromotionGoodsQueryRequest) ([]ApplicationList, error) {
	var resp PromotionGoodsQueryResponse
	if err := clt.Do(req, &resp, ""); err != nil {
		return nil, err
	}

	return resp.Response.ApplicationList, nil
}
