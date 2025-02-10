package migration

import (
	"database/sql"
	"log"
)

func CreateTableCustomer(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS customers (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        phone VARCHAR(20) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Gagal membuat tabel Customer:", err)
	}

	log.Println("Tabel Customer berhasil dibuat!")
}
