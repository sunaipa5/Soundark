package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Soundark",
		Width:  1150,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       app.ondomready,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(0, 0, 0),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				LightModeTitleBar:  windows.RGB(0, 0, 0),
				LightModeTitleText: windows.RGB(200, 200, 200),
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
