package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	err error
)

const (
	host = "localhost"
	user = "admin"
	password = "postgres"
	dialect = "postgres"
	port = 5432
	dbname = "assignment2"
)

func handleDatabaseConnection(){
	pqslInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open(dialect, pqslInfo)

	if err != nil {
		log.Panic("error occured while validating database arguments:", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while opening a connection to database:", err.Error())
	}

}

func createRequiredTables(){
	createOrdersTableQuery := `
		CREATE TABLE IF NOT EXISTS "orders" (
			order_id SERIAL PRIMARY KEY NOT NULL,
			customer_name VARCHAR(32) NOT NULL,
			ordered_at timestamptz DEFAULT (now() AT TIME ZONE 'utc'),
			created_at timestamptz DEFAULT (now() AT TIME ZONE 'utc'),
			updated_at timestamptz DEFAULT (now() AT TIME ZONE 'utc')
		);
	`

	_, err = db.Exec(createOrdersTableQuery)

	if err != nil {
		log.Fatal("error while creating orders table:", err.Error())
	}

	createItemsTableQuery := `
		CREATE TABLE IF NOT EXISTS "items" (
			item_id SERIAL PRIMARY KEY NOT NULL,
			item_code VARCHAR(16) NOT NULL,
			description VARCHAR(255) NOT NULL,
			quantity int NOT NULL,
			created_at timestamptz DEFAULT (now() AT TIME ZONE 'utc'),
			updated_at timestamptz DEFAULT (now() AT TIME ZONE 'utc'),
			order_id int NOT NULL,
			CONSTRAINT fk_order FOREIGN KEY(order_id) REFERENCES orders(order_id)
		);
	`

	_, err := db.Exec(createItemsTableQuery)

	if err != nil {
		log.Fatal("error while create items table:", err.Error())
	}
	
}

func InitializeDatabase(){
	handleDatabaseConnection()
	createRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}