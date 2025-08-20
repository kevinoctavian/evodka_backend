package database

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/kevinoctavian/evodka_backend/pkg/config"
)

type DB struct{ *sqlx.DB }

var defaultDB = &DB{}

func (db *DB) connect(dbCfg *config.DB) (err error) {
	dbUri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Name,
		dbCfg.SslMode,
	)

	if dbCfg.Password != "" {
		dbUri += fmt.Sprintf(" password=%s", dbCfg.Password)
	}

	db.DB, err = sqlx.Connect("pgx", dbUri)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(dbCfg.MaxOpenConn)
	db.SetMaxIdleConns(dbCfg.MaxIdleConn)
	db.SetConnMaxLifetime(dbCfg.MaxConnLifetime)

	if err = db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("failed to send ping to database: %w", err)
	}

	return nil
}

func GetDB() *DB {
	if defaultDB == nil {
		log.Fatal("Default database connection is not initialized")
	}
	return defaultDB
}

func ConnectDB() error {
	return defaultDB.connect(config.DBCfg())
}
