package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikelpsv/digital-label/internal/usecase"
	"github.com/mikelpsv/digital-label/pkg/model"
	app "github.com/mlplabs/app-utils"
	"github.com/mlplabs/app-utils/pkg/http/errors"
	resp "github.com/mlplabs/app-utils/pkg/http/response"
	"html/template"
	"io"
	"log"
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
		Name:          "GetLinkData",
		Method:        "GET",
		Pattern:       "/l/{key_link}",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Raw(c.GetData),
	})
	routeItems = append(routeItems, app.Route{
		Name:          "WriteData",
		Method:        "POST",
		Pattern:       "/write",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   wrap.Data(c.WriteData),
	})
	return routeItems
}

func (c *Controller) GetData(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	p := mux.Vars(r)
	keyLink := p["key_link"]
	linkData, err := c.service.GetLink(keyLink)
	if err != nil {
		return "", errors.NewInvalidInputData(err)
	}
	if linkData.KeyLink == "" {
		return "", errors.NewNotFoundError(fmt.Errorf("сould not find information about the box"))
	}

	viewData := new(model.ViewData)
	err = json.Unmarshal([]byte(linkData.Payload), &viewData)
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/type%d.html", linkData.Type))
	if err != nil {
		log.Printf(err.Error())
		return "", errors.NewServerError(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, viewData)
	if err != nil {
		log.Printf(err.Error())
		return "", errors.NewServerError(err)
	}

	return buf.String(), nil
}

func (c *Controller) WriteData(r *http.Request) (interface{}, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("could not read body: %w", err))
	}
	linkData := model.LinkData{}
	err = json.Unmarshal(b, &linkData)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("could not unmarshal body: %w", err))
	}
	rawMsg := json.RawMessage{}
	err = json.Unmarshal([]byte(linkData.Payload), &rawMsg)
	if err != nil {
		return nil, errors.NewInvalidInputData(fmt.Errorf("could not unmarshal payload: %w", err))
	}
	err = c.service.WriteData(&linkData)
	if err != nil {
		return nil, errors.NewServerError(err)
	}
	return "", nil
}
