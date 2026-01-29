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

	// Batasi koneksi untuk efisiensi di Railway/Supabase
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)

	// Kita nonaktifkan Ping dan Migrate untuk menghindari EOF di Pooler
	// Pastikan tabel sudah dibuat manual di dashboard Supabase
	log.Println("Database connection string initialized")

	return db, nil
}

// Migrate tetap ada jika ingin dipanggil manual, tapi tidak dijalankan otomatis
func Migrate(db *sql.DB) error {
	log.Println("Running manual migration...")

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

	log.Println("Manual migration completed")
	return nil
}
