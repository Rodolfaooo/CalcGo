package main

import (
	"net/http"
	"strings"
	"fmt"
)

var Method = map[string]map[string]func(w http.ResponseWriter, r *http.Request) {

}

func MethoResult(res http.ResponseWriter, req *http.Request) {
	operation := req.URL.Query().Get("op")
	sliceOperation := strings.Split(operation, " ")

	if len(sliceOperation) < 3 {
		//error
		return
	}

	op := sliceOperation[1]
	
}