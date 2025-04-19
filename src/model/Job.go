package model

import (
	"fmt"
	"strings"
	"time"
)

type Job struct {
	Title     string // Judul pekerjaan
	Company   string // Nama perusahaan
	Location  string // (Optional) Lokasi
	Link      string // URL job
	PostedAgo string // Waktu posting (e.g., "2 hours ago")
}

func (j *Job) ToDiscordEmbed() DiscordEmbed {
	return DiscordEmbed{
		Title:       j.Title,
		URL:         j.Link,
		Description: fmt.Sprintf("🏢 %s\n🕒 %s", j.Company, j.PostedAgo),
		Color:       0x00ADD8, // warna biru Golang
		Footer: &EmbedFooter{
			Text: "Sumber: LinkedIn",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func Job_BuildDiscordPayload(listJob []Job) DiscordPayload {
	embeds := []DiscordEmbed{}
	for _, job := range listJob {
		embeds = append(embeds, job.ToDiscordEmbed())
	}

	return DiscordPayload{
		Content: "🚀 **Lowongan Golang Baru!**",
		Embeds:  embeds,
	}
}

func (j *Job) ExtractID(raw string) string {
	parts := strings.Split(raw, "-")
	if len(parts) == 0 {
		return raw
	}
	last := parts[len(parts)-1]
	return strings.TrimSpace(strings.Split(last, "?")[0])
}
