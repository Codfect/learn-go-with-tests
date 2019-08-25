package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/larien/learn-go-with-tests/linha-de-comando/v3"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Vamos jogar poker")
	fmt.Println("Digite {Nome} venceu para registrar uma vitoria")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
