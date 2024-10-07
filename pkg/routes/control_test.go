package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	type PingResponseTest struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}
	wh := GetWrapHandlers()
	r := httptest.NewRequest("GET", "http://127.0.0.1", nil)
	w := httptest.NewRecorder()
	wh.Ping(w, r)
	data := PingResponseTest{}
	err := json.Unmarshal(w.Body.Bytes(), &data)
	if err != nil {
		t.Error(err)
	}
	if data.Code != http.StatusOK {
		t.Errorf("Invalid")
	}
}
