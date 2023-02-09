package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/igsr5/proglog/internal/server"
)

func TestHandleConsume(t *testing.T) {
	reqBodyJson, err := json.Marshal(&server.ConsumeRequest{Offset: 0})
	if err != nil {
		t.Fatal(err)
	}
	reqBody := bytes.NewBuffer(reqBodyJson)

	req := httptest.NewRequest("GET", "/", reqBody)
	res := httptest.NewRecorder()
	srv := server.NewHTTPServer(":8080")

	srv.Handler.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("want %d, but %d", http.StatusOK, res.Code)
	}
}
