package control

import (
	"github.com/mikelpsv/digital-label/pkg/routes"
	"github.com/mikelpsv/digital-label/pkg/usecase"
	app "github.com/mlplabs/app-utils"
	resp "github.com/mlplabs/app-utils/pkg/http/response"
	"net/http"
)

type Controller struct {
	service *usecase.Service
}

func NewController(service *usecase.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) RegisterHandlers(routeItems app.Routes) app.Routes {
	wrap := resp.NewWrapper()
	routeItems = append(routeItems, app.Route{
		Name:          "Ping",
		Method:        "GET",
		Pattern:       "/ping",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Data(c.Ping),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "health",
		Method:        "GET",
		Pattern:       "/health",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Data(c.GetHealth),
	})

	return routeItems
}

func (c *Controller) Ping(r *http.Request) (interface{}, error) {
	ping := routes.PingResponse{
		Status:      "ок",
		Code:        http.StatusOK,
		Description: "",
		Version:     "v1.0",
	}
	return &ping, nil
}

func (c *Controller) GetHealth(r *http.Request) (interface{}, error) {
	//pingErr := c.service.Ping()
	return routes.HealthResponse{
		Database: true,
	}, nil
}

func (c *Controller) Custom404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	//wh.Log.Error.Printf("HTTP INVALID ROUTE: %s", r.RequestURI)
}
