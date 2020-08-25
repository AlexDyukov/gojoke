# Архитектура приложения

## Авторизация
ТЗ требует идентификации пользователей, но среди эндпоинтов нет URL для авторизации (допустим за него "не заплатили").
Необходим механизм идентификации такой, что при расширении ТЗ методом авторизации процесс миграции был:
- максимально прозрачен для пользователей
- максимально прост для служб эксплуатации
- с минимальными простоями сервиса при миграции

Обычно для хранения данных с привязкой к пользователю используют структуру таблицы с foreign key:
```
CREATE TABLE users (
	userid SERIAL PRIMARY KEY,
	name VARCHAR(20) NOT NULL,
	email VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE jokes (
	jokeid SERIAL PRIMARY KEY,
	userid INTEGER REFERENCES users (userid),
	joke_str VARCHAR UNIQUE NOT NULL
)
```
С текущим ТЗ нам нет нужды держать users и занимать место в table space, однако рассчитывать на её создание нужно, поэтому описываем только таблицу jokes с расчётом в будущем добавить users:
```
CREATE SEQUENCE IF NOT EXISTS user_sequence;
CREATE TABLE IF NOT EXISTS jokes (
	jokeid SERIAL PRIMARY KEY,
	userid INTEGER NOT NULL,
	joke_str VARCHAR UNIQUE NOT NULL
);
```

## БД
Модуль работы с БД должен обеспечивать:
- прозрачные миграции БД. Сервис должен производить миграции самостоятельно т.к. код миграции, интегрированный в сам сервис, можно покрыть тестами, а значит исключить фактор человеческой ошибки
- поддержку нескольких БД (не считая экзотических вариантов типа berkeleyDB, необходимым минимумом считается MariaDB/PostgreSQL/SQLite)

