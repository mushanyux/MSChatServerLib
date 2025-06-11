package common

import (
	"bytes"
	"encoding/json"
	"errors"
)

var ErrData = errors.New("数据格式有误！")

type ChannelType uint8 // 频道类型

const (
	ChannelTypeNone            ChannelType = iota // 没有指定频道
	ChannelTypePerson                             // 个人频道
	ChannelTypeGroup                              // 群频道
	ChannelTypeCustomerService                    // 客服频道
	ChannelTypeCommunity                          // 社区
	ChannelTypeCommunityTopic                     // 话题
	ChannelTypeInfo                               // 资讯
)

type GroupMemberStatus int // 群成员状态

const (
	GroupMemberStatusNormal    GroupMemberStatus = 1 // 正常
	GroupMemberStatusBlacklist GroupMemberStatus = 2 // 黑名单
)

type GroupMemberRole int // 群成员角色

const (
	GroupMemberRoleCreater GroupMemberRole = 1 // 群主
	GroupMemberRoleManager GroupMemberRole = 2 // 管理员
	GroupMemberRoleNormal  GroupMemberRole = 0 // 群员
)

type GroupAllowViewHistoryMsgStatus int

const (
	GroupAllowViewHistoryMsgDisabled GroupAllowViewHistoryMsgStatus = 0
	GroupAllowViewHistoryMsgEnabled  GroupAllowViewHistoryMsgStatus = 1
)

func (c ChannelType) Uint8() uint8 {
	return uint8(c)
}

const (
	GroupMemberSeqKey        = "groupMember"           // 群成员序列号key
	GroupSettingSeqKey       = "groupSetting"          // 群设置序列号key
	GroupSeqKey              = "group"                 // 群序列号key
	UserSettingSeqKey        = "userSetting"           // 用户设置序列号key
	UserSeqKey               = "user"                  // 用户序列号
	FriendSeqKey             = "friend"                // 好友
	MessageExtraSeqKey       = "messageExtra"          // 消息扩展序号
	MessageReactionSeqKey    = "messageReaction"       // 消息回应序号
	RobotSeqKey              = "robot"                 // 机器人序号
	RobotEventSeqKey         = "robotEventSeq"         // 机器人事件序号
	SensitiveWordsKey        = "sensitiveWords"        // 敏感词序号
	RemindersKey             = "reminders"             // 提醒项序号
	SyncConversationExtraKey = "syncConversationExtra" // 同步最近会话扩展
	ProhibitWordKey          = "ProhibitWord"          // 违禁词
)

// 群属性key
const (
	GroupAttrKeyName               = "name"                        // 群名称
	GroupAttrKeyNotice             = "notice"                      // 群公告
	GroupAttrKeyForbidden          = "forbidden"                   // 群禁言
	GroupAttrKeyInvite             = "invite"                      // 要请确认
	GroupAttrKeyForbiddenAddFriend = "forbidden_add_friend"        // 群内禁止加好友
	GroupAttrKeyStatus             = "status"                      // 群状态
	GroupAllowViewHistoryMsg       = "allow_view_history_msg"      // 是否允许新成员查看历史消息
	GroupAllowMemberPinnedMessage  = "allow_member_pinned_message" // 是否允许成员置顶消息
)

// 命令消息
const (
	CMDChannelUpdate           = "channelUpdate"          // 频道信息更新
	CMDGroupMemberUpdate       = "memberUpdate"           // 群成员更新
	CMDConversationUnreadClear = "unreadClear"            // 未读数清空
	CMDGroupAvatarUpdate       = "groupAvatarUpdate"      // 群头像更新
	CMDCommunityAvatarUpdate   = "communityAvatarUpdate"  // 社区头像更新
	CMDCommunityCoverUpdate    = "communityCoverUpdate"   // 社区封面更新
	CMDConversationDelete      = "conversationDelete"     // 删除最近会话
	CMDFriendRequest           = "friendRequest"          // 好友申请
	CMDFriendAccept            = "friendAccept"           // 接受好友申请
	CMDFriendDeleted           = "friendDeleted"          // 好友被删除
	CMDUserAvatarUpdate        = "userAvatarUpdate"       // 个人头像更新
	CMDTyping                  = "typing"                 // 输入中
	CMDOnlineStatus            = "onlineStatus"           // 在线状态
	CMDMomentMsg               = "momentMsg"              // 动态点赞或评论消息
	CMDSyncMessageExtra        = "syncMessageExtra"       // 同步消息扩展数据
	CMDSyncMessageReaction     = "syncMessageReaction"    // 同步消息回应数据
	CMDPCQuit                  = "pcQuit"                 // 退出pc登录
	CMDConversationDeleted     = "conversationDeleted"    // 最近会话被删除
	CMDSyncReminders           = "syncReminders"          // 同步提醒项
	CMDSyncConversationExtra   = "syncConversationExtra"  // 同步最近会话扩展
	CMDOrganizationInfoUpdate  = "organizationInfoUpdate" // 组织信息更新
	CMDQuitOrganization        = "quitOrganization"       // 退出某个组织
	CMDJoinOrganization        = "joinOrganization"       // 加入某个组织
	CMDSyncPinnedMessage       = "syncPinnedMessage"      // 同步置顶消息
	CMDMessageErase            = "messageErase"           // 消息擦除
)

const UserDeviceTokenPrefix = "userDeviceToken:" // 用户设备token缓存前缀

const UserDeviceBadgePrefix = "userDeviceBadge" // 用户设备红点

const QRCodeCachePrefix = "qrcode:" // 二维码缓存前缀

const AuthCodeCachePrefix = "authcode:" // 授权code

const DeviceCacheUUIDPrefix = "deviceCacheUUID:" // 设备UUID前缀

type AuthCodeType string // 认证代码类型

const (
	AuthCodeTypeJoinGroup         AuthCodeType = "joinGroup"         // 进群授权code
	AuthCodeTypeGroupMemberInvite AuthCodeType = "groupMemberInvite" // 群成员邀请
	AuthCodeTypeScanLogin         AuthCodeType = "scanLogin"         // 扫描登录
)

type DeviceType string // 设备类型

const (
	DeviceTypeIOS      DeviceType = "IOS"      // iOS设备
	DeviceTypeMI       DeviceType = "MI"       // 小米设备
	DeviceTypeHMS      DeviceType = "HMS"      // 华为设备
	DeviceTypeFirebase DeviceType = "FIREBASE" // 海外GOOGLE Firebase推送
	DeviceTypeOPPO     DeviceType = "OPPO"     // oppo设备
	DeviceTypeVIVO     DeviceType = "VIVO"     // vivo设备
)

type QRCodeType string // 二维码类型

const (
	QRCodeTypeGroup     QRCodeType = "group"     // 群聊
	QRCodeTypeScanLogin QRCodeType = "scanLogin" // 扫描登录
)

type ScanLoginStatus string // 扫码状态

const (
	ScanLoginStatusExpired  ScanLoginStatus = "expired"  // 二维码过期
	ScanLoginStatusWaitScan ScanLoginStatus = "waitScan" // 等待扫码
	ScanLoginStatusScanned  ScanLoginStatus = "scanned"  // 已扫描
	ScanLoginStatusAuthed   ScanLoginStatus = "authed"   // 已授权
)

type VercodeType int // 加好友验证码类型

const (
	User           VercodeType = 1 // 搜索
	GroupMember    VercodeType = 2 // 群成员
	QRCode         VercodeType = 3 // 二维码
	Friend         VercodeType = 4 // 好友
	MailList       VercodeType = 5 // 手机联系人
	InvitationCode VercodeType = 6 // 邀请码
)

type QRCodeModel struct {
	Type QRCodeType             `json:"type"` // 二维码类型
	Data map[string]interface{} `json:"data"`
}

type UserStatus int // 用户状态

const (
	UserAvailable UserStatus = 1 // 可用
	UserDisable   UserStatus = 0 // 禁用
)

func NewQRCodeModel(typ QRCodeType, data map[string]interface{}) *QRCodeModel {
	return &QRCodeModel{
		Type: typ,
		Data: data,
	}
}

func (q QRCodeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(string(q))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (q *QRCodeType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*q = QRCodeType(j)
	return nil
}

type RTCCallType int

const (
	RTCCallTypeAudio RTCCallType = 0 // 语音通话
	RTCCallTypeVideo RTCCallType = 1 // 视频通话
)

type RTCResultType int

const (
	RTCResultTypeCancel RTCResultType = 0 // 取消通话
	RTCResultTypeHangup RTCResultType = 1 // 挂断通话
	RTCResultTypeMissed RTCResultType = 2 // 未接听
	RTCResultTypeRefuse RTCResultType = 3 // 拒绝接听
)
