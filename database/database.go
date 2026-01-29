package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	if connectionString == "" {
		log.Fatal("DB_CONN environment variable is empty")
	}

	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Pengaturan koneksi agar stabil dengan Pooler Supabase
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Coba ping sekali untuk memastikan kredensial benar
	// Jika gagal di sini, berarti URL atau Password salah
	err = db.Ping()
	if err != nil {
		log.Printf("Ping failed: %v", err)
		// Kita tidak log.Fatal agar server tidak restart terus-menerus di Railway
		// Tapi kita kembalikan error agar terlihat di log
		return db, err
	}

	log.Println("Database connection established and verified")
	return db, nil
}

func Migrate(db *sql.DB) error {
	log.Println("Attempting manual migration...")
	// Query tetap sama
	queryCategories := `CREATE TABLE IF NOT EXISTS categories (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, description TEXT);`
	queryProducts := `CREATE TABLE IF NOT EXISTS products (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, price INT NOT NULL, stock INT NOT NULL, category_id INT REFERENCES categories(id) ON DELETE SET NULL);`

	if _, err := db.Exec(queryCategories); err != nil {
		return err
	}
	if _, err := db.Exec(queryProducts); err != nil {
		return err
	}
	log.Println("Manual migration successful")
	return nil
}
