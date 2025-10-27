package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
)

type Migrate struct {
	path           string
	db             *sql.DB
	migrationFiles []fs.DirEntry
	txn            *sql.Tx
}

func NewMigrate(db *sql.DB, path string) *Migrate {
	return &Migrate{
		db:   db,
		path: path,
	}
}

func (m *Migrate) RunMigrations() error {
	rawEntries, err := os.ReadDir(m.path)
	if err != nil {
		fmt.Printf("error reading directory: %v", err)
		return err
	}

	files := getFilesFromDirectory(rawEntries)

	m.migrationFiles = files

	m.txn, err = m.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer m.txn.Rollback()

	err = m.parseAndMigrateSQL()

	if err != nil {
		return err
	}

	if err = m.txn.Commit(); err != nil {
		return err
	}

	return nil
}

func getFilesFromDirectory(dirEntries []fs.DirEntry) []fs.DirEntry {
	result := []fs.DirEntry{}
	for _, v := range dirEntries {
		if !v.IsDir() {
			result = append(result, v)
		}
	}
	return result
}

func (m *Migrate) parseAndMigrateSQL() error {
	for _, file := range m.migrationFiles {
		filepath := m.path + "/" + file.Name()
		fmt.Printf("Reading file: %s\n", filepath)
		bytes, err := os.ReadFile(filepath)
		if err != nil {
			return err
		}
		content := string(bytes)
		log.Print(content)
		_, err = m.txn.Exec(content)
		if err != nil {
			return err
		}
	}
	return nil
}
