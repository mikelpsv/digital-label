package routes

import (
	app "github.com/mlplabs/app-utils"
	"net/http"
	"time"
)

type PingResponse struct {
	Status      string `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"-"`
	Version     string `json:"version"`
}

type HealthResponse struct {
	Database bool `json:"database"`
}

func RegisterControlHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
	routeItems = append(routeItems, app.Route{
		Name:          "Ping",
		Method:        "GET",
		Pattern:       "/{version}/ping",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.Ping,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "health",
		Method:        "GET",
		Pattern:       "/{version}/health",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetHealth,
	})

	return routeItems
}

func (wh *WrapHttpHandlers) Ping(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	resp := PingResponse{
		Status:      "ок",
		Code:        http.StatusOK,
		Description: "",
		Version:     "v1.0",
	}
	bytes := ResponseJSONLocal(w, http.StatusOK, resp)
	wh.Log.Info.Printf("HTTP PING: %d %f", bytes, time.Since(t).Seconds())
}

func (wh *WrapHttpHandlers) GetHealth(w http.ResponseWriter, r *http.Request) {
	pingErr := app.Db.Ping()
	app.ResponseJSON(w, http.StatusOK, HealthResponse{
		Database: pingErr == nil,
	})
}

func (wh *WrapHttpHandlers) Custom404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	wh.Log.Error.Printf("HTTP INVALID ROUTE: %s", r.RequestURI)
}
