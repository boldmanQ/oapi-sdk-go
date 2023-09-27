// Package application code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkapplication

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *ApplicationService {
	a := &ApplicationService{config: config}
	a.AppRecommendRule = &appRecommendRule{service: a}
	a.Application = &application{service: a}
	a.ApplicationAppUsage = &applicationAppUsage{service: a}
	a.ApplicationAppVersion = &applicationAppVersion{service: a}
	a.ApplicationContactsRange = &applicationContactsRange{service: a}
	a.ApplicationFeedback = &applicationFeedback{service: a}
	a.ApplicationVisibility = &applicationVisibility{service: a}
	a.Bot = &bot{service: a}
	return a
}

type ApplicationService struct {
	config                   *larkcore.Config
	AppRecommendRule         *appRecommendRule         // 我的常用推荐规则
	Application              *application              // 应用
	ApplicationAppUsage      *applicationAppUsage      // 应用使用情况
	ApplicationAppVersion    *applicationAppVersion    // 事件
	ApplicationContactsRange *applicationContactsRange // application.contacts_range
	ApplicationFeedback      *applicationFeedback      // 应用反馈
	ApplicationVisibility    *applicationVisibility    // 事件
	Bot                      *bot                      // 事件
}

type appRecommendRule struct {
	service *ApplicationService
}
type application struct {
	service *ApplicationService
}
type applicationAppUsage struct {
	service *ApplicationService
}
type applicationAppVersion struct {
	service *ApplicationService
}
type applicationContactsRange struct {
	service *ApplicationService
}
type applicationFeedback struct {
	service *ApplicationService
}
type applicationVisibility struct {
	service *ApplicationService
}
type bot struct {
	service *ApplicationService
}

// 获取当前设置的推荐规则列表
//
// - 获取当前设置的推荐规则列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/app_recommend_rule/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/list_appRecommendRule.go
func (a *appRecommendRule) List(ctx context.Context, req *ListAppRecommendRuleReq, options ...larkcore.RequestOptionFunc) (*ListAppRecommendRuleResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/app_recommend_rules"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListAppRecommendRuleResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *appRecommendRule) ListByIterator(ctx context.Context, req *ListAppRecommendRuleReq, options ...larkcore.RequestOptionFunc) (*ListAppRecommendRuleIterator, error) {
	return &ListAppRecommendRuleIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.List,
		options:  options,
		limit:    req.Limit}, nil
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=contacts_range_configuration&project=application&resource=application&version=v6
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/contactsRangeConfiguration_application.go
func (a *application) ContactsRangeConfiguration(ctx context.Context, req *ContactsRangeConfigurationApplicationReq, options ...larkcore.RequestOptionFunc) (*ContactsRangeConfigurationApplicationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/contacts_range_configuration"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ContactsRangeConfigurationApplicationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取应用信息
//
// - 根据app_id获取应用的基础信息
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/get_application.go
func (a *application) Get(ctx context.Context, req *GetApplicationReq, options ...larkcore.RequestOptionFunc) (*GetApplicationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetApplicationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新应用分组信息
//
// - 更新应用的分组信息（分组会影响应用在工作台中的分类情况，请谨慎更新）
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/patch_application.go
func (a *application) Patch(ctx context.Context, req *PatchApplicationReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查看待审核的应用列表
//
// - 查看本企业下所有待审核的自建应用列表
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/underauditlist
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/underauditlist_application.go
func (a *application) Underauditlist(ctx context.Context, req *UnderauditlistApplicationReq, options ...larkcore.RequestOptionFunc) (*UnderauditlistApplicationResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/underauditlist"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnderauditlistApplicationResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *application) UnderauditlistByIterator(ctx context.Context, req *UnderauditlistApplicationReq, options ...larkcore.RequestOptionFunc) (*UnderauditlistApplicationIterator, error) {
	return &UnderauditlistApplicationIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.Underauditlist,
		options:  options,
		limit:    req.Limit}, nil
}

// 获取多部门应用使用概览（灰度租户可见）
//
// - 查看应用在某一天/某一周/某一个月的使用数据，可以根据部门做多层子部门的筛选
//
// - 1. 仅支持企业版/旗舰版租户使用;2. 一般每天早上10点产出前一天的数据;3. 已经支持的指标包括：应用的活跃用户数、累计用户数、新增用户数、访问页面数、打开次数;4. 按照部门查看数据时，可以分别展示当前部门以及其子部门的使用情况;5. 如果查询的部门在查询日期没有使用过应用，只返回指标：应用的活跃用户数指标;6. 数据从飞书4.10版本开始统计，使用飞书版本4.10及以下版本的用户数据不会被统计到;7. 调用频控为100次/分
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_usage/department_overview
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/departmentOverview_applicationAppUsage.go
func (a *applicationAppUsage) DepartmentOverview(ctx context.Context, req *DepartmentOverviewApplicationAppUsageReq, options ...larkcore.RequestOptionFunc) (*DepartmentOverviewApplicationAppUsageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/app_usage/department_overview"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DepartmentOverviewApplicationAppUsageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取应用使用概览
//
// - 查看应用在某一天/某一周/某一个月的使用数据，可以查看租户整体对应用的使用情况，也可以分部门查看。
//
// - 1. 仅支持企业版/旗舰版租户使用;2. 一般每天早上10点产出前一天的数据;3. 已经支持的指标包括：应用的活跃用户数、累计用户数、新增用户数、访问页面数、打开次数;4. 数据从飞书4.10版本开始统计，使用飞书版本4.10及以下版本的用户数据不会被统计到;5. 按照部门查看数据时，会展示当前部门以及其子部门的整体使用情况;6. 调用频控为100次/分
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_usage/overview
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/overview_applicationAppUsage.go
func (a *applicationAppUsage) Overview(ctx context.Context, req *OverviewApplicationAppUsageReq, options ...larkcore.RequestOptionFunc) (*OverviewApplicationAppUsageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/app_usage/overview"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &OverviewApplicationAppUsageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//
//
// - 获取应用版本通讯录权限范围建议
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=contacts_range_suggest&project=application&resource=application.app_version&version=v6
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/contactsRangeSuggest_applicationAppVersion.go
func (a *applicationAppVersion) ContactsRangeSuggest(ctx context.Context, req *ContactsRangeSuggestApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*ContactsRangeSuggestApplicationAppVersionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/app_versions/:version_id/contacts_range_suggest"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ContactsRangeSuggestApplicationAppVersionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取应用版本信息
//
// - 根据应用 ID 和应用版本 ID 来获取同租户下的应用版本的信息
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/get_applicationAppVersion.go
func (a *applicationAppVersion) Get(ctx context.Context, req *GetApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*GetApplicationAppVersionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/app_versions/:version_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetApplicationAppVersionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取应用版本列表
//
// - 根据 app_id 获取对应应用版本列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/list_applicationAppVersion.go
func (a *applicationAppVersion) List(ctx context.Context, req *ListApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*ListApplicationAppVersionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/app_versions"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListApplicationAppVersionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *applicationAppVersion) ListByIterator(ctx context.Context, req *ListApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*ListApplicationAppVersionIterator, error) {
	return &ListApplicationAppVersionIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 更新应用审核状态
//
// - 通过接口来更新应用版本的审核结果：通过后应用可以直接上架；拒绝后则开发者可以看到拒绝理由，并在修改后再次申请发布。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/patch_applicationAppVersion.go
func (a *applicationAppVersion) Patch(ctx context.Context, req *PatchApplicationAppVersionReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationAppVersionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/app_versions/:version_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationAppVersionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=patch&project=application&resource=application.contacts_range&version=v6
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/patch_applicationContactsRange.go
func (a *applicationContactsRange) Patch(ctx context.Context, req *PatchApplicationContactsRangeReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationContactsRangeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/contacts_range"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationContactsRangeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取应用反馈列表
//
// - 查询应用的反馈数据
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/list_applicationFeedback.go
func (a *applicationFeedback) List(ctx context.Context, req *ListApplicationFeedbackReq, options ...larkcore.RequestOptionFunc) (*ListApplicationFeedbackResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/feedbacks"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListApplicationFeedbackResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新应用反馈
//
// - 更新应用的反馈数据
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/patch_applicationFeedback.go
func (a *applicationFeedback) Patch(ctx context.Context, req *PatchApplicationFeedbackReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationFeedbackResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/feedbacks/:feedback_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationFeedbackResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=check_white_black_list&project=application&resource=application.visibility&version=v6
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/checkWhiteBlackList_applicationVisibility.go
func (a *applicationVisibility) CheckWhiteBlackList(ctx context.Context, req *CheckWhiteBlackListApplicationVisibilityReq, options ...larkcore.RequestOptionFunc) (*CheckWhiteBlackListApplicationVisibilityResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/visibility/check_white_black_list"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CheckWhiteBlackListApplicationVisibilityResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

//
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=patch&project=application&resource=application.visibility&version=v6
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/applicationv6/patch_applicationVisibility.go
func (a *applicationVisibility) Patch(ctx context.Context, req *PatchApplicationVisibilityReq, options ...larkcore.RequestOptionFunc) (*PatchApplicationVisibilityResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/application/v6/applications/:app_id/visibility"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchApplicationVisibilityResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
