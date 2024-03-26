// Code generated by Lark OpenAPI.

package larkwiki

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
)

type V2 struct {
	Space        *space        // 知识空间
	SpaceMember  *spaceMember  // 空间成员
	SpaceNode    *spaceNode    // 节点
	SpaceSetting *spaceSetting // 空间设置
	Task         *task         // 云文档
}

func New(config *larkcore.Config) *V2 {
	return &V2{
		Space:        &space{config: config},
		SpaceMember:  &spaceMember{config: config},
		SpaceNode:    &spaceNode{config: config},
		SpaceSetting: &spaceSetting{config: config},
		Task:         &task{config: config},
	}
}

type space struct {
	config *larkcore.Config
}
type spaceMember struct {
	config *larkcore.Config
}
type spaceNode struct {
	config *larkcore.Config
}
type spaceSetting struct {
	config *larkcore.Config
}
type task struct {
	config *larkcore.Config
}

// Create 创建知识空间
//
// - 此接口用于创建知识空间
//
// - 此接口不支持tenant access token（应用身份访问）
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/create_space.go
func (s *space) Create(ctx context.Context, req *CreateSpaceReq, options ...larkcore.RequestOptionFunc) (*CreateSpaceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpaceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Get 获取知识空间信息
//
// - 此接口用于根据知识空间ID来查询知识空间的信息。;;空间类型（type）：;- 个人空间：归个人管理。一人仅可拥有一个个人空间，无法添加其他管理员。;- 团队空间：归团队（多人)管理，可添加多个管理员。;;空间可见性（visibility）：;- 公开空间：租户所有用户可见，默认为成员权限。无法额外添加成员，但可以添加管理员。;- 私有空间：仅对知识空间管理员、成员可见，需要手动添加管理员、成员。
//
// - 本接口要求知识库权限：;- 需要为知识空间成员（管理员）
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/get_space.go
func (s *space) Get(ctx context.Context, req *GetSpaceReq, options ...larkcore.RequestOptionFunc) (*GetSpaceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpaceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// GetNode 获取知识空间节点信息
//
// - 获取知识空间节点信息
//
// - 知识库权限要求，当前使用的 access token 所代表的应用或用户拥有：;- 节点阅读权限
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get_node
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/getNode_space.go
func (s *space) GetNode(ctx context.Context, req *GetNodeSpaceReq, options ...larkcore.RequestOptionFunc) (*GetNodeSpaceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/get_node"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetNodeSpaceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// List 获取知识空间列表
//
// - 此接口用于获取有权限访问的知识空间列表。;;此接口为分页接口。由于权限过滤，可能返回列表为空，但分页标记（has_more）为true，可以继续分页请求。;;对于知识空间各项属性描述请参阅[获取知识空间信息](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get)
//
// - 使用tenant access token调用时，请确认应用/机器人拥有部分知识空间的访问权限，否则返回列表容易为空。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/list_space.go
func (s *space) List(ctx context.Context, req *ListSpaceReq, options ...larkcore.RequestOptionFunc) (*ListSpaceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListSpaceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *space) ListByIterator(ctx context.Context, req *ListSpaceReq, options ...larkcore.RequestOptionFunc) (*ListSpaceIterator, error) {
	return &ListSpaceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: s.List,
		options:  options,
		limit:    req.Limit}, nil
}

// Create 添加知识空间成员
//
// - 添加知识空间成员或管理员。
//
// - 知识空间具有[类型](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)和[可见性](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)的概念。不同的类型或可见性可以对本操作做出限制：;- 可见性限制：公开知识空间（visibility为public）对租户所有用户可见，因此不支持再添加成员，但可以添加管理员。;- 类型限制：个人知识空间 （type为person）为个人管理的知识空间，不支持添加其他管理员（包括应用/机器人）。但可以添加成员。;;;知识空间权限要求，当前使用的 access token 所代表的应用或用户拥有：;- 为知识空间管理员
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/create_spaceMember.go
func (s *spaceMember) Create(ctx context.Context, req *CreateSpaceMemberReq, options ...larkcore.RequestOptionFunc) (*CreateSpaceMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/members"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpaceMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Delete 删除知识空间成员
//
// - 此接口用于删除知识空间成员或管理员。
//
// - 知识空间具有[类型](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)和[可见性](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)的概念。不同的类型或可见性可以对本操作做出限制：;- 可见性限制：公开知识空间（visibility为public）对租户所有用户可见，因此不支持再删除成员，但可以删除管理员。;- 类型限制：个人知识空间 （type为person）为个人管理的知识空间，不支持删除管理员。但可以删除成员。;;;知识空间权限要求，当前使用的 access token 所代表的应用或用户拥有：;- 为知识空间管理员
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/delete_spaceMember.go
func (s *spaceMember) Delete(ctx context.Context, req *DeleteSpaceMemberReq, options ...larkcore.RequestOptionFunc) (*DeleteSpaceMemberResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/members/:member_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpaceMemberResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Copy 创建知识空间节点副本
//
// - 此接口用于在知识空间创建节点副本到指定位置。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/copy
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/copy_spaceNode.go
func (s *spaceNode) Copy(ctx context.Context, req *CopySpaceNodeReq, options ...larkcore.RequestOptionFunc) (*CopySpaceNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/copy"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CopySpaceNodeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create 创建知识空间节点
//
// - 此接口用于在知识节点里创建[节点](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)到指定位置。
//
// - 知识空间权限要求，当前使用的 access token 所代表的应用或用户拥有：;- **父节点**容器编辑权限
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/create_spaceNode.go
func (s *spaceNode) Create(ctx context.Context, req *CreateSpaceNodeReq, options ...larkcore.RequestOptionFunc) (*CreateSpaceNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/nodes"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpaceNodeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// List 获取知识空间子节点列表
//
// - 此接口用于分页获取Wiki节点的子节点列表。;;此接口为分页接口。由于权限过滤，可能返回列表为空，但分页标记（has_more）为true，可以继续分页请求。
//
// - 知识库权限要求，当前使用的 access token 所代表的应用或用户拥有：;- 父节点阅读权限
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/list_spaceNode.go
func (s *spaceNode) List(ctx context.Context, req *ListSpaceNodeReq, options ...larkcore.RequestOptionFunc) (*ListSpaceNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/nodes"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListSpaceNodeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *spaceNode) ListByIterator(ctx context.Context, req *ListSpaceNodeReq, options ...larkcore.RequestOptionFunc) (*ListSpaceNodeIterator, error) {
	return &ListSpaceNodeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: s.List,
		options:  options,
		limit:    req.Limit}, nil
}

// Move 移动知识空间节点
//
// - 此方法用于在Wiki内移动节点，支持跨知识空间移动。如果有子节点，会携带子节点一起移动。
//
// - 知识库权限要求：;- 节点编辑权限;- 原父节点容器编辑权限;- 目的父节点容器编辑权限
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/move_spaceNode.go
func (s *spaceNode) Move(ctx context.Context, req *MoveSpaceNodeReq, options ...larkcore.RequestOptionFunc) (*MoveSpaceNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/move"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveSpaceNodeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// MoveDocsToWiki 移动云空间文档至知识空间
//
// - 该接口允许移动云空间文档至知识空间，并挂载在指定位置
//
// - ### 移动操作 ###;移动后，文档将从“我的空间”或“共享空间”转移至“知识库”后，无法从下列入口查看到文档：;- 云空间主页：快速访问;- 我的空间;- 共享空间;;### 权限变更 ###;移动后，文档会向所有可查看“页面树”的用户显示，默认继承父页面的权限设置。;</md-alert
//
// - 此接口为异步接口。若移动已完成（或文档已在Wiki中），则直接返回结果（Wiki token）。若尚未完成，则返回task id。请使用[获取任务结果](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get)接口进行查询。;;知识库权限要求，当前使用的 access token 所代表的应用或用户拥有：;- 文档可管理权限;- 原文件夹编辑权限;- 目标父节点容器编辑权限
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move_docs_to_wiki
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/moveDocsToWiki_spaceNode.go
func (s *spaceNode) MoveDocsToWiki(ctx context.Context, req *MoveDocsToWikiSpaceNodeReq, options ...larkcore.RequestOptionFunc) (*MoveDocsToWikiSpaceNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/nodes/move_docs_to_wiki"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveDocsToWikiSpaceNodeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// UpdateTitle 更新知识空间节点标题
//
// - 此接口用于更新节点标题
//
// - 此接口目前仅支持文档(doc)、新版文档(docx)和快捷方式。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/update_title
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/updateTitle_spaceNode.go
func (s *spaceNode) UpdateTitle(ctx context.Context, req *UpdateTitleSpaceNodeReq, options ...larkcore.RequestOptionFunc) (*UpdateTitleSpaceNodeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/update_title"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateTitleSpaceNodeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Update 更新知识空间设置
//
// - 根据space_id更新知识空间公共设置
//
// - 知识库权限要求：;- 为知识空间管理员
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-setting/update
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/update_spaceSetting.go
func (s *spaceSetting) Update(ctx context.Context, req *UpdateSpaceSettingReq, options ...larkcore.RequestOptionFunc) (*UpdateSpaceSettingResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/spaces/:space_id/setting"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateSpaceSettingResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, s.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Get 获取任务结果
//
// - 该方法用于获取wiki异步任务的结果
//
// - 知识库权限要求，当前 access token 所代表的用户或应用（机器人）：;- 为任务创建者
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/task/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/wikiv2/get_task.go
func (t *task) Get(ctx context.Context, req *GetTaskReq, options ...larkcore.RequestOptionFunc) (*GetTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/wiki/v2/tasks/:task_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, t.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
