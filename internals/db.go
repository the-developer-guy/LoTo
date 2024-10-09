package internals

import (
	"database/sql"
	"fmt"
	"log"
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
	row := dh.Database.QueryRow("SELECT reservedBy FROM services WHERE name=?", service.Name)
	var reservedBy sql.NullString
	err := row.Scan(&reservedBy)
	if err != nil {
		err = dh.InsertService(service)
		return err
	}
	if !reservedBy.Valid {
		service.Locked = false
		return nil
	}
	service.Locked = reservedBy.String != ""

	return nil
}

func (dh *DatabaseHelper) InsertService(service *Service) error {
	insert := `INSERT INTO services(name, url, reservedBy) VALUES (?, ?, ?)`
	statement, err := dh.Database.Prepare(insert)
	if err != nil {
		return err
	}

	lockedBy := ""
	if service.Locked {
		lockedBy = "placeholder"
	}
	_, err = statement.Exec(service.Name, service.Url, lockedBy)
	if err != nil {
		return err
	}

	return nil
}

func (dh *DatabaseHelper) Lock(name string) error {
	lockStatement := `UPDATE services SET reservedBy='placeholder'
					  WHERE name=?`
	statement, err := dh.Database.Prepare(lockStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	return err
}

func (dh *DatabaseHelper) Unlock(name string) error {
	lockStatement := `UPDATE services SET reservedBy=''
					  WHERE name=?`
	statement, err := dh.Database.Prepare(lockStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name)
	return err
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

func (dh *DatabaseHelper) dbDump() error {
	row, err := dh.Database.Query(`SELECT name, url, reservedBy FROM services`)
	if err != nil {
		return err
	}
	defer row.Close()

	for row.Next() {
		var name, url, reservedBy sql.NullString
		row.Scan(&name, &url, &reservedBy)
		if name.Valid {
			log.Print(name.String)
		}
		if url.Valid {
			log.Print(url.String)
		}
		if reservedBy.Valid {
			log.Print(reservedBy.String)
		}
		log.Println()

	}
	return err
}
