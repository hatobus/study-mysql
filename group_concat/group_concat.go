package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hatobus/study-mysql/database"
	"github.com/hatobus/study-mysql/database/config"
)

func main() {
	mysqlcnf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = database.PrepareDB("group_concat", mysqlcnf)
	if err != nil {
		log.Fatal(err)
	}

}
