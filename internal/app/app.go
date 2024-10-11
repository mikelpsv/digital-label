package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikelpsv/digital-label/pkg/config"
	"github.com/mikelpsv/digital-label/pkg/repositories"
	"github.com/mikelpsv/digital-label/pkg/routes"
	"github.com/mikelpsv/digital-label/pkg/routes/control"
	"github.com/mikelpsv/digital-label/pkg/routes/service"
	"github.com/mikelpsv/digital-label/pkg/routes/service_utils"
	"github.com/mikelpsv/digital-label/pkg/usecase"
	utils "github.com/mlplabs/app-utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init(cfg *config.Service) {
	ctx, cancel := context.WithCancel(context.Background())
	utils.Log.Init("", "")
	//utils.InitDb(cfg.DbHost, cfg.DbName, cfg.DbUser, cfg.DbPassword)
	//defer utils.Db.Close()

	repository := repositories.NewServiceRepository(utils.Db)
	thisService := usecase.NewService(repository)
	ctrlBase := service.NewController(thisService)
	ctrlUtils := service_utils.NewController(thisService)
	ctrlSystem := control.NewController(thisService)

	routeItems := utils.Routes{}
	routeItems = ctrlBase.RegisterHandlers(routeItems)
	routeItems = ctrlUtils.RegisterHandlers(routeItems)
	routeItems = ctrlSystem.RegisterHandlers(routeItems)

	wHandlers := routes.NewWrapHandlers()
	wHandlers.Log = routes.WrapHttpLog{
		Trace:   utils.Log.Trace,
		Info:    utils.Log.Info,
		Warning: utils.Log.Warning,
		Error:   utils.Log.Error,
	}

	router := NewRouter(routeItems)
	//router.NotFoundHandler = http.HandlerFunc(wHandlers.Custom404)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.AppAddr, cfg.AppPort),
		Handler: router,
	}
	fmt.Printf("Service started. Listen %s:%s", cfg.AppAddr, cfg.AppPort)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			utils.Log.Error.Fatal(err)
		}
	}()
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigterm

	cancel()

	if err := server.Shutdown(ctx); err != nil {
		utils.Log.Error.Fatalf("server shutdown failed:%+v", err)
	}

	fmt.Println("Service stopped")
	time.Sleep(5 * time.Second)
}

func NewRouter(routeItems utils.Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routeItems {
		handlerFunc := route.HandlerFunc
		if route.ValidateToken {
			handlerFunc = routes.SetMiddlewareApiKey(handlerFunc)
		}

		if route.SetHeaderJSON {
			handlerFunc = utils.SetMiddlewareJSON(handlerFunc)
		}
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handlerFunc)
	}
	return router
}
