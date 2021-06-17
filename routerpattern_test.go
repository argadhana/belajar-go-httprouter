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

func TestRouterParamNamed(t *testing.T) {
	router := httprouter.New()
	router.GET("/profil/:id/users/:detail", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		id := p.ByName("id")
		detail := p.ByName("detail")
		fmt.Fprintf(w, "Profil id : " + id + " detail : " + detail)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/profil/2/users/cv", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	ressult, _ := io.ReadAll(response.Body)

	strress := string(ressult)

	assert.Equal(t, "Profil id : 2 detail : cv", strress)
}

func TestRouterCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/profil/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		profilImage := p.ByName("image")
		fmt.Fprintf(w, "Profil image : " + profilImage )
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/profil/images/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	ressult, _ := io.ReadAll(response.Body)

	strress := string(ressult)

	assert.Equal(t, "Profil image : /images/profile.png", strress)
}
