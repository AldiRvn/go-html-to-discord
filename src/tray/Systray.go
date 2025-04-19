package tray

import "github.com/getlantern/systray"

func Init() {
	systray.Run(onReady, onExit)
}
