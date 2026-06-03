package main

import (
	"log"
	"net/http"
	"time"
)

func openServer() {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        FunctionHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type FunctionHandler struct{}

func (f FunctionHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if !ValidServer(req) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	//metodo
}

func ValidServer(req *http.Request) bool {
	return true
}
