package common

import (
	"fmt"
	"strings"

	imlog "github.com/mushanyux/MSChatServerLib/pkg/log"
	"github.com/mushanyux/MSChatServerLib/pkg/util"
	"go.uber.org/zap"
)

// ContentType 正文类型
type ContentType int

const (
	// 聊天类
	Text     ContentType = 1 // 文本消息
	Image    ContentType = 2 // 图片消息
	GIF      ContentType = 3 // GIF消息
	Voice    ContentType = 4 // 语音消息
	Video    ContentType = 5 // 视频消息
	Location ContentType = 6 // 位置
	Card     ContentType = 7 // 名片
	File     ContentType = 8 // 文件

	MultipleForward ContentType = 11 //	合并转发
	VectorSticker   ContentType = 12 // 矢量表情
	EmojiSticker    ContentType = 13 // 矢量emoji表情
	RichText        ContentType = 14 // 富文本消息

	InviteJoinOrganization ContentType = 16 // 邀请加入组织

	ContentError ContentType = 97 // 消息正文
	SignalError  ContentType = 98 // 解密失败
	CMD          ContentType = 99 // CMD消息

	// 系统类
	Tip ContentType = 2000 // 只作为提醒

	FriendApply          ContentType = 1000 // 好友申请
	GroupCreate          ContentType = 1001 // 群创建
	GroupMemberAdd       ContentType = 1002 // 群成员添加
	GroupMemberRemove    ContentType = 1003 // 群成员移除
	FriendSure           ContentType = 1004 // 好友确认
	GroupUpdate          ContentType = 1005 // 群更新
	RevokeMessage        ContentType = 1006 // 撤回消息
	GroupMemberScanJoin  ContentType = 1007 // 扫码进群
	GroupTransferGrouper ContentType = 1008 // 转让群主
	GroupMemberInvite    ContentType = 1009 // 群成员邀请
	GroupMemberBeRemove  ContentType = 1010 // 群成员被移除
	GroupMemberQuit      ContentType = 1011 // 群成员主动退群
	GroupUpgrade         ContentType = 1012 // 群升级

	// 客服类
	HotlineAssignTo ContentType = 1200 // 分配客服
	HotlineSolved   ContentType = 1201 // 已解决
	HotlineReopen   ContentType = 1202 // 会话重开

	// 音视频
	VideoCallResult ContentType = 9989 // 音视频通话
)

func GetDisplayText(contentType int) string {
	if contentType == Text.Int() {
		return "文本消息"
	} else if contentType == Image.Int() {
		return "图片消息"
	} else if contentType == GIF.Int() {
		return "GIF"
	} else if contentType == Voice.Int() {
		return "语音"
	} else if contentType == Video.Int() {
		return "视频"
	} else if contentType == Location.Int() {
		return "位置"
	} else if contentType == Card.Int() {
		return "名片"
	} else if contentType == File.Int() {
		return "文件"
	} else if contentType == MultipleForward.Int() {
		return "合并转发消息"
	} else if contentType == VectorSticker.Int() {
		return "贴纸"
	} else if contentType == EmojiSticker.Int() {
		return "emoji"
	} else if contentType == RichText.Int() {
		return "富文本消息"
	}
	return "未知消息类型"
}

func (c ContentType) String() string {
	switch c {
	case Text:
		return "Text"
	case Image:
		return "Image"
	case GIF:
		return "GIF"
	case Voice:
		return "Voice"
	case CMD:
		return "CMD"
	case FriendApply:
		return "FriendApply"
	case GroupCreate:
		return "GroupCreate"
	case GroupMemberAdd:
		return "GroupMemberAdd"
	case GroupMemberRemove:
		return "GroupMemberRemove"
	case FriendSure:
		return "FriendSure"
	case GroupUpdate:
		return "GroupUpdate"
	case RevokeMessage:
		return "RevokeMessage"
	}
	return fmt.Sprintf("%d", c)
}

func (c ContentType) Int() int { return int(c) }

func GetFakeChannelIDWith(fromUID, toUID string) string {
	fromUIDHash := util.HashCrc32(fromUID)
	toUIDHash := util.HashCrc32(toUID)
	if fromUIDHash > toUIDHash {
		return fmt.Sprintf("%s@%s", fromUID, toUID)
	}
	if fromUIDHash == toUIDHash {
		imlog.Warn("生成的fromUID和toUID的hash是相同的！", zap.Uint32("fromUIDHash", fromUIDHash), zap.Uint32("toUIDHash", toUIDHash), zap.String("fromUID", fromUID), zap.String("toUID", toUID))
	}
	return fmt.Sprintf("%s@%s", toUID, fromUID)
}

func IsFakeChannel(channelID string) bool {
	return strings.Contains(channelID, "@")
}

func GetToChannelIDWithFakeChannelID(fakeChannelID string, uid string) string {
	channelIDs := strings.Split(fakeChannelID, "@")
	toChannelID := fakeChannelID
	if len(channelIDs) == 2 {
		if channelIDs[0] == uid {
			toChannelID = channelIDs[1]
		} else {
			toChannelID = channelIDs[0]
		}
	}
	return toChannelID
}
