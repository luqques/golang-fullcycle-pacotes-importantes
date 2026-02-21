package main

import (
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
