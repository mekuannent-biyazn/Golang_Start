package config

import(
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(){
	connStr:="host=localhost port=5432 user=postgres password=Mekub@@3116 dbname=GolangDB sslmode=disable"

	db, err:= sql.Open(
		"postgres",
		connStr,
	)

	if err!= nil{
		panic(err)
	}

	errr:= db.Ping()
	if err!= nil{
		panic(errr)
	}

	fmt.Println("Database Connected Successfully!!")

	DB = db
}