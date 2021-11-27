package main

import (
	"fmt"
	"io"

	"github.com/nna774/times-carshare-point-program/http"
)

func main() {
	resp, err := http.Get("http://www.kmc.gr.jp/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	r, _ := io.ReadAll(resp.Body)
	fmt.Printf("status: %v\n", resp.StatusCode)
	fmt.Printf("headers: %v\n", resp.Header)
	fmt.Printf("body: %v\n", string(r))
}
