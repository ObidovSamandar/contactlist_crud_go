package dbconnector

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DBClient *sqlx.DB

const schema = `
CREATE SEQUENCE contactlistSquence START 1;

CREATE TABLE contactlistdb (
	id integer,
	firstname varchar(40),
	lastname varchar(40),
	phone varchar(40),
	email varchar(40)
)
`

func DBClientConnector() {

	godotenv.Load()

	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=mylocaldatabase sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(schema)

	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}
