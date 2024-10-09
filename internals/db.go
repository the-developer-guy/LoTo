package internals

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseHelper struct {
	Path     string
	Database *sql.DB
}

func NewDatabaseHelper() (*DatabaseHelper, error) {

	dbPath := os.Getenv("LOTO_SQLITE_PATH")
	if dbPath == "" {
		return nil, fmt.Errorf("LOTO_SQLITE_PATH not set")
	}

	dh := DatabaseHelper{
		Path: dbPath,
	}

	var err error
	dh.Database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	err = dh.dbInit()
	if err != nil {
		return nil, err
	}

	return &dh, nil
}

func (dh *DatabaseHelper) InitConfig(service *Service) error {

	return nil
}

func (dh *DatabaseHelper) dbInit() error {
	initStatement := `CREATE TABLE IF NOT EXISTS services (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT NOT NULL UNIQUE,
		"url" TEXT,
		"reservedBy" TEXT	
	  );`
	statement, err := dh.Database.Prepare(initStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	return err
}
