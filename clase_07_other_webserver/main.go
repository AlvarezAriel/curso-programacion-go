package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)
import "encoding/json"

type Pokemon struct {
	Name string
	Order int
	Sprites Images
}

type Images struct {
	FrontDefault string `json:"front_default"`
}

func main() {
	http.HandleFunc("/", servePage)

	http.ListenAndServe(":8080", nil)
}

func servePage(writer http.ResponseWriter, request *http.Request) {
	duck := "https://pokeapi.co/api/v2/pokemon/" + strings.ReplaceAll(request.URL.Path, "/", "")

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

	pokemon := Pokemon{}
	json.Unmarshal(body, &pokemon)


	htmlBody := `<!DOCTYPE html>
	<html lang="en">
		<head>
			<!-- Required meta tags -->
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
			<title>Golang HTML Server</title>
		</head>
		<body>
			<div class="container">
 				<img src="`+ pokemon.Sprites.FrontDefault+`" alt="Girl in a jacket" width="500" height="600"/>
			</div>
		</body>
	</html>
	`

	io.WriteString(writer,htmlBody)
}
