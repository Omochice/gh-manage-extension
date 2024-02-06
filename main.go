package main

import (
	"log"

	"github.com/Omochice/gh-manage-extension/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
