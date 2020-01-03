package main

import (
	"market/controllers"
	"market/db"
	"net/http"

	"github.com/jackc/pgx"
	"github.com/julienschmidt/httprouter"
)

var (
	pgPool *pgx.ConnPool
	router = httprouter.New()
)

func main() {

	//Подключаемся к Базе Данных
	//Оставляем на автовыключение
	db.Connect()
	defer db.Close()

	router.GET("/products", controllers.GetProductsHandler)
	router.POST("/newclient", controllers.CreateNewClientHandler)
	router.GET("/clients", controllers.GetClientsHandler)
	router.DELETE("/deleteclient/:id", controllers.DeleteClientHandler)

	http.ListenAndServe(":8080", router)
}

/*
func afterConnect(con *pgx.Conn) (err error) {
	return
}


//Connect - connects database
func Connect() {

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			User:     "postgres",
			Password: "918813181s",
			Database: "datamarket",
		},
		MaxConnections: runtime.NumCPU() * 2,
		AfterConnect:   afterConnect,
	}

	var err error
	pgPool, err = pgx.NewConnPool(connPoolConfig)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}

	fmt.Printf("Hello world\n")
}

// Close - disconnects database connection
func Close() {
	pgPool.Close()
}

///*/
