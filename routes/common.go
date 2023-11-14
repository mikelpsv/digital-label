package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WrapHttpLog struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

type WrapHttpHandlers struct {
	Log WrapHttpLog
}

func NewWrapHandlers() *WrapHttpHandlers {
	wHandlers := new(WrapHttpHandlers)
	return wHandlers
}

type CounterWriter struct {
	http.ResponseWriter
	Count int
}

func (cw *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = cw.ResponseWriter.Write(p)
	cw.Count += n
	return
}

func ResponseJSONLocal(w http.ResponseWriter, statusCode int, data interface{}) int {
	cw := &CounterWriter{
		ResponseWriter: w,
	}

	cw.WriteHeader(statusCode)
	err := json.NewEncoder(cw).Encode(data)
	if err != nil {
		return 0
		fmt.Fprintf(cw, "%s", err.Error())
	}
	return cw.Count
}

// SetMiddlewareApiKey проверяет api key, определяет клиента, и пробрасывает дальше
func SetMiddlewareApiKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//apiKey := r.URL.Query().Get("api_key")
		//if apiKey == "" {
		//	app.ResponseERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		//	return
		//}
		//
		//client, err := new(models.Client).FindByApiKey(apiKey)
		//if err != nil {
		//	app.ResponseERROR(w, http.StatusForbidden, errors.New("unauthorized"))
		//	return
		//}
		//if !client.IsActive() {
		//	app.ResponseERROR(w, http.StatusForbidden, errors.New("unauthorized"))
		//	return
		//}
		//
		//ctx := context.WithValue(r.Context(), "client", client)
		//
		//next(w, r.WithContext(ctx))
		next(w, r)
	}
}
