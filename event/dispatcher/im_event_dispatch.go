// Package dispatcher code generated by oapi sdk gen
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

package dispatcher

import (
	"context"
	"github.com/assassinLL/oapi-sdk-go/v3/service/im/v1"
)

// 群解散
//
// - 群组被解散后触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==解散群== 事件;- 事件会向群内订阅了该事件的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/disbanded
func (dispatcher *EventDispatcher) OnP2ChatDisbandedV1(handler func(ctx context.Context, event *larkim.P2ChatDisbandedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.disbanded_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.disbanded_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.disbanded_v1"] = larkim.NewP2ChatDisbandedV1Handler(handler)
	return dispatcher
}

// 群配置修改
//
// - 群组配置被修改后触发此事件，包含：;- 群主转移;- 群基本信息修改(群头像/群名称/群描述/群国际化名称);- 群权限修改(加人入群权限/群编辑权限/at所有人权限/群分享权限)。
//
// - 注意事项：; - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==群配置修改== 事件;- 事件会向群内订阅了该事件的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/updated
func (dispatcher *EventDispatcher) OnP2ChatUpdatedV1(handler func(ctx context.Context, event *larkim.P2ChatUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.updated_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.updated_v1"] = larkim.NewP2ChatUpdatedV1Handler(handler)
	return dispatcher
}

// 机器人进群
//
// - 机器人被用户添加至群聊时触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==机器人进群== 事件;- 事件会向进群的机器人进行推送;- 机器人邀请机器人不会触发事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/added
func (dispatcher *EventDispatcher) OnP2ChatMemberBotAddedV1(handler func(ctx context.Context, event *larkim.P2ChatMemberBotAddedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.bot.added_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.bot.added_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.bot.added_v1"] = larkim.NewP2ChatMemberBotAddedV1Handler(handler)
	return dispatcher
}

// 机器人被移出群
//
// - 机器人被移出群聊后触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==机器人被移出群== 事件;- 事件会向被移出群的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/deleted
func (dispatcher *EventDispatcher) OnP2ChatMemberBotDeletedV1(handler func(ctx context.Context, event *larkim.P2ChatMemberBotDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.bot.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.bot.deleted_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.bot.deleted_v1"] = larkim.NewP2ChatMemberBotDeletedV1Handler(handler)
	return dispatcher
}

// 用户进群
//
// - 新用户进群（包含话题群）触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==用户进群== 事件;- 事件会向群内订阅了该事件的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/added
func (dispatcher *EventDispatcher) OnP2ChatMemberUserAddedV1(handler func(ctx context.Context, event *larkim.P2ChatMemberUserAddedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.user.added_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.user.added_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.user.added_v1"] = larkim.NewP2ChatMemberUserAddedV1Handler(handler)
	return dispatcher
}

// 用户出群
//
// - 用户主动退群或被移出群聊时推送事件。
//
// - 注意事项：;- 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)并且机器人所在群发生上述变化;- 机器人需要订阅 ==消息与群组== 分类下的 ==用户主动退群或被移出群聊== 事件;- 事件会向群内订阅了该事件的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/deleted
func (dispatcher *EventDispatcher) OnP2ChatMemberUserDeletedV1(handler func(ctx context.Context, event *larkim.P2ChatMemberUserDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.user.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.user.deleted_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.user.deleted_v1"] = larkim.NewP2ChatMemberUserDeletedV1Handler(handler)
	return dispatcher
}

// 撤销拉用户进群
//
// - 撤销拉用户进群后触发此事件。
//
// - 注意事项：;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability);- 需要订阅 ==消息与群组== 分类下的 ==撤销拉用户进群== 事件;- 事件会向群内订阅了该事件的机器人进行推送
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/withdrawn
func (dispatcher *EventDispatcher) OnP2ChatMemberUserWithdrawnV1(handler func(ctx context.Context, event *larkim.P2ChatMemberUserWithdrawnV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.chat.member.user.withdrawn_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.chat.member.user.withdrawn_v1")
	}
	dispatcher.eventType2EventHandler["im.chat.member.user.withdrawn_v1"] = larkim.NewP2ChatMemberUserWithdrawnV1Handler(handler)
	return dispatcher
}

// 消息已读
//
// - 用户阅读机器人发送的单聊消息后触发此事件。
//
// - 注意事项:;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)  ;- 需要订阅 ==消息与群组== 分类下的 ==消息已读== 事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/message_read
func (dispatcher *EventDispatcher) OnP2MessageReadV1(handler func(ctx context.Context, event *larkim.P2MessageReadV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.message_read_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.message_read_v1")
	}
	dispatcher.eventType2EventHandler["im.message.message_read_v1"] = larkim.NewP2MessageReadV1Handler(handler)
	return dispatcher
}

// 消息撤回事件
//
// - 消息被撤回后触发此事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/recalled
func (dispatcher *EventDispatcher) OnP2MessageRecalledV1(handler func(ctx context.Context, event *larkim.P2MessageRecalledV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.recalled_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.recalled_v1")
	}
	dispatcher.eventType2EventHandler["im.message.recalled_v1"] = larkim.NewP2MessageRecalledV1Handler(handler)
	return dispatcher
}

// 接收消息
//
// - 机器人接收到用户发送的消息后触发此事件。
//
// - 注意事项:;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)，并订阅 ==消息与群组== 分类下的 ==接收消息v2.0== 事件才可接收推送;- 同时，将根据应用具备的权限，判断可推送的信息：;	- 当具备==获取用户发给机器人的单聊消息==权限或者==读取用户发给机器人的单聊消息（历史权限）==，可接收与机器人单聊会话中用户发送的所有消息;	- 当具备==获取群组中所有消息== 权限时，可接收与机器人所在群聊会话中用户发送的所有消息;	- 当具备==获取用户在群组中@机器人的消息== 权限或者==获取用户在群聊中@机器人的消息（历史权限）==，可接收机器人所在群聊中用户 @ 机器人的消息
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
func (dispatcher *EventDispatcher) OnP2MessageReceiveV1(handler func(ctx context.Context, event *larkim.P2MessageReceiveV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.receive_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.receive_v1")
	}
	dispatcher.eventType2EventHandler["im.message.receive_v1"] = larkim.NewP2MessageReceiveV1Handler(handler)
	return dispatcher
}

// 新增消息表情回复
//
// - 消息被添加某一个表情回复后触发此事件
//
// - 注意事项:;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)  ;- 具备==获取单聊、群组消息== 或 ==获取与发送单聊、群组消息==权限，并订阅 ==消息与群组== 分类下的 ==消息被reaction== 事件才可接收推送;- 机器人只能收到所在群聊内的消息被添加表情回复事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/events/created
func (dispatcher *EventDispatcher) OnP2MessageReactionCreatedV1(handler func(ctx context.Context, event *larkim.P2MessageReactionCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.reaction.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.reaction.created_v1")
	}
	dispatcher.eventType2EventHandler["im.message.reaction.created_v1"] = larkim.NewP2MessageReactionCreatedV1Handler(handler)
	return dispatcher
}

// 删除消息表情回复
//
// - 消息被删除某一个表情回复后触发此事件
//
// - 注意事项:;- 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)  ;- 具备==获取单聊、群组消息== 或 ==获取与发送单聊、群组消息==权限，并订阅 ==消息与群组== 分类下的 ==消息被取消reaction== 事件才可接收推送;- 机器人只能收到所在群聊内的消息被删除表情回复事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/events/deleted
func (dispatcher *EventDispatcher) OnP2MessageReactionDeletedV1(handler func(ctx context.Context, event *larkim.P2MessageReactionDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["im.message.reaction.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "im.message.reaction.deleted_v1")
	}
	dispatcher.eventType2EventHandler["im.message.reaction.deleted_v1"] = larkim.NewP2MessageReactionDeletedV1Handler(handler)
	return dispatcher
}
