package main

import (
	"log"

	serverInit "github.com/danilotadeu/r-customer-code-information/server"
	"github.com/joho/godotenv"
)

var (
	server serverInit.Server
)

func init() {
	//Deletar devido nao funcionar assim no ambiente real
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server = serverInit.New()
}

func main() {
	server.Start()
}
