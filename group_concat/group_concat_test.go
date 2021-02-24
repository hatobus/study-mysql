package group_concat

import (
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/google/go-cmp/cmp"
	"github.com/hatobus/study-mysql/database"
	"github.com/hatobus/study-mysql/database/config"
	"github.com/jmoiron/sqlx"
)

func prepareDB(t testing.TB, mysqlcnf *config.MySQLConfig) *sqlx.DB {
	t.Helper()

	err := database.PrepareDB(mysqlcnf.DataBase, mysqlcnf)
	if err != nil {
		t.Fatal(err)
	}

	files := []string{
		"./scheme/scheme.sql",
		"./scheme/init_data.sql",
	}

	for _, f := range files {
		var fname string
		var err error
		if fname, err = filepath.Abs(f); err != nil {
			t.Fatal(err)
		}
		t.Log(fname)
		if err = database.RunSQLScript(fname, mysqlcnf.DataBase, mysqlcnf); err != nil {
			t.Fatal(err)
		}
	}

	db, err := database.ConnectDB(mysqlcnf)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func attachmentsFromUserID(userID int, db *sqlx.DB) ([]string, error) {
	var attachments string
	err := db.Get(&attachments, "select GROUP_CONCAT(attachment.attachment_name) from user, attachment where user.id = attachment.user_id and user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return strings.Split(attachments, ","), nil
}

func TestGroupConcat(t *testing.T) {
	mysqlcnf, err := config.Init()
	if err != nil {
		t.Fatal(err)
	}

	mysqlcnf.DataBase = "group_concat"

	db := prepareDB(t, mysqlcnf)

	user1Data, err := attachmentsFromUserID(1, db)
	if err != nil {
		t.Fatal(err)
	}

	expectUser1Data := []string{
		"user1_file1",
		"user1_file2",
		"user1_file3",
	}

	if diff := cmp.Diff(user1Data, expectUser1Data); diff != "" {
		t.Fatalf("unexpected output diff: %v", diff)
	}
}
