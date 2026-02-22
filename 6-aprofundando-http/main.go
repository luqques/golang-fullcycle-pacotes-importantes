package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	viaCep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(viaCep) // forma mais curta somente para retornar o valor direto sem guardar o resultado parseado em uma variável.

	// responseData, err := json.Marshal(viaCep) // forma mais longa onde o resultado do parseamento é guardado em uma variável.
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(responseData)
}

func BuscaCep(cep string) (*ViaCep, error) {
	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		println("Erro ao buscar CEP:", err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		println("CEP não encontrado:", response.Status)
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		println("Erro ao ler resposta:", err)
		return nil, err
	}

	var viaCep ViaCep
	err = json.Unmarshal(responseData, &viaCep)
	if err != nil {
		println("Erro ao fazer unmarshal:", err)
		return nil, err
	}

	return &viaCep, nil
}
