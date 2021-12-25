package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Gak Boleh")
	})
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "GET")
	})

	request := httptest.NewRequest(http.MethodPost, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	results := recorder.Result()
	body, _ := io.ReadAll(results.Body)

	assert.Equal(t, "Gak Boleh", string(body))
}
