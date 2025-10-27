package persistance

import (
	"database/sql"
	"fmt"

	"github.com/Gurveer1510/url-shortner/internal/config"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s&channel_binding=%s",cfg.PGUSER, cfg.PGPASSWORD, cfg.PGHOST, cfg.PGDATABASE, cfg.PGSSLMODE, cfg.PGCHANNELBINDING)

	fmt.Println("DATABASE URL:", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return  &Database{db: db}, nil
}

func (d *Database) GetDB() *sql.DB{
	return d.db
}