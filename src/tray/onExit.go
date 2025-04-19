package tray

import (
	"log/slog"
	"os"
)

func onExit() {
	slog.Info("Exiting...")
	os.Exit(0)
}
