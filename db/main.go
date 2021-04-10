package db

import (
	"database/sql"
	"fmt"
)

type DbCon struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type NewDbCon interface {
	Connect() bool
}

func (d *DbCon) Connect() bool {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.User, d.Password, d.Dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if !CheckError(err) {
		return false
	}

	// close database
	defer db.Close()
	return true
}

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
