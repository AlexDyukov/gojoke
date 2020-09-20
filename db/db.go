package db

import (
	"errors"
	"os"
)

type Joke struct {
	Id  int    `json:"id"`
	Str string `json:"str"`
}

type JokeDatabase interface {
	Add(user_id int, joke_str string) (int, error)
	Select(joke_id int, user_id int) (string, error)
	List(user_id int) ([]Joke, error)
	Update(joke_id int, user_id int, joke_str string) error
	Delete(joke_id int, user_id int) error
}

func Init() (JokeDatabase, error) {
	if pg_url := os.Getenv("PG_URL"); len(pg_url) > 0 {
		return PsqlInit(pg_url)
	}
	return nil, errors.New("[db/Init] No db specified")
}
