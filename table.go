package main

//Classe que guarda a operação e o resultado

type Operation struct {
	Result float64 `json:"result"`
	Op     string  `json:"op"`
}

//Tabela de operações

var OperationsTable = map[string]Searchable{
	"+": Sum{},
	"-": Sub{},
	"*": Mul{},
	"/": Div{},
	"^": Pow{},
}
