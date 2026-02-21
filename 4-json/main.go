package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int     `json:"numero_conta"`
	Saldo  float64 `json:"saldo_conta"`
}

func main() {
	conta := Conta{Numero: 12345, Saldo: 1000.50}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero_conta":67890,"saldo_conta":2500.75}`)
	var outraConta Conta
	err = json.Unmarshal(jsonPuro, &outraConta)
	if err != nil {
		panic(err)
	}
	println(outraConta.Numero, outraConta.Saldo)
}
