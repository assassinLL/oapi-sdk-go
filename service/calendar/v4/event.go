// Code generated by lark suite oapi sdk gen
package v4

import (
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/event"
)

type CalendarChangedEventHandler struct {
	Fn func(*core.Context, *CalendarChangedEvent) error
}

func (h *CalendarChangedEventHandler) GetEvent() interface{} {
	return &CalendarChangedEvent{}
}

func (h *CalendarChangedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*CalendarChangedEvent))
}

func SetCalendarChangedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *CalendarChangedEvent) error) {
	event.SetTypeHandler(conf, "calendar.calendar.changed_v4", &CalendarChangedEventHandler{Fn: fn})
}

type CalendarEventChangedEventHandler struct {
	Fn func(*core.Context, *CalendarEventChangedEvent) error
}

func (h *CalendarEventChangedEventHandler) GetEvent() interface{} {
	return &CalendarEventChangedEvent{}
}

func (h *CalendarEventChangedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*CalendarEventChangedEvent))
}

func SetCalendarEventChangedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *CalendarEventChangedEvent) error) {
	event.SetTypeHandler(conf, "calendar.calendar.event.changed_v4", &CalendarEventChangedEventHandler{Fn: fn})
}

type CalendarAclCreatedEventHandler struct {
	Fn func(*core.Context, *CalendarAclCreatedEvent) error
}

func (h *CalendarAclCreatedEventHandler) GetEvent() interface{} {
	return &CalendarAclCreatedEvent{}
}

func (h *CalendarAclCreatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*CalendarAclCreatedEvent))
}

func SetCalendarAclCreatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *CalendarAclCreatedEvent) error) {
	event.SetTypeHandler(conf, "calendar.calendar.acl.created_v4", &CalendarAclCreatedEventHandler{Fn: fn})
}

type CalendarAclDeletedEventHandler struct {
	Fn func(*core.Context, *CalendarAclDeletedEvent) error
}

func (h *CalendarAclDeletedEventHandler) GetEvent() interface{} {
	return &CalendarAclDeletedEvent{}
}

func (h *CalendarAclDeletedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*CalendarAclDeletedEvent))
}

func SetCalendarAclDeletedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *CalendarAclDeletedEvent) error) {
	event.SetTypeHandler(conf, "calendar.calendar.acl.deleted_v4", &CalendarAclDeletedEventHandler{Fn: fn})
}
