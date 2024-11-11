package safesql

import (
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}
type Rows = sql.Rows
type Result = sql.Result

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	rows, err := db.db.QueryContext(ctx, string(query.s), args...)
	return rows, err
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (sql.Result, error) {
	return db.db.ExecContext(ctx, string(query.s), args...)
}

func (db *DB) Open(driverName, dataSourceName string) error {
	dbConn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	db.db = dbConn
	return nil
}
