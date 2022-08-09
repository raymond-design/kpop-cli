package main

import (
	"cli-tunes/api"
	"fmt"
	"os"
	"os/signal"
)

const JPOP string = "https://listen.moe/fallback"
const KPOP string = "https://listen.moe/kpop/fallback"

const JSOCKET string = "wss://listen.moe/gateway_v2"
const KSOCKET string = "wss://listen.moe/kpop/gateway_v2"

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	mode := "kpop"
	var stream string
	var socket string

	if len(os.Args) == 2 {
		mode = os.Args[1]
	}

	switch mode {
	case "kpop":
		stream = KPOP
		socket = KSOCKET
	case "jpop":
		stream = JPOP
		socket = JSOCKET
	default:
		fmt.Println("Error")
		os.Exit(1)
	}

	api.Play(socket)
	api.Play(stream)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	fmt.Println("Exiting Player")
	api.Stop()
	api.Stop()
}
