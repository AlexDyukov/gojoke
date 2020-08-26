package db

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

type psql struct {
	pool *sql.DB
}

func PsqlInit(max_connections int) (JokeDatabase, error) {
	conn_config, err := pgx.ParseEnvLibpq()
	if err != nil {
		return nil, errors.New("[plsql_init] Unexpected error from pgx.ParseEnvLibpq")
	}

	config := pgx.ConnPoolConfig{ConnConfig: conn_config, MaxConnections: max_connections}
	pool, err := pgx.NewConnPool(config)
	if err != nil {
		return nil, errors.New("[plsql_init] Unable to create connection pool")
	}

	return psql{pool: stdlib.OpenDBFromPool(pool)}, nil
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
