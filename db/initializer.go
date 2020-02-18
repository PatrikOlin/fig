package db

import (
	"database/sql"
	"os"
	"bufio"
	"log"
		
	_ "github.com/mattn/go-sqlite3"
)

func main() {
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}		

func Initdb() {
	initFirstnameTable()
	initSurnameTable()
	initStreetPrefixTable()
	initStreetSuffixTable()
}		
		
func initFirstnameTable() {
	fname := "./db/fname"
	initStmt := "CREATE TABLE IF NOT EXISTS firstnames (id INTEGER PRIMARY KEY, firstname TEXT)"
	inputStmt := "INSERT INTO firstnames (firstname) VALUES (?)"

	initTable(fname, initStmt, inputStmt)
}

func initSurnameTable() {
	sname := "./db/lname"
	initStmt := "CREATE TABLE IF NOT EXISTS surnames (id INTEGER PRIMARY KEY, surname TEXT)"
	inputStmt := "INSERT INTO surnames (surname) VALUES (?)"

	initTable(sname, initStmt, inputStmt)
}

func initStreetPrefixTable() {
	stpre := "./db/stpre"
	initStmt := "CREATE TABLE IF NOT EXISTS streetprefixes (id INTEGER PRIMARY KEY, streetprefix TEXT)"
	inputStmt := "INSERT INTO streetprefixes (streetprefix) VALUES (?)"

	initTable(stpre, initStmt, inputStmt)
}

func initStreetSuffixTable() {
	stsuf := "./db/stsuf"
	initStmt := "CREATE TABLE IF NOT EXISTS streetsuffixes (id INTEGER PRIMARY KEY, streetsuffix TEXT)"
	inputStmt := "INSERT INTO streetsuffixes (streetsuffix) VALUES (?)"

	initTable(stsuf, initStmt, inputStmt)
}

func initTable(sourcefile string, initStmt string, inputStmt string) {
	file, err := os.Open(sourcefile)
	check(err)

	scanner := bufio.NewScanner(file)
	db, _ := sql.Open("sqlite3", "./db/idparts.db")
	statement, _ := db.Prepare(initStmt)
	statement.Exec()
	statement, _ = db.Prepare(inputStmt)
	for scanner.Scan() {
		statement.Exec(scanner.Text())
	}
	statement.Close()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	db.Close()
}
