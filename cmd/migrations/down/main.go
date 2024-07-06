package main

import (
	"log"

	"github.com/albugowy15/nearby-shops/internal/config"
)

var schema = `
DROP TABLE IF EXISTS shops;
`

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	db.MustExec(schema)
	db.Close()
	log.Println("DB Migrations Down success")
}
