package routes

import (
	"github.com/gorilla/mux"
	"github.com/mikelpsv/digital-label/pkg/model"
	app "github.com/mlplabs/app-utils"
	"net/http"
	"strconv"
)

func RegisterUtilsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
	routeItems = append(routeItems, app.Route{
		Name:          "Encode",
		Method:        "GET",
		Pattern:       "/utils/encode/{num}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.Encode,
	})
	routeItems = append(routeItems, app.Route{
		Name:          "Encode",
		Method:        "GET",
		Pattern:       "/utils/decode/{str}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wHandlers.Decode,
	})
	return routeItems
}

func (wh *WrapHttpHandlers) Encode(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	numStr := p["num"]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		//
	}
	str := model.NewEnc62("").Encode(uint64(num))
	app.ResponseJSON(w, http.StatusOK, str)
}
func (wh *WrapHttpHandlers) Decode(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	str := p["str"]
	num := model.NewEnc62("").Decode(str)
	app.ResponseJSON(w, http.StatusOK, num)
}
