package group_concat

import (
	"log"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

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

	mysqlcnf.DataBase = dbName

	var db *sqlx.DB
	db, err = database.ConnectDB(mysqlcnf)
	if err != nil {
		log.Fatal(err)
	}

	var attachments string
	err = db.Get(&attachments, "select GROUP_CONCAT(attachment.attachment_name) from user, attachment where user.id = attachment.user_id and user_id = 1;")
	if err != nil {
		log.Println(err)
	}

	log.Println(attachments)

	log.Println(strings.Split(attachments, ","))
}
