package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

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
	createDb()
	initFirstnameTable()
	initSurnameTable()
	initStreetPrefixTable()
	initStreetSuffixTable()
	initPostalAddressTable()
	initPasswordTable()
	initEmailDomainsTable()
	initCompanynameTable()
}

func createDb() {
	path := "./db/idparts.db"
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

func initFirstnameTable() {
	src := "./db/fname"
	dropStmt := "DROP TABLE IF EXISTS firstnames"
	initStmt := "CREATE TABLE IF NOT EXISTS firstnames (id INTEGER PRIMARY KEY, firstname TEXT)"
	inputStmt := "INSERT INTO firstnames (firstname) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initSurnameTable() {
	src := "./db/lname"
	dropStmt := "DROP TABLE IF EXISTS surnames"
	initStmt := "CREATE TABLE IF NOT EXISTS surnames (id INTEGER PRIMARY KEY, surname TEXT)"
	inputStmt := "INSERT INTO surnames (surname) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initStreetPrefixTable() {
	src := "./db/stpre"
	dropStmt := "DROP TABLE IF EXISTS streetprefixes"
	initStmt := "CREATE TABLE IF NOT EXISTS streetprefixes (id INTEGER PRIMARY KEY, streetprefix TEXT)"
	inputStmt := "INSERT INTO streetprefixes (streetprefix) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initStreetSuffixTable() {
	src := "./db/stsuf"
	dropStmt := "DROP TABLE IF EXISTS streetsuffixes"
	initStmt := "CREATE TABLE IF NOT EXISTS streetsuffixes (id INTEGER PRIMARY KEY, streetsuffix TEXT)"
	inputStmt := "INSERT INTO streetsuffixes (streetsuffix) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initPasswordTable() {
	src := "./db/pwords"
	dropStmt := "DROP TABLE IF EXISTS passwords"
	initStmt := "CREATE TABLE IF NOT EXISTS passwords (id INTEGER PRIMARY KEY, password TEXT)"
	inputStmt := "INSERT INTO passwords (password) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initEmailDomainsTable() {
	src := "./db/emailaddresses"
	dropStmt := "DROP TABLE IF EXISTS emaildomains"
	initStmt := "CREATE TABLE IF NOT EXISTS emaildomains (id INTEGER PRIMARY KEY, emaildomain TEXT)"
	inputStmt := "INSERT INTO emaildomains (emaildomain) VALUES (?)"

	initTable(dropStmt, initStmt)
	populateTable(src, inputStmt)
}

func initCompanynameTable() {
	src := "./db/foretag"
	dropStmt := "DROP TABLE IF EXISTS companynameparts"
	initStmt := "CREATE TABLE IF NOT EXISTS companynameparts (id INTEGER PRIMARY KEY, companynamepart TEXT)"
	inputStmt := "INSERT INTO companynameparts (companynamepart) VALUES (?)"

	initTable(dropStmt, initStmt)
	file, err := os.Open(src)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	db, _ := sql.Open("sqlite3", "./db/idparts.db")
	statement, _ := db.Prepare(inputStmt)
	for scanner.Scan() {
		statement.Exec(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	statement.Close()
	db.Close()
}

func initPostalAddressTable() {
	src := "./db/pnumort"
	dropStmt := "DROP TABLE IF EXISTS postaladdresses"
	initStmt := "CREATE TABLE IF NOT EXISTS postaladdresses (id INTEGER PRIMARY KEY, postalcode TEXT, posttown TEXT)"
	inputStmt := "INSERT INTO postaladdresses (postalcode, posttown) VALUES (?, ?)"

	initTable(dropStmt, initStmt)

	file, err := os.Open(src)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	db, _ := sql.Open("sqlite3", "./db/idparts.db")
	statement, _ := db.Prepare(inputStmt)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
		if len(line) == 2 {
			statement.Exec(line[0], line[1])
			line = line[2:]
		}
	}
	statement.Close()
	db.Close()

}

func initTable(dropStmt string, initStmt string) {
	db, _ := sql.Open("sqlite3", "./db/idparts.db")
	statement, _ := db.Prepare(dropStmt)
	statement.Exec()
	statement, _ = db.Prepare(initStmt)
	statement.Exec()
	statement.Close()

	db.Close()
}

func populateTable(sourcefile string, inputStmt string) {
	file, err := os.Open(sourcefile)
	check(err)

	scanner := bufio.NewScanner(file)
	db, _ := sql.Open("sqlite3", "./db/idparts.db")
	statement, _ := db.Prepare(inputStmt)
	for scanner.Scan() {
		statement.Exec(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	statement.Close()
	db.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
