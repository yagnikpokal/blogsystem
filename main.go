package main

import (
	"backend/api"
	portconst "backend/pkg/const"
	"backend/pkg/db"
	"backend/pkg/repository/dbrepo"
	"time"

	_ "github.com/lib/pq"

	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// set application config
	var app api.Application
	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=postgres port=5432 user=postgres password=postgres dbname=articles sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()
	fmt.Println("Wait for database container to start up")
	time.Sleep(5 * time.Second)

	// connect to the database
	conn, err := db.ConnectToDB(app.DSN)
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	// Create the table if not exist in container
	app.DB.CreateTable()

	log.Println("Starting application on port", portconst.Port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", portconst.Port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
