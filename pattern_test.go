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

func TestNamedParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:idProduct/items/:idItem", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		idProduct, idItem := p.ByName("idProduct"), p.ByName("idItem")
		text := "Product " + idProduct + " Item " + idItem
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest(http.MethodGet, "/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	results := recorder.Result()
	body, _ := io.ReadAll(results.Body)

	assert.Equal(t, "Product 1 Item 2", string(body))
}

func TestCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*images", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		images := p.ByName("images")
		fmt.Fprint(rw, "Path: "+images)
	})

	request := httptest.NewRequest(http.MethodGet, "/images/subdir/avatar.jpg", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	results := recorder.Result()
	body, _ := io.ReadAll(results.Body)

	assert.Equal(t, "Path: /subdir/avatar.jpg", string(body))
}
