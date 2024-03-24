package api

import (
	"assignment3/ent"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/gofiber/fiber/v2"
)

var ctx *fiber.Ctx
var DB *ent.Client

func ConnectDB() *ent.Client {
	db, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
		os.Exit(1)
	}
	fmt.Println("Connected to postgres!")
	DB = db
	return DB
}
