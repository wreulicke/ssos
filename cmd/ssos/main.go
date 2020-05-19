package main

import (
	"log"
)

func main() {
	if err := NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
