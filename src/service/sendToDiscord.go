package service

import (
	"bytes"
	"encoding/json"
	"go-html-monitor/src/model"
	"io"
	"log/slog"
	"net/http"
)

func sendToDiscord(webhook string, payload model.DiscordPayload) {
	if len(payload.Embeds) == 0 {
		slog.Warn("Skip, embeds is empty.")
		return
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(body))
	if err != nil {
		slog.Error("❌ Gagal kirim ke Discord", "err", err)
		return
	}
	defer resp.Body.Close()
	slog.Info("✅ Terkirim ke Discord", "status", resp.Status)

	if resp.StatusCode == http.StatusBadRequest {
		respBodyBadRequest, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		slog.Warn("Response Bad Request", "body", string(respBodyBadRequest))
	}
}
