package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

type psql struct {
	pool *sql.DB
}

func PsqlInit(db_url string) (JokeDatabase, error) {
	db, err := sql.Open("pgx", db_url)
	if err != nil {
		format_str := "[db/PsqlInit] Unexpected error from sql.Open: %s"
		return nil, fmt.Errorf(format_str, err.Error())
	}
	return psql{pool: db}, nil
}

func (p psql) Add(user_id int, joke_str string) (int, error) {
	return 0, nil
}
func (p psql) Select(joke_id int, user_id int) (string, error) {
	return "", nil
}
func (p psql) List(user_id int) ([]Joke, error) {
	return nil, nil
}
func (p psql) Update(joke_id int, user_id int, joke_str string) error {
	return nil
}
func (p psql) Delete(joke_id int, user_id int) error {
	return nil
}
