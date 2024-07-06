package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewDatabase(viper *viper.Viper) *sqlx.DB {
	db, err := sqlx.Connect("postgres", viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error connecting with database: %v", err)
	}
	return db
}
