package main

import (
	"log"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hatobus/study-mysql/database"
	"github.com/hatobus/study-mysql/database/config"
)

func main() {
	mysqlcnf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	dbName := "group_concat"

	err = database.PrepareDB(dbName, mysqlcnf)
	if err != nil {
		log.Fatal(err)
	}

	files := []string{
		"./scheme/scheme.sql",
		"./scheme/init_data.sql",
	}

	for _, f := range files {
		var fname string
		var err error
		if fname, err = filepath.Abs(f); err != nil {
			log.Fatal(err)
		}
		log.Println(fname)
		if err = database.RunSQLScript(fname, dbName, mysqlcnf); err != nil {
			log.Fatal(err)
		}
	}

}
