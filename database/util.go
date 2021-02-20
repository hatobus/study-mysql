package database

import (
	"database/sql"
	"fmt"

	"github.com/hatobus/study-mysql/database/config"
)

func ConnectDB(conf config.MySQLConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%%2FTokyo",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DataBase,
	)

	return sql.Open("mysql", dsn)
}
