package main

import (
	"go-html-monitor/src/service"
	"go-html-monitor/src/tray"
	"go-html-monitor/src/util"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/phsym/console-slog"
)

const cacheFile = "gen/sent_ids.json"

func main() {
	util.LoadCache(cacheFile)
	defer util.SaveCache(cacheFile)

	level := slog.LevelInfo
	if os.Getenv("DEBUG") == "1" {
		level = slog.LevelDebug
	}

	handler := console.NewHandler(os.Stdout, &console.HandlerOptions{
		Level:     level,
		AddSource: true,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Debug("Debug mode ON")
	slog.Info("Starting App...")

	go tray.Init()
	service.Init(
		"https://www.linkedin.com/jobs/search/?geoId=104370960&keywords=(%22golang%22%20OR%20%22go%22)%20AND%20(%22backend%22%20OR%20%22engineer%22)",
		"linkedInJob",
	)
}
