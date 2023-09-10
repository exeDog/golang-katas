package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	const URL = "https://www.google.com"

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	_, err = resp.Body.Read(bs)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(bs))
}
