// Code generated by Lark OpenAPI.

package larkoptical_char_recognition

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
)

type V1 struct {
	Image *image // 图片识别
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		Image: &image{config: config},
	}
}

type image struct {
	config *larkcore.Config
}

// BasicRecognize 基础图片识别 (OCR)
//
// - 可识别图片中的文字，按图片中的区域划分，分段返回文本列表
//
// - 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/optical_char_recognitionv1/basicRecognize_image.go
func (i *image) BasicRecognize(ctx context.Context, req *BasicRecognizeImageReq, options ...larkcore.RequestOptionFunc) (*BasicRecognizeImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/optical_char_recognition/v1/image/basic_recognize"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BasicRecognizeImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, i.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
