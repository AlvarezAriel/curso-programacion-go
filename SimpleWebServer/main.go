package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", servePage)

	http.ListenAndServe(":8080", nil)
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	duck := "https://api.duckduckgo.com/?q=" +
		strings.ReplaceAll(request.URL.Path, "/", "") +
		"&format=json&pretty=1"
	fmt.Println(duck)

	resp, err := http.Get(duck)

	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	io.WriteString(writer, `{"hola":"soy ariel", "resultado":`+string(body)+`}`)
}
