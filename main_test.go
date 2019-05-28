package main

import (
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/",RootEndpoint).Methods("GET")
	return router
}

func TestEndPoint(t *testing.T) {
	request,_ := http.NewRequest("GET","/",nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t,200,response.Code,"Response is ok")
	//assert.Equal(t,"Hello World",response.Body.String(),"It is all right")
}