package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Todo struct {
	Test int    `db:"TEST"`
	Text string `db:"TEXT"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "test"
)

func main() {
	username := "xcode"
	password := "password"
	db, err := sqlx.Connect("godror", username+"/"+password+"@localhost:1521/XEPDB1")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Successfully connected to the Oracle database!")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	dbpsql, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	defer dbpsql.Close()

	// check db
	err = dbpsql.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected! Postgresql")

	todo := Todo{}
	rows, err := db.Queryx("SELECT * FROM TEST")

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		err := rows.StructScan(&todo)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("row[Test]: %d, row[Text]: %s\n", todo.Test, todo.Text)
		// insertStmt := `INSERT INTO public.test ("id", "text") VALUES (1, "awda")`
		_, error := dbpsql.Exec(`INSERT INTO public.test (id, text) VALUES ($1, $2);`, todo.Test, todo.Text)
		if error != nil {
			log.Fatalf("error insert: %v", error)
		}
	}

}
