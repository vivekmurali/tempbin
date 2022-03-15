package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

type DBEnv struct {
	DB *pgxpool.Pool
}

var Env DBEnv

func InitDB() {

	godotenv.Load()
	DB, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	Env.DB = DB
}
