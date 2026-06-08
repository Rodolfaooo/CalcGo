package main

type Operation struct {
	Result float64 `json:"result"`
	Op 		string `json:"op"`
}

var OperationsTable = map[string]Searchable {
	"+": Sum{},
	"-": Sub{},
	"*": Mul{},
	"/": Div{},
	"^": Pow{},
}