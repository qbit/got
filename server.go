package main

import (
	//"github.com/eaigner/shield"
	"fmt"
	"github.com/qbit/goirc"
	"log"
)

func errr(e error, msg string) {
	if e != nil {
		log.Printf("[!]: %s - %s", msg, e)
	}
}

func main() {
	client, err := goirc.NewIrc("config.json")

	if client == nil {
		errr(err, "Can't connect to IRC")
	}

	client.Connect()

	fmt.Println(client.Recv(0, true))
}
