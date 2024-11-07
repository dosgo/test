package main

import (
	"fmt"
	"log"

	"github.com/wlynxg/anet"
)

func main() {

	interfaces, err := anet.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		log.Printf("i:%+v\r\n", i)
	}
	fmt.Scanln()
}
