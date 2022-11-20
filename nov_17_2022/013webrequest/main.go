package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("jit web request")

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response of the type : %T", resp)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	content := string(body)
	fmt.Println(content)
}
