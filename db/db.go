package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() {
	conn := "server=localhost;database=crm_exotel;trusted_connection=yes"
	var err error
	DB, err = sql.Open("sqlserver", conn)
	if err != nil {
		log.Fatal("❌ Cannot open DB:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Cannot connect to DB:", err)
	}
	fmt.Println("✅ Connected to SQL Server")
}
