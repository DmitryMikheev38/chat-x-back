package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DB ...
var DB *pgx.Conn

// Connect ...
func Connect() error {
	var err error
	dbConfig := viper.GetStringMap("postgres")

	DB, err = pgx.Connect(context.Background(), fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig["username"].(string),
		dbConfig["password"].(string),
		dbConfig["dbname"].(string),
		dbConfig["port"].(string),
	))

	if err != nil {
		return errors.Wrap(err, "Database connect #1: ")
	}

	err = DB.Ping(context.Background())
	if err != nil {
		return err
	}

	log.Println("Start postgres port=", dbConfig["port"])
	return nil
}
