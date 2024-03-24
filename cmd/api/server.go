package api

import (
	"assignment3/ent"
	"assignment3/ent/user"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"entgo.io/ent/dialect"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var ctx *fiber.Ctx
var DB *ent.Client

func ConnectDB() *ent.Client {
	db, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=root dbname=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
		os.Exit(1)
	}
	// Run migration.
	ctx := context.Background()
	err = db.Schema.Create(ctx)
	if err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
	//Create admin if do not have
	_, err = db.User.Query().Where(user.UsernameEQ("Admin")).Only(ctx)
	if err != nil {
		current_time := time.Now()
		adminPassword := "password"
		hashedPassword, err := HashPassword(adminPassword)
		if err != nil {
			log.Fatal(err)
		}
		role_admin := db.Role.Create().
			SetRoleName("ADMIN").
			SaveX(ctx)
		_ = db.User.Create().
			SetUsername("Admin").
			SetPassword(hashedPassword).
			SetRole(role_admin).
			SetCreatedAt(current_time).
			SaveX(ctx)
	}

	fmt.Println("Connected to postgres!")
	DB = db
	return DB
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
