package tray

import (
	"log/slog"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	appName := "Go HTML Monitor App"
	systray.SetIcon(icon.Data)

	systray.SetTitle(appName)
	systray.SetTooltip(appName)

	quitItem := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-quitItem.ClickedCh
		slog.Info("Quit clicked, shutting down...")
		systray.Quit()
	}()
}
