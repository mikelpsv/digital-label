package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	app "github.com/mlplabs/app-utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wib-project/conf"
	"wib-project/routes"
)

var StopAllTasks chan bool

func main() {
	StopAllTasks = make(chan bool)

	conf.ReadEnv()

	app.Log.Init("", "")

	app.InitDb(conf.Cfg.DbHost, conf.Cfg.DbName, conf.Cfg.DbUser, conf.Cfg.DbPassword)

	wHandlers := routes.NewWrapHandlers()
	wHandlers.Log = routes.WrapHttpLog{
		Trace:   app.Log.Trace,
		Info:    app.Log.Info,
		Warning: app.Log.Warning,
		Error:   app.Log.Error,
	}

	routeItems := app.Routes{}
	routeItems = RegisterHandlers(routeItems, wHandlers)

	router := NewRouter(routeItems)
	router.NotFoundHandler = http.HandlerFunc(wHandlers.Custom404)

	StartHttpServer(conf.Cfg.AppAddr, conf.Cfg.AppPort, router, CloseAll)
}

func RegisterHandlers(routeItems app.Routes, wh *routes.WrapHttpHandlers) app.Routes {
	routeItems = routes.RegisterControlHandlers(routeItems, wh)
	routeItems = routes.RegisterUtilsHandlers(routeItems, wh)
	routeItems = routes.RegisterDataHandlers(routeItems, wh)

	return routeItems
}

func NewRouter(routeItems app.Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeItems {
		handlerFunc := route.HandlerFunc
		if route.ValidateToken {
			handlerFunc = routes.SetMiddlewareApiKey(handlerFunc)
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

func StartHttpServer(Ip string, Port string, Router *mux.Router, deferFunc func()) {
	app.Log.Info.Printf("Starting Api server on %s:%s", Ip, Port)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", Ip, Port),
		Handler: Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			app.Log.Error.Fatal(err)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigterm

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Log.Error.Fatalf("server shutdown failed:%+v", err)
	}
	if deferFunc != nil {
		deferFunc()
	}

	app.Log.Info.Printf("Stopped Api server on %s port", Port)
}

func CloseAll() {
	if err := app.Db.Close(); err != nil {
		app.Log.Error.Fatalf("database close failed:%+v", err)
	}
	app.Log.Info.Println("Bye!")
}
