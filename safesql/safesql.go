package safesql

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/empijei/def-prog-exercises/safesql/internal/raw"
)

func init() {
	raw.TrustedSQLCtor =
		func(unsafe string) TrustedSQL {
			return TrustedSQL{unsafe}
		}
}

/***********
* Safe SQL *
************/

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string(text)}
}

func NewFromInt(i int) TrustedSQL {
	return TrustedSQL{strconv.Itoa(i)}
}

/***********
* SQL Wrap *
************/

/* Known safe types */

type (
	Result = sql.Result
	Rows   = sql.Rows
)

/* Wrappers */

type DB struct {
	db *sql.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	d, err := sql.Open(driverName, dataSourceName)
	return &DB{d}, err
}

func (db *DB) QueryContext(ctx context.Context,
	query TrustedSQL, args ...any) (*Rows, error) {
	return db.db.QueryContext(ctx, query.s, args...)
}
func (db *DB) ExecContext(ctx context.Context,
	query TrustedSQL, args ...any) (Result, error) {
	return db.db.ExecContext(ctx, query.s, args...)
}
