package service

import (
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
		Name:          "Encode",
		Method:        "GET",
		Pattern:       "/l/{key_link}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Empty(c.GetData),
	})
	return routeItems
}

func (c *Controller) GetData(r *http.Request) error {
	return nil
}
