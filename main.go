package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

// Configuração do Banco de Dados
var server = "HOST"
var port = 8080
var user = "sa"
var password = "123"

func main() {
	http.HandleFunc("/data/api/info/", handler)
	http.ListenAndServe(":3000", nil)
}

//Função de Handler para retorno da rota
func handler(writer http.ResponseWriter, request *http.Request) {
	var db = getDb()

	var (
		result string
	)
	result = "{}"

	rows, err := db.Query("SELECT @@version")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Data: ", result)
		fmt.Fprintf(writer, "{Data: "+result+"}")
	}

	defer db.Close()
}

//Cria uma conexão e retorna
func getDb() *sql.DB {
	var db *sql.DB
	var err error

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	return db
}
