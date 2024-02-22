package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/src/SMS/docs
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sms2wails",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 0, B: 0, A: 0},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			BackdropType:         windows.Acrylic,
			// BackdropType: windows.Tabbed,
			// BackdropType: windows.Mica,
			EnableSwipeGestures: true,
		},
		// OnStartup:        app.startup,
		Bind: []interface{}{
			// app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
