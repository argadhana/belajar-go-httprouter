package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer, "Panic : ", i)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		panic("ups")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	ressult, _ := io.ReadAll(response.Body)

	strress := string(ressult)

	assert.Equal(t, "Panic : ups", strress)
}
