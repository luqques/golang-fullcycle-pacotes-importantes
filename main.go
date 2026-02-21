package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Bytes escritos:", tamanho)

	f.Close()

	//leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	//leitura com buffer
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("Erro ao acabar o arquivo:", err)
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	arquivo2.Close()

	//removendo arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
