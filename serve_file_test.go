package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestServeFileHello(t *testing.T) {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("resources"))

	request := httptest.NewRequest(http.MethodGet, "/static/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	results := recorder.Result()
	body, _ := io.ReadAll(results.Body)

	assert.Equal(t, "hello world\n", string(body))
}

func TestServeFileBye(t *testing.T) {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("resources"))

	request := httptest.NewRequest(http.MethodGet, "/static/bye/goodbye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	results := recorder.Result()
	body, _ := io.ReadAll(results.Body)

	assert.Equal(t, "bye world\n", string(body))
}
