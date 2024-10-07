package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mikelpsv/digital-label/pkg/model"
	app "github.com/mlplabs/app-utils"
	"html/template"
	"net/http"
)

type ViewData struct {
	Owner     string `json:"owner"`
	Title     string `json:"title"`
	OrderNum  string `json:"order_num"`
	Client    string `json:"client"`
	Address   string `json:"address"`
	BoxNumber string `json:"box_number"`
	BoxOneOf  int    `json:"box_one_of"`
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
	scheme := r.URL.Query().Get("raw")
	p := mux.Vars(r)
	link := p["key_link"]
	if link == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Данные по запросу не найдены"))
		return
	}

	code := new(model.Code)
	err := code.Get(link)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Данные по запросу не найдены"))
		return
	}

	if scheme == "" {
		data := new(ViewData)
		err = json.Unmarshal([]byte(code.Payload), data)
		var tmpl, _ = template.ParseFiles("templates/type1.html")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ошибка отображения данных"))
			return
		}
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write([]byte(code.Payload))
	}
}
