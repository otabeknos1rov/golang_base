package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root123"
	dbname   = "golang"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// insertStmt := `insert into "students"("name", "degree") values('Saloh', 4)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	deleteQuery := `delete from "students" where name='Saloh'`
	_, e := db.Exec(deleteQuery)
	CheckError(e)

	// // dynamic
	// insertDynStmt := `insert into "students"("name", "degree") values('Otabek', 5)`
	// _, e = db.Exec(insertDynStmt)
	// CheckError(e)

	selectAllStudents := `select * from "students"`
	students, er := db.Query(selectAllStudents)
	CheckError(er)

	// updateStmt := `update "students" set "name"=$1, "degree"=$2 where "degree"=1`
	// _, e := db.Exec(updateStmt, "Changed Name", 3)
	// CheckError(e)

	defer students.Close()
	for students.Next() {

		var name string
		var roll int

		err = students.Scan(&name, &roll)
		CheckError(err)

		fmt.Println(name, roll)

	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
