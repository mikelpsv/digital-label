package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"html/template"
	"net/http"
	"wib-project/model"
)

type ViewData struct {
	Title    string `json:"title"`
	OrderNum string `json:"order_num"`
	Client   string `json:"client"`
	Address  string `json:"address"`
}

func RegisterDataHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
	routeItems = append(routeItems, app.Route{
		Name:          "Encode",
		Method:        "GET",
		Pattern:       "/{key_link}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.GetData,
	})
	return routeItems
}

func (wh *WrapHttpHandlers) GetData(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	link := p["key_link"]
	if link == "" {
		// ошибку нарисовать
	}

	code := new(model.Code)
	err := code.Get(link)
	if err != nil {
		// ошибку нарисовать
	}
	data := new(ViewData)
	_ = json.Unmarshal([]byte(code.Payload), data)
	data.Title = "Просто текст"
	data.Client = "ООО Клиент"
	data.Address = "660077, Красноярский край, Красноярск г, Весны ул, дом № 11"
	var tmpl, _ = template.ParseFiles("templates/type1.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_ = tmpl.Execute(w, data)

}
