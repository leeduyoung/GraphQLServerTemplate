package db

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

type multiDriver struct {
	replica dialect.Driver
	master  dialect.Driver
}

var _ dialect.Driver = (*multiDriver)(nil)

func (d *multiDriver) Query(ctx context.Context, query string, args, v any) error {
	return d.replica.Query(ctx, query, args, v)
}

func (d *multiDriver) Exec(ctx context.Context, query string, args, v any) error {
	return d.master.Exec(ctx, query, args, v)
}

func (d *multiDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.master.Tx(ctx)
}

func (d *multiDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	return d.master.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
}

func (d *multiDriver) Close() error {
	rerr := d.replica.Close()
	werr := d.master.Close()
	if rerr != nil {
		return rerr
	}
	if werr != nil {
		return werr
	}
	return nil
}

func (d *multiDriver) Dialect() string {
	return dialect.MySQL
}
