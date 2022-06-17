// Package vc code generated by oapi sdk gen
package vc

import (
	"context"
	"net/http"

	"github.com/feishu/oapi-sdk-go/core"
)

/**
构建业务域服务实例
**/
func NewService(httpClient *http.Client, config *core.Config) *VcService {
	v := &VcService{httpClient: httpClient, config: config}
	v.Meeting = &meeting{service: v}
	v.MeetingRecording = &meetingRecording{service: v}
	v.Report = &report{service: v}
	v.Reserve = &reserve{service: v}
	v.RoomConfig = &roomConfig{service: v}
	return v
}

/**
业务域服务定义
**/
type VcService struct {
	httpClient       *http.Client
	config           *core.Config
	Meeting          *meeting
	MeetingRecording *meetingRecording
	Report           *report
	Reserve          *reserve
	RoomConfig       *roomConfig
}

/**
资源服务定义
**/
type meeting struct {
	service *VcService
}
type meetingRecording struct {
	service *VcService
}
type report struct {
	service *VcService
}
type reserve struct {
	service *VcService
}
type roomConfig struct {
	service *VcService
}

/**
资源服务方法定义
**/
func (m *meeting) End(ctx context.Context, req *EndMeetingReq, options ...core.RequestOptionFunc) (*EndMeetingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/end", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &EndMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) Get(ctx context.Context, req *GetMeetingReq, options ...core.RequestOptionFunc) (*GetMeetingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/vc/v1/meetings/:meeting_id", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) Invite(ctx context.Context, req *InviteMeetingReq, options ...core.RequestOptionFunc) (*InviteMeetingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/invite", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &InviteMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) Kickout(ctx context.Context, req *KickoutMeetingReq, options ...core.RequestOptionFunc) (*KickoutMeetingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPost,
		"/open-apis/vc/v1/meetings/:meeting_id/kickout", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &KickoutMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meeting) ListByNo(ctx context.Context, req *ListByNoMeetingReq, options ...core.RequestOptionFunc) (*ListByNoMeetingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/vc/v1/meetings/list_by_no", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListByNoMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

/**如果是分页查询，则添加迭代器函数**/
func (m *meeting) ListByNoMeeting(ctx context.Context, req *ListByNoMeetingReq, options ...core.RequestOptionFunc) (*ListByNoMeetingIterator, error) {
	return &ListByNoMeetingIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.ListByNo,
		options:  options,
		limit:    req.Limit}, nil
}
func (m *meeting) SetHost(ctx context.Context, req *SetHostMeetingReq, options ...core.RequestOptionFunc) (*SetHostMeetingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/set_host", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SetHostMeetingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) Get(ctx context.Context, req *GetMeetingRecordingReq, options ...core.RequestOptionFunc) (*GetMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodGet,
		"/open-apis/vc/v1/meetings/:meeting_id/recording", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) SetPermission(ctx context.Context, req *SetPermissionMeetingRecordingReq, options ...core.RequestOptionFunc) (*SetPermissionMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/recording/set_permission", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SetPermissionMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) Start(ctx context.Context, req *StartMeetingRecordingReq, options ...core.RequestOptionFunc) (*StartMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/recording/start", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &StartMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *meetingRecording) Stop(ctx context.Context, req *StopMeetingRecordingReq, options ...core.RequestOptionFunc) (*StopMeetingRecordingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, m.service.config, http.MethodPatch,
		"/open-apis/vc/v1/meetings/:meeting_id/recording/stop", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &StopMeetingRecordingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *report) GetDaily(ctx context.Context, req *GetDailyReportReq, options ...core.RequestOptionFunc) (*GetDailyReportResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reports/get_daily", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetDailyReportResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *report) GetTopUser(ctx context.Context, req *GetTopUserReportReq, options ...core.RequestOptionFunc) (*GetTopUserReportResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reports/get_top_user", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetTopUserReportResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Apply(ctx context.Context, req *ApplyReserveReq, options ...core.RequestOptionFunc) (*ApplyReserveResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodPost,
		"/open-apis/vc/v1/reserves/apply", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ApplyReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Delete(ctx context.Context, req *DeleteReserveReq, options ...core.RequestOptionFunc) (*DeleteReserveResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodDelete,
		"/open-apis/vc/v1/reserves/:reserve_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Get(ctx context.Context, req *GetReserveReq, options ...core.RequestOptionFunc) (*GetReserveResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reserves/:reserve_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) GetActiveMeeting(ctx context.Context, req *GetActiveMeetingReserveReq, options ...core.RequestOptionFunc) (*GetActiveMeetingReserveResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/reserves/:reserve_id/get_active_meeting", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetActiveMeetingReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *reserve) Update(ctx context.Context, req *UpdateReserveReq, options ...core.RequestOptionFunc) (*UpdateReserveResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodPut,
		"/open-apis/vc/v1/reserves/:reserve_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateReserveResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *roomConfig) Query(ctx context.Context, req *QueryRoomConfigReq, options ...core.RequestOptionFunc) (*QueryRoomConfigResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodGet,
		"/open-apis/vc/v1/room_configs/query", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryRoomConfigResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (r *roomConfig) Set(ctx context.Context, req *SetRoomConfigReq, options ...core.RequestOptionFunc) (*SetRoomConfigResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, r.service.config, http.MethodPost,
		"/open-apis/vc/v1/room_configs/set", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SetRoomConfigResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}