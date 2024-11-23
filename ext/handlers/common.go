package handlers

import (
	"github.com/hasnainkk-07/gotgbot/v2"
	"github.com/hasnainkk-07/gotgbot/v2/ext"
)

type Response func(b *gotgbot.Bot, ctx *ext.Context) error
