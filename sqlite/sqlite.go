package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateConnection() (*sql.DB, bool) {
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./nilav.db")

	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	fmt.Println(db)
	//	defer db.Close()
	return db, true

}

func createTable(query string, db *sql.DB) {
	_, err := db.Exec(query)
	//	_, err = db.Exec("create table employee (id INTEGER PRIMARY KEY   AUTOINCREMENT, name TEXT NOT NULL, address TEXT NOT NULL)")
	if err != nil {
		fmt.Println("error while creating")
		fmt.Println(err)
	}
}

func insertTable(query string, db *sql.DB) {

	stmt, err := db.Prepare(query)
	//	stmt, err := db.Prepare("INSERT INTO employee(name, address) values(?,?)")
	if err != nil {
		fmt.Println("error while inserting")
		log.Fatal(err)
	}

	defer stmt.Close()
	res, err := stmt.Exec("nilav", 123)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}

func retrieveTable(query string, db *sql.DB) {

	rows, err := db.Query("select * from employee")
	if err != nil {
		log.Fatal(err)
	}

	var name string
	var address string
	var id int

	for rows.Next() {
		err = rows.Scan(&id, &name, &address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
		fmt.Println(address)
	}
	rows.Close()
}

func deleteTable(query string, db *sql.DB) {
	stmt, err := db.Prepare("delete from employee where id=?")
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}

	res, err := stmt.Exec(1)
	if err != nil {
		fmt.Println("res error")
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(affect)
}

func main() {
	//	var create string
	//	var insert string
	//	var retrieve string
	//	var remove string

	create := "create table employee (id INTEGER PRIMARY KEY   AUTOINCREMENT, name TEXT NOT NULL, address TEXT NOT NULL)"
	insert := "INSERT INTO employee(name, address) values(?,?)"
	retrieve := "select * from employee"
	remove := "delete from employee where id=?"
	var db *sql.DB

	db, _ = CreateConnection()

	createTable(create, db)
	insertTable(insert, db)
	retrieveTable(retrieve, db)
	deleteTable(remove, db)
}
