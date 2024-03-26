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
	"github.com/larksuite/oapi-sdk-go/v3/service/hire/v1"
)

// 删除投递
//
// - 删除投递
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/events/deleted
func (dispatcher *EventDispatcher) OnP2ApplicationDeletedV1(handler func(ctx context.Context, event *larkhire.P2ApplicationDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.application.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.application.deleted_v1")
	}
	dispatcher.eventType2EventHandler["hire.application.deleted_v1"] = larkhire.NewP2ApplicationDeletedV1Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/event/application-stage-changed
func (dispatcher *EventDispatcher) OnP2ApplicationStageChangedV1(handler func(ctx context.Context, event *larkhire.P2ApplicationStageChangedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.application.stage_changed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.application.stage_changed_v1")
	}
	dispatcher.eventType2EventHandler["hire.application.stage_changed_v1"] = larkhire.NewP2ApplicationStageChangedV1Handler(handler)
	return dispatcher
}

// 帐号绑定
//
// - 招聘管理员添加三方服务商帐号时，系统会推送事件给应用开发者，开发者可根据事件获取用户添加的帐号类型（背调 或 笔试）和 帐号自定义字段信息，并根据这些信息识别用户在服务商处的身份，完成三方服务商帐号 和 招聘帐号之间的绑定，并根据用户服务商身份推送对应的背调套餐或试卷列表。
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_account/events/created
func (dispatcher *EventDispatcher) OnP2EcoAccountCreatedV1(handler func(ctx context.Context, event *larkhire.P2EcoAccountCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.eco_account.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.eco_account.created_v1")
	}
	dispatcher.eventType2EventHandler["hire.eco_account.created_v1"] = larkhire.NewP2EcoAccountCreatedV1Handler(handler)
	return dispatcher
}

// 终止背调
//
// - 用户在招聘系统终止背调后，系统会推送事件给对应的应用开发者。开发者可根据事件获取背调 ID，完成在三方服务商处的订单取消等后续操作。
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check/events/canceled
func (dispatcher *EventDispatcher) OnP2EcoBackgroundCheckCanceledV1(handler func(ctx context.Context, event *larkhire.P2EcoBackgroundCheckCanceledV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.eco_background_check.canceled_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.eco_background_check.canceled_v1")
	}
	dispatcher.eventType2EventHandler["hire.eco_background_check.canceled_v1"] = larkhire.NewP2EcoBackgroundCheckCanceledV1Handler(handler)
	return dispatcher
}

// 创建背调
//
// - 用户在招聘系统安排背调后，系统会推送事件给对应的应用开发者。开发者可根据事件获取候选人信息、委托人信息和自定义字段信息，并根据这些信息完成在三方服务商处的背调订单创建。
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check/events/created
func (dispatcher *EventDispatcher) OnP2EcoBackgroundCheckCreatedV1(handler func(ctx context.Context, event *larkhire.P2EcoBackgroundCheckCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.eco_background_check.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.eco_background_check.created_v1")
	}
	dispatcher.eventType2EventHandler["hire.eco_background_check.created_v1"] = larkhire.NewP2EcoBackgroundCheckCreatedV1Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2EcoExamCreatedV1(handler func(ctx context.Context, event *larkhire.P2EcoExamCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.eco_exam.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.eco_exam.created_v1")
	}
	dispatcher.eventType2EventHandler["hire.eco_exam.created_v1"] = larkhire.NewP2EcoExamCreatedV1Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/event/import-ehr
func (dispatcher *EventDispatcher) OnP2EhrImportTaskImportedV1(handler func(ctx context.Context, event *larkhire.P2EhrImportTaskImportedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.ehr_import_task.imported_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.ehr_import_task.imported_v1")
	}
	dispatcher.eventType2EventHandler["hire.ehr_import_task.imported_v1"] = larkhire.NewP2EhrImportTaskImportedV1Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2EhrImportTaskForInternshipOfferImportedV1(handler func(ctx context.Context, event *larkhire.P2EhrImportTaskForInternshipOfferImportedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.ehr_import_task_for_internship_offer.imported_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.ehr_import_task_for_internship_offer.imported_v1")
	}
	dispatcher.eventType2EventHandler["hire.ehr_import_task_for_internship_offer.imported_v1"] = larkhire.NewP2EhrImportTaskForInternshipOfferImportedV1Handler(handler)
	return dispatcher
}

// Offer 状态变更
//
// - 当 Offer 状态发生变更时将触发该事件。
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/offer/events/status_changed
func (dispatcher *EventDispatcher) OnP2OfferStatusChangedV1(handler func(ctx context.Context, event *larkhire.P2OfferStatusChangedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.offer.status_changed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.offer.status_changed_v1")
	}
	dispatcher.eventType2EventHandler["hire.offer.status_changed_v1"] = larkhire.NewP2OfferStatusChangedV1Handler(handler)
	return dispatcher
}

// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ReferralAccountAssetsUpdateV1(handler func(ctx context.Context, event *larkhire.P2ReferralAccountAssetsUpdateV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["hire.referral_account.assets_update_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "hire.referral_account.assets_update_v1")
	}
	dispatcher.eventType2EventHandler["hire.referral_account.assets_update_v1"] = larkhire.NewP2ReferralAccountAssetsUpdateV1Handler(handler)
	return dispatcher
}
