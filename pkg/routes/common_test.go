package routes

import (
	app "github.com/mlplabs/app-utils"
)

// GetWrapHandlers общий метод инициализации "обертки методов"
func GetWrapHandlers() *WrapHttpHandlers {
	wh := NewWrapHandlers()
	app.Log.Init("", "")
	wh.Log = WrapHttpLog(app.Log)
	return wh
}
