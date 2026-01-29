package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Open database (sql.Open tidak langsung melakukan koneksi ke DB)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Set limit seminimal mungkin
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(1)

	// Kita hapus Ping() agar server tidak crash saat startup jika DB delay
	log.Println("Database driver initialized (Lazy Connection)")

	return db, nil
}

func Migrate(db *sql.DB) error {
	// Fungsi ini tetap ada tapi tidak dipanggil otomatis di main.go
	return nil
}
