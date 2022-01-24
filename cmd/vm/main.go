package main

import (
	"github.com/DamirLuketic/virtual_minds/config"
	"github.com/DamirLuketic/virtual_minds/db"
	"log"
)

func main() {
	c := config.NewServerConfig()
	db.NewMariaDBDataStore(c)
	log.Printf("App has landed.")
}
