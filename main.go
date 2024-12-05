package main

import (
	"Nextlaunch/src/config"
	"Nextlaunch/src/logging"
	"Nextlaunch/src/tsd"
	"Nextlaunch/src/tui"
	"Nextlaunch/src/tui/widgets"
)

var logger *logging.Logger

var ll2 tsd.LL2Client
var snapi tsd.SnapiClient

func init() {
	logger = logging.NewLogger("main")
	logger.Log("Starting")
	//logging.EnterTui()

	config.LoadConfig()
}

func main() {

	logger.Log("Waiting for initialization to finish")

	logger.Log("Initialized")

	logger.Log("Preparing caches")

	ll2 = *tsd.NewLL2Client()
	snapi = *tsd.NewSnapiClient()

	logger.Log("Starting application")

	config.LoadConfig()

	context := tui.Model{
		KeybindingManager: tui.NewKeybindManager(),
		CursorPosition:    tui.CursorPosition{0, 0},
		CursorBlink:       false,
		CursorVisible:     false,
		CursorStyle:       tui.CursorStyleNone,
		Data:              make(map[string]interface{}),
		LL2:               ll2,
		Snapi:             snapi,
		Page:              0,
		Frame:             widgets.NewWindow("NextLaunch", 10, 10, false),
	}

	tui.StartBubbletea(&context)
	logging.ShouldExit = true
}
