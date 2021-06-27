package main

import (
	"log"

	"github.com/larkintuckerllc/caller/internal/caller"
)

func main() {
	err := caller.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
