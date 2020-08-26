package db

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

func DBInit(max_connections int) JokeDatabase {
	if true {
		jdb, err := PsqlInit(max_connections)
		if err != nil {
			panic(err)
		}
		return jdb
	}
	return nil
}
