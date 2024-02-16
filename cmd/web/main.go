package main

import (
	"log"

	"github.com/erry-az/templ-exmpl/internal/app"
)

func main() {
	err := app.NewWeb()
	if err != nil {
		log.Println(err)
	}
}
