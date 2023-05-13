package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	testSqlStatement := `INSERT INTO "user" (username) VALUES ('test03')`
	_, err = db.Exec(testSqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Println("All done!")
}
