package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// 1. Inisialisasi driver dengan pgx
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	// 2. Setting koneksi agar tidak cepat timeout
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(15 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// 3. Tes koneksi saat startup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database!")
	return db, nil
}

func Migrate(db *sql.DB) error {
	// Membuat tabel jika belum ada
	query := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT
	);
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		stock INT NOT NULL,
		category_id INT REFERENCES categories(id) ON DELETE SET NULL
	);`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Database migration completed successfully!")
	return nil
}
