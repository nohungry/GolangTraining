package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://github.com/GeekwiseAcademy")
	if err != nil {
		log.Fatal(err)
	}
	// page, err := ioutil.ReadAll(res.Body)
	page, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
