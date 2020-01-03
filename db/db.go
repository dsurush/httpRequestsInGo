package db

import (
	"fmt"
	"log"
	"market/config"
	"runtime"

	"github.com/jackc/pgx"
)

var (
	pgPool *pgx.ConnPool
)

func afterConnect(con *pgx.Conn) (err error) {
	_, err = con.Prepare("getProducts",
		`Select *From public.product
		ORDER BY name`)
	if err != nil {
		log.Println("getProducts")
		return
	}

	_, err = con.Prepare("addClient",
		`INSERT INTO "costumer"(
			id, name, surname, address, email, phonenumbers)
		VALUES ($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		log.Println("addClient", err)
		return
	}
	_, err = con.Prepare("getClients",
		`Select *From public.costumer
	ORDER BY name`)
	if err != nil {
		log.Println("getClients", err)
		return
	}
	_, err = con.Prepare("deleteClient",
		`DELETE FROM "costumer" c
		WHERE c.id = $1`)
	if err != nil {
		log.Println("deleteClient", err)
		return
	}
	return
}

// Connect "Подключение к базе данных"
func Connect() {
	db := config.Peek().Database
	db.Addr = "localhost"
	db.User = "postgres"
	db.Pass = "918813181s"
	db.DBName = "datamarket"
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     db.Addr,
			User:     db.User,
			Password: db.Pass,
			Database: db.DBName,
		},
		MaxConnections: runtime.NumCPU() * 2,
		AfterConnect:   afterConnect,
	}

	var err error
	pgPool, err = pgx.NewConnPool(connPoolConfig)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Printf("Hello world\n")
}

// Close "Закрывает доступ к базе данных"
func Close() {
	pgPool.Close()
}
