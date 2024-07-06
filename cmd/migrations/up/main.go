package main

import (
	"log"

	"github.com/albugowy15/nearby-shops/internal/config"
)

var schema = `
CREATE TABLE "shops" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" text,
  "city" varchar(50) NOT NULL,
  "location" geography(POINT) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now())
);
`

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	db.MustExec(schema)
	db.Close()
	log.Println("DB Migrations Up success")
}
