package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	const URL = "https://www.google.com"

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	_, err2 := io.Copy(os.Stdout, resp.Body)

	if err2 != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	_, err3 := io.Copy(logWriter{}, resp.Body)

	if err3 != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes:", len(bs), "bytes")
	return len(bs), nil
}
