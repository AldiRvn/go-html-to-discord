package service

import (
	"os"
)

func Init(linkHtml, web string) {
	payload := getDiscordPayload(linkHtml, web)

	sendToDiscord(os.Getenv("DC_WEBHOOK"), payload)
}
