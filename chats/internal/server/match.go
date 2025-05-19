package server

import (
	"github.com/go-telegram/bot/models"
)

func messageTextMatch(update *models.Update) bool {
	if update.Message != nil {
		return update.Message.Text != ""
	}
	return false
}

func memberKickedMatch(update *models.Update) bool {
	if update.MyChatMember != nil {
		return (
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeBanned ||
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeLeft)
	}
	return false
}