// Package acs code generated by oapi sdk gen
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

package larkacs

import (
	"bytes"
	"context"
	"net/http"

	"github.com/assassinLL/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *AcsService {
	a := &AcsService{config: config}
	a.AccessRecord = &accessRecord{service: a}
	a.AccessRecordAccessPhoto = &accessRecordAccessPhoto{service: a}
	a.Device = &device{service: a}
	a.User = &user{service: a}
	a.UserFace = &userFace{service: a}
	return a
}

type AcsService struct {
	config                  *larkcore.Config
	AccessRecord            *accessRecord            // 门禁记录
	AccessRecordAccessPhoto *accessRecordAccessPhoto // access_record.access_photo
	Device                  *device                  // 门禁设备
	User                    *user                    // 用户管理
	UserFace                *userFace                // user.face
}

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

// 获取门禁记录列表
//
// - 用户在门禁考勤机上成功开门或打卡后，智能门禁应用都会生成一条门禁记录。;;该接口返回满足查询参数的识别记录。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record/list
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/list_accessRecord.go
func (a *accessRecord) List(ctx context.Context, req *ListAccessRecordReq, options ...larkcore.RequestOptionFunc) (*ListAccessRecordResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/access_records"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListAccessRecordResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *accessRecord) ListByIterator(ctx context.Context, req *ListAccessRecordReq, options ...larkcore.RequestOptionFunc) (*ListAccessRecordIterator, error) {
	return &ListAccessRecordIterator{
		ctx:      ctx,
		req:      req,
		listFunc: a.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 下载开门时的人脸识别图片
//
// - 用户在门禁考勤机上成功开门或打卡后，智能门禁应用都会生成一条门禁记录，对于使用人脸识别方式进行开门的识别记录，还会有抓拍图。;;可以用该接口下载开门时的人脸识别照片。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/access_record-access_photo/get
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/get_accessRecordAccessPhoto.go
func (a *accessRecordAccessPhoto) Get(ctx context.Context, req *GetAccessRecordAccessPhotoReq, options ...larkcore.RequestOptionFunc) (*GetAccessRecordAccessPhotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/access_records/:access_record_id/access_photo"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetAccessRecordAccessPhotoResp{ApiResp: apiResp}
	// 如果是下载，则设置响应结果
	if apiResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(apiResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(apiResp.Header)
		return resp, err
	}
	err = apiResp.JSONUnmarshalBody(resp, a.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取门禁设备列表
//
// - 使用该接口获取租户内所有门禁设备。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/device/list
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/list_device.go
func (d *device) List(ctx context.Context, options ...larkcore.RequestOptionFunc) (*ListDeviceResp, error) {
	// 发起请求
	apiReq := &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	apiReq.ApiPath = "/open-apis/acs/v1/devices"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, d.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDeviceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, d.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取单个用户信息
//
// - 该接口用于获取智能门禁中单个用户的信息。
//
// - 只能获取已加入智能门禁权限组的用户
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/get
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/get_user.go
func (u *user) Get(ctx context.Context, req *GetUserReq, options ...larkcore.RequestOptionFunc) (*GetUserResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/users/:user_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, u.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取用户列表
//
// - 使用该接口获取智能门禁中所有用户信息。
//
// - 只能获取已加入智能门禁权限组的用户。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/list
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/list_user.go
func (u *user) List(ctx context.Context, req *ListUserReq, options ...larkcore.RequestOptionFunc) (*ListUserResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/users"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListUserResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, u.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (u *user) ListByIterator(ctx context.Context, req *ListUserReq, options ...larkcore.RequestOptionFunc) (*ListUserIterator, error) {
	return &ListUserIterator{
		ctx:      ctx,
		req:      req,
		listFunc: u.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 修改用户部分信息
//
// - 飞书智能门禁在人脸识别成功后会有韦根信号输出，输出用户的卡号。;对于使用韦根协议的门禁系统，企业可使用该接口录入用户卡号。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/patch
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/patch_user.go
func (u *user) Patch(ctx context.Context, req *PatchUserReq, options ...larkcore.RequestOptionFunc) (*PatchUserResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/users/:user_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchUserResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, u.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 下载人脸图片
//
// - 对于已经录入人脸图片的用户，可以使用该接口下载用户人脸图片。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/get
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/get_userFace.go
func (u *userFace) Get(ctx context.Context, req *GetUserFaceReq, options ...larkcore.RequestOptionFunc) (*GetUserFaceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/users/:user_id/face"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserFaceResp{ApiResp: apiResp}
	// 如果是下载，则设置响应结果
	if apiResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(apiResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(apiResp.Header)
		return resp, err
	}
	err = apiResp.JSONUnmarshalBody(resp, u.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 上传人脸图片
//
// - 用户需要录入人脸图片才可以使用门禁考勤机。使用该 API 上传门禁用户的人脸图片。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user-face/update
//
// - 使用Demo链接:https://github.com/assassinLL/oapi-sdk-go/tree/v3_main/sample/apiall/acsv1/update_userFace.go
func (u *userFace) Update(ctx context.Context, req *UpdateUserFaceReq, options ...larkcore.RequestOptionFunc) (*UpdateUserFaceResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/acs/v1/users/:user_id/face"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateUserFaceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, u.service.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
