package service

import "strings"

func isRecent(text string) bool {
	if strings.Contains(text, "day") {
		return strings.Contains(text, "1 day") || strings.Contains(text, "0 day")
	}
	if strings.Contains(text, "hour") || strings.Contains(text, "minute") || strings.Contains(text, "second") {
		return true
	}
	return false
}
