package main

import (
	"go-linha-comando/app"
	"log"
	"os"
)

func main() {

	applicao := app.Gerar()

	err := applicao.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
