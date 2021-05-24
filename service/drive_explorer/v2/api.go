// Code generated by lark suite oapi sdk gen
package v2

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf    *config.Config
	Files   *FileService
	Folders *FolderService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Files = newFileService(s)
	s.Folders = newFolderService(s)
	return s
}

type FileService struct {
	service *Service
}

func newFileService(service *Service) *FileService {
	return &FileService{
		service: service,
	}
}

type FolderService struct {
	service *Service
}

func newFolderService(service *Service) *FolderService {
	return &FolderService{
		service: service,
	}
}

type FolderChildrenReqCall struct {
	ctx         *core.Context
	folders     *FolderService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *FolderChildrenReqCall) SetFolderToken(folderToken string) {
	rc.pathParams["folderToken"] = folderToken
}
func (rc *FolderChildrenReqCall) SetTypes(types ...string) {
	rc.queryParams["types"] = types
}

func (rc *FolderChildrenReqCall) Do() (*FolderChildrenResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &FolderChildrenResult{}
	req := request.NewRequest("drive/explorer/v2/folder/:folderToken/children", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.folders.service.conf, req)
	return result, err
}

func (folders *FolderService) Children(ctx *core.Context, optFns ...request.OptFn) *FolderChildrenReqCall {
	return &FolderChildrenReqCall{
		ctx:         ctx,
		folders:     folders,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type FileCopyReqCall struct {
	ctx        *core.Context
	files      *FileService
	body       *FileCopyReqBody
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *FileCopyReqCall) SetFileToken(fileToken string) {
	rc.pathParams["fileToken"] = fileToken
}

func (rc *FileCopyReqCall) Do() (*FileCopyResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &FileCopyResult{}
	req := request.NewRequest("drive/explorer/v2/file/copy/files/:fileToken", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) Copy(ctx *core.Context, body *FileCopyReqBody, optFns ...request.OptFn) *FileCopyReqCall {
	return &FileCopyReqCall{
		ctx:        ctx,
		files:      files,
		body:       body,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type FolderCreateReqCall struct {
	ctx        *core.Context
	folders    *FolderService
	body       *FolderCreateReqBody
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *FolderCreateReqCall) SetFolderToken(folderToken string) {
	rc.pathParams["folderToken"] = folderToken
}

func (rc *FolderCreateReqCall) Do() (*FolderCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &FolderCreateResult{}
	req := request.NewRequest("drive/explorer/v2/folder/:folderToken", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.folders.service.conf, req)
	return result, err
}

func (folders *FolderService) Create(ctx *core.Context, body *FolderCreateReqBody, optFns ...request.OptFn) *FolderCreateReqCall {
	return &FolderCreateReqCall{
		ctx:        ctx,
		folders:    folders,
		body:       body,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type FileCreateReqCall struct {
	ctx        *core.Context
	files      *FileService
	body       *FileCreateReqBody
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *FileCreateReqCall) SetFolderToken(folderToken string) {
	rc.pathParams["folderToken"] = folderToken
}

func (rc *FileCreateReqCall) Do() (*FileCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &FileCreateResult{}
	req := request.NewRequest("drive/explorer/v2/file/:folderToken", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) Create(ctx *core.Context, body *FileCreateReqBody, optFns ...request.OptFn) *FileCreateReqCall {
	return &FileCreateReqCall{
		ctx:        ctx,
		files:      files,
		body:       body,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type FileDocsDeleteReqCall struct {
	ctx        *core.Context
	files      *FileService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *FileDocsDeleteReqCall) SetDocToken(docToken string) {
	rc.pathParams["docToken"] = docToken
}

func (rc *FileDocsDeleteReqCall) Do() (*FileDocsDeleteResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &FileDocsDeleteResult{}
	req := request.NewRequest("drive/explorer/v2/file/docs/:docToken", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) DocsDelete(ctx *core.Context, optFns ...request.OptFn) *FileDocsDeleteReqCall {
	return &FileDocsDeleteReqCall{
		ctx:        ctx,
		files:      files,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type FolderMetaReqCall struct {
	ctx        *core.Context
	folders    *FolderService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *FolderMetaReqCall) SetFolderToken(folderToken string) {
	rc.pathParams["folderToken"] = folderToken
}

func (rc *FolderMetaReqCall) Do() (*FolderMetaResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &FolderMetaResult{}
	req := request.NewRequest("drive/explorer/v2/folder/:folderToken/meta", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.folders.service.conf, req)
	return result, err
}

func (folders *FolderService) Meta(ctx *core.Context, optFns ...request.OptFn) *FolderMetaReqCall {
	return &FolderMetaReqCall{
		ctx:        ctx,
		folders:    folders,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type FolderRootMetaReqCall struct {
	ctx     *core.Context
	folders *FolderService
	optFns  []request.OptFn
}

func (rc *FolderRootMetaReqCall) Do() (*FolderRootMetaResult, error) {
	var result = &FolderRootMetaResult{}
	req := request.NewRequest("drive/explorer/v2/root_folder/meta", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.folders.service.conf, req)
	return result, err
}

func (folders *FolderService) RootMeta(ctx *core.Context, optFns ...request.OptFn) *FolderRootMetaReqCall {
	return &FolderRootMetaReqCall{
		ctx:     ctx,
		folders: folders,
		optFns:  optFns,
	}
}

type FileSpreadsheetsDeleteReqCall struct {
	ctx        *core.Context
	files      *FileService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *FileSpreadsheetsDeleteReqCall) SetSpreadsheetToken(spreadsheetToken string) {
	rc.pathParams["spreadsheetToken"] = spreadsheetToken
}

func (rc *FileSpreadsheetsDeleteReqCall) Do() (*FileSpreadsheetsDeleteResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &FileSpreadsheetsDeleteResult{}
	req := request.NewRequest("drive/explorer/v2/file/spreadsheets/:spreadsheetToken", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.files.service.conf, req)
	return result, err
}

func (files *FileService) SpreadsheetsDelete(ctx *core.Context, optFns ...request.OptFn) *FileSpreadsheetsDeleteReqCall {
	return &FileSpreadsheetsDeleteReqCall{
		ctx:        ctx,
		files:      files,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}
