package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var Method = map[string]map[string]func(w http.ResponseWriter, r *http.Request) {
	"GET":
}

var MethodFunction = map[string]func(w http.ResponseWriter, r *http.Request) {
	"/result": methodFunction,
}

func MethoResult(res http.ResponseWriter, req *http.Request) {
	operation := req.URL.Query().Get("op")
	sliceOperation := strings.Split(operation, " ")

	if len(sliceOperation) < 3 {
		//error
		return
	}

	op := sliceOperation[1]
	Searchable, ok := OperationsTable[op]

	if !ok {
		//Printer err
		PrintErro(res, operation)
		return
	}

	n1, err1 := strconv.ParseFloat(sliceOperation[0], 64)

	if err1 != nil {
		PrintErro(res, "Falha do parse do primeiro número")
		return
	}

	n2, err2 := strconv.ParseFloat(sliceOperation[2], 64)

	if err2 != nil {
		PrintErro(res, "Falha do parse do primeiro número")
		return
	}

	r := Operation{
		Op: op,
		Result: Searchable.ShowAnswers(n1, n2),
	}

	res.Write(ChangeJson(r))
	
}

func PrintErro(res http.ResponseWriter, op string) {
	err := map[string]string{"result" : "invalid expression", "op": op}
	json, _ := json.Marshal(err)
	res.Write(json)
	return
}

func ChangeJson (s Operation) []byte {
	json, _ := json.Marshal(s)
	return json
}