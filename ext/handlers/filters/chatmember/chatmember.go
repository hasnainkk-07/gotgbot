package chatmember

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func All(_ *gotgbot.ChatMemberUpdated) bool {
	return true
}

func UserId(id int64) filters.ChatMember {
	return func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm.NewChatMember.GetUser().Id == id
	}
}

func FromUserId(id int64) filters.ChatMember {
	return func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm.From.Id == id
	}
}

func ChatId(id int64) filters.ChatMember {
	return func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm.Chat.Id == id
	}
}

func Private(cm *gotgbot.ChatMemberUpdated) bool {
	return cm.Chat.Type == "private"
}

func Group(cm *gotgbot.ChatMemberUpdated) bool {
	return cm.Chat.Type == "group"
}

func Supergroup(cm *gotgbot.ChatMemberUpdated) bool {
	return cm.Chat.Type == "supergroup"
}

func Channel(cm *gotgbot.ChatMemberUpdated) bool {
	return cm.Chat.Type == "channel"
}

func InviteLink(cm *gotgbot.ChatMemberUpdated) bool {
	return cm.InviteLink != nil
}

func NewStatus(status string) filters.ChatMember {
	return func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm.NewChatMember.GetStatus() == status
	}
}

func OldStatus(status string) filters.ChatMember {
	return func(cm *gotgbot.ChatMemberUpdated) bool {
		return cm.OldChatMember.GetStatus() == status
	}
}
