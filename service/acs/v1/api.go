// Package acs code generated by oapi sdk gen
package acs

import (
	"bytes"
	"context"
	"net/http"

	"github.com/feishu/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *AcsService {
	a := &AcsService{httpClient: httpClient, config: config}
	a.AccessRecord = &accessRecord{service: a}
	a.AccessRecordAccessPhoto = &accessRecordAccessPhoto{service: a}
	a.Device = &device{service: a}
	a.User = &user{service: a}
	a.UserFace = &userFace{service: a}
	return a
}

/**
业务域服务定义
**/
type AcsService struct {
	httpClient              *http.Client
	config                  *core.Config
	AccessRecord            *accessRecord
	AccessRecordAccessPhoto *accessRecordAccessPhoto
	Device                  *device
	User                    *user
	UserFace                *userFace
}

/**
资源服务定义
**/
type accessRecord struct {
	service *AcsService
}
type accessRecordAccessPhoto struct {
	service *AcsService
}
type device struct {
	service *AcsService
}
type user struct {
	service *AcsService
}
type userFace struct {
	service *AcsService
}

/**
资源服务方法定义
**/
func (a *accessRecord) List(ctx context.Context, req *ListAccessRecordReq, options ...core.RequestOptionFunc) (*ListAccessRecordResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/acs/v1/access_records", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListAccessRecordResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (a *accessRecord) ListAccessRecord(ctx context.Context, req *ListAccessRecordReq, options ...core.RequestOptionFunc) (*ListAccessRecordIterator, error) {
	return &ListAccessRecordIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (a *accessRecordAccessPhoto) Get(ctx context.Context, req *GetAccessRecordAccessPhotoReq, options ...core.RequestOptionFunc) (*GetAccessRecordAccessPhotoResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, a.service.config, http.MethodGet,
		"/open-apis/acs/v1/access_records/:access_record_id/access_photo", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetAccessRecordAccessPhotoResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *device) List(ctx context.Context, options ...core.RequestOptionFunc) (*ListDeviceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, d.service.config, http.MethodGet,
		"/open-apis/acs/v1/devices", []core.AccessTokenType{core.AccessTokenTypeTenant}, nil, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDeviceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) Get(ctx context.Context, req *GetUserReq, options ...core.RequestOptionFunc) (*GetUserResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/acs/v1/users/:user_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) List(ctx context.Context, req *ListUserReq, options ...core.RequestOptionFunc) (*ListUserResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/acs/v1/users", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (u *user) ListUser(ctx context.Context, req *ListUserReq, options ...core.RequestOptionFunc) (*ListUserIterator, error) {
	return &ListUserIterator{
		ctx:      ctx,
		req:      req,
		listFunc: u.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (u *user) Patch(ctx context.Context, req *PatchUserReq, options ...core.RequestOptionFunc) (*PatchUserResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPatch,
		"/open-apis/acs/v1/users/:user_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchUserResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userFace) Get(ctx context.Context, req *GetUserFaceReq, options ...core.RequestOptionFunc) (*GetUserFaceResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodGet,
		"/open-apis/acs/v1/users/:user_id/face", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserFaceResp{RawResponse: rawResp}
	// 如果是下载，则设置响应结果
	if rawResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(rawResp.RawBody)
		resp.FileName = core.FileNameByHeader(rawResp.Header)
		return resp, err
	}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *userFace) Update(ctx context.Context, req *UpdateUserFaceReq, options ...core.RequestOptionFunc) (*UpdateUserFaceResp, error) {
	options = append(options, core.WithFileUpload())
	// 发起请求
	rawResp, err := core.SendRequest(ctx, u.service.config, http.MethodPut,
		"/open-apis/acs/v1/users/:user_id/face", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateUserFaceResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}