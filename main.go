package main

import (
	"log"
)

func main() {
	log.Print("Start Application")

	router := FiberRouter().InitRouter()
	router.Listen(8080)
}
