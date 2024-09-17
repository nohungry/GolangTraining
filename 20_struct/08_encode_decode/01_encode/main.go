package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	tmp := json.NewEncoder(os.Stdout).Encode(p1)
	fmt.Println(tmp)
}
