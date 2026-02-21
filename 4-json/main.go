package main

import "encoding/json"

type Conta struct {
	Numero int
	Saldo  float64
}

func main() {
	conta := Conta{Numero: 12345, Saldo: 1000.50}
	json, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(json))
}
