package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Supabase Free Tier punya limit koneksi kecil
	// Kita set lebih rendah agar tidak kena kicked oleh Pooler
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	// Run auto migration
	err = Migrate(db)
	if err != nil {
		log.Println("Migration failed:", err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	log.Println("Running auto migration...")

	queryCategories := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT
	);`

	queryProducts := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		stock INT NOT NULL,
		category_id INT REFERENCES categories(id) ON DELETE SET NULL
	);`

	_, err := db.Exec(queryCategories)
	if err != nil {
		return err
	}

	_, err = db.Exec(queryProducts)
	if err != nil {
		return err
	}

	log.Println("Auto migration completed")
	return nil
}
