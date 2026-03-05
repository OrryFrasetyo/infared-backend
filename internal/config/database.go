package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnectDB() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Tidak menemukan file .env, menggunakan environment system")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Error: DATABASE_URL belum di-set di file .env")
	}

	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak merespons: %v", err)
	}

	fmt.Println("✅ [Database] Berhasil terhubung ke PostgreSQL InfaRed!")
	return db
}
