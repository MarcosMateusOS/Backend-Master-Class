package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"fmt"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

/* Conversão em testes unitarios do GoLang, o TestMain chama todos os teste*/
func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err = testDB.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Ping status: OK")
	}


	testQueries = New(testDB)
	os.Exit(m.Run())
}