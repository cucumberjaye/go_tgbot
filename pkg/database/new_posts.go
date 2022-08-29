package database

import (
	"database/sql"
	"github.com/cucumberjaye/go_tgbot/pkg/habr"
	_ "github.com/lib/pq"
	"log"
)

func InitLastPost() error {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	defer db.Close()

	data := `INSERT INTO last_post(post) VALUES ($1);`

	ids, err := habr.GetPostsIds()
	if err != nil {
		return err
	}
	if _, err = db.Exec(data, ids.Id[0]); err != nil {
		log.Fatalf("lol %s", err.Error())
		return err
	}

	return nil
}

func GetUrls() ([]string, error) {
	var post string

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(`SELECT post FROM last_post;`)
	err = row.Scan(&post)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	ids, err := habr.GetPostsIds()
	if err != nil {
		return nil, err
	}
	return habr.GetNewPostsUrls(post, ids), nil
}
