package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikelpsv/digital-label/pkg/config"
	"github.com/mikelpsv/digital-label/pkg/model"
	routes2 "github.com/mikelpsv/digital-label/pkg/routes"
	app "github.com/mlplabs/app-utils"
	"github.com/segmentio/kafka-go"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app.Log.Init("", "")
	config.ReadEnv()

	app.InitDb(config.Cfg.DbHost, config.Cfg.DbName, config.Cfg.DbUser, config.Cfg.DbPassword)
	defer app.Db.Close()

	wHandlers := routes2.NewWrapHandlers()
	wHandlers.Log = routes2.WrapHttpLog{
		Trace:   app.Log.Trace,
		Info:    app.Log.Info,
		Warning: app.Log.Warning,
		Error:   app.Log.Error,
	}

	routeItems := app.Routes{}
	routeItems = RegisterHandlers(routeItems, wHandlers)

	router := NewRouter(routeItems)
	router.NotFoundHandler = http.HandlerFunc(wHandlers.Custom404)

	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:         []string{config.Cfg.KafkaHost0},
			GroupID:         config.Cfg.KafkaDataGroup,
			Topic:           config.Cfg.KafkaDataTopic,
			MinBytes:        10e3, // 10KB
			MaxBytes:        10e6, // 10MB
			ReadLagInterval: 500 * time.Millisecond,
		})
		defer reader.Close()
		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				fmt.Println(err)
			}
			code, err := model.ConvertMessage(msg)
			err = code.Write()
			if err != nil {
				fmt.Println(err)
			}
		}

	}()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Cfg.AppAddr, config.Cfg.AppPort),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			app.Log.Error.Fatal(err)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigterm

	cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Log.Error.Fatalf("server shutdown failed:%+v", err)
	}

	time.Sleep(5 * time.Second)
}

func RegisterHandlers(routeItems app.Routes, wh *routes2.WrapHttpHandlers) app.Routes {
	routeItems = routes2.RegisterControlHandlers(routeItems, wh)
	routeItems = routes2.RegisterUtilsHandlers(routeItems, wh)
	routeItems = routes2.RegisterDataHandlers(routeItems, wh)

	return routeItems
}

func NewRouter(routeItems app.Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeItems {
		handlerFunc := route.HandlerFunc
		if route.ValidateToken {
			handlerFunc = routes2.SetMiddlewareApiKey(handlerFunc)
		}

		if route.SetHeaderJSON {
			handlerFunc = app.SetMiddlewareJSON(handlerFunc)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handlerFunc)
	}

	return router
}
