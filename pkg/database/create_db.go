package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	host     = os.Getenv("HOST")
	port     = os.Getenv("PORT")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbname   = os.Getenv("DBNAME")
	sslmode  = os.Getenv("SSLMODE")
	dbInfo   = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
)

func CreateTable() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	defer db.Close()

	if _, err = db.Query("CREATE TABLE last_post(POST varchar(255));"); err != nil && fmt.Sprintf("%s", err.Error()) != `pq: relation "last_post" already exists` {
		log.Fatalf(err.Error())
		return err
	}

	if err := InitLastPost(); err != nil {
		log.Fatalf(err.Error())
		return err
	}

	return nil
}
