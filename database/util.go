package database

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	"github.com/hatobus/study-mysql/database/config"
)

func ConnectDB(conf *config.MySQLConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%%2FTokyo",
		conf.DBUser,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DataBase,
	)

	return sql.Open("mysql", dsn)
}

func PrepareDB(dbName string, conf *config.MySQLConfig) error {
	db, err := ConnectDB(conf)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", dbName))
	if err != nil {
		return err
	}

	return nil
}

func RunSQLScript(filePath, dbName string, conf *config.MySQLConfig) error {
	if exist := fileExists(filePath); !exist {
		return fmt.Errorf("file: %v is not found", filePath)
	}

	err := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf("mysql -h %s -u %s -p%s %s < %s", conf.Host, conf.DBUser, conf.Password, dbName, filePath),
	).Run()

	if err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
