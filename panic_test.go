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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(rw http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprint(rw, "Panic: ", i)
	}
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Upps")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	results := recorder.Result()
	body, _ := io.ReadAll(results.Body)

	assert.Equal(t, "Panic: Upps", string(body))
}
