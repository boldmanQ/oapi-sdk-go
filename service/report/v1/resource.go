// Code generated by Lark OpenAPI.

package larkreport

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
)

type V1 struct {
	Rule     *rule     // 规则
	RuleView *ruleView // 规则看板
	Task     *task     // 任务
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		Rule:     &rule{config: config},
		RuleView: &ruleView{config: config},
		Task:     &task{config: config},
	}
}

type rule struct {
	config *larkcore.Config
}
type ruleView struct {
	config *larkcore.Config
}
type task struct {
	config *larkcore.Config
}

// Query 规则查询
//
// - 规则查询
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/report/report-v1/rule/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/reportv1/query_rule.go
func (r *rule) Query(ctx context.Context, req *QueryRuleReq, options ...larkcore.RequestOptionFunc) (*QueryRuleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/report/v1/rules/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, r.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryRuleResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, r.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Remove 移除规则看板
//
// - 移除规则看板
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/report/report-v1/rule-view/remove
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/reportv1/remove_ruleView.go
func (r *ruleView) Remove(ctx context.Context, req *RemoveRuleViewReq, options ...larkcore.RequestOptionFunc) (*RemoveRuleViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/report/v1/rules/:rule_id/views/remove"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, r.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RemoveRuleViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, r.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Query 任务查询
//
// - 任务查询
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/report/report-v1/task/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/reportv1/query_task.go
func (t *task) Query(ctx context.Context, req *QueryTaskReq, options ...larkcore.RequestOptionFunc) (*QueryTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/report/v1/tasks/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
