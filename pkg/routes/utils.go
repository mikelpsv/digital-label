package routes

//
//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"github.com/mikelpsv/digital-label/pkg/model"
//	app "github.com/mlplabs/app-utils"
//	"github.com/mlplabs/app-utils/pkg/http/errors"
//	resp "github.com/mlplabs/app-utils/pkg/http/response"
//	"net/http"
//	"strconv"
//)
//
//func RegisterUtilsHandlers(routeItems app.Routes, wHandlers *WrapHttpHandlers) app.Routes {
//	wrap := resp.NewWrapper()
//
//	routeItems = append(routeItems, app.Route{
//		Name:          "Encode",
//		Method:        "GET",
//		Pattern:       "/utils/encode/{num}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wrap.Data(wHandlers.Encode),
//	})
//	routeItems = append(routeItems, app.Route{
//		Name:          "Encode",
//		Method:        "GET",
//		Pattern:       "/utils/decode/{str}",
//		SetHeaderJSON: true,
//		ValidateToken: false,
//		HandlerFunc:   wrap.Data(wHandlers.Decode),
//	})
//	return routeItems
//}
//
//func (wh *WrapHttpHandlers) Encode(r *http.Request) (interface{}, error) {
//	p := mux.Vars(r)
//	numStr := p["num"]
//	num, err := strconv.Atoi(numStr)
//	if err != nil {
//		return "", errors.NewInvalidInputData(fmt.Errorf("error converting input data"))
//	}
//	return model.NewEnc62("").Encode(uint64(num)), nil
//}
//
//func (wh *WrapHttpHandlers) Decode(r *http.Request) (interface{}, error) {
//	p := mux.Vars(r)
//	str := p["str"]
//	num := model.NewEnc62("").Decode(str)
//	return num, nil
//}
