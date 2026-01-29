package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// 1. Inisialisasi driver
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// 2. Setting koneksi agar tidak cepat timeout
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(15 * time.Minute)

	// 3. Tes koneksi saat startup
	// Kita gunakan loop kecil agar memberi waktu DB untuk siap
	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("Successfully connected to database!")
			return db, nil
		}
		log.Printf("Attempt %d: Could not connect to database, retrying... (%v)", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return db, err
}

func Migrate(db *sql.DB) error {
	// Kita kembalikan fitur Migrate agar tabel dibuat otomatis
	// asalkan kita sudah pindah ke Port 5432 (Session Mode)
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

	_, err := db.Exec(query)
	return err
}
