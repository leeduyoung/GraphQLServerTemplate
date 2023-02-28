package db

import (
	"context"
	"fmt"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/ent/migrate"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/config"
	"log"
	"sync"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	client *ent.Client = nil
	once   sync.Once
)

func New(mode pkg.Mode) *ent.Client {
	if client == nil {
		once.Do(func() {
			master, err := createDriver(MasterMode)
			if err != nil {
				log.Fatal(err)
			}

			replica, err := createDriver(ReplicaMode)
			if err != nil {
				log.Fatal(err)
			}

			client = createClient(mode, master, replica)

			cb := func(ctx context.Context, client *ent.Client) error {
				// DO SOMETHING
				return nil
			}

			ctx := context.Background()
			err = migration(ctx, client, cb)

			if err != nil {
				panic(err)
			}
		})
	}

	return client
}

func createDriver(driverMode DriverMode) (*sql.Driver, error) {
	return sql.Open(loadDialect(), loadDSN(driverMode))
}

func createClient(mode pkg.Mode, master *sql.Driver, replica *sql.Driver) *ent.Client {
	switch mode {
	case pkg.TestMode:
	case pkg.DevMode:
		return ent.NewClient(
			ent.Driver(
				&multiDriver{replica: replica, master: master},
			),
		).Debug()
	}

	return ent.NewClient(
		ent.Driver(
			&multiDriver{replica: replica, master: master},
		),
	)
}

func migration(ctx context.Context, client *ent.Client, fn func(ctx context.Context, client *ent.Client) error) error {
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(false),
		migrate.WithDropColumn(false),
	)
	if err != nil {
		return err
	}

	err = fn(ctx, client)
	if err != nil {
		return err
	}

	return nil
}

func loadDialect() string {
	return dialect.MySQL
}

func loadDSN(mode DriverMode) string {
	var (
		hostName    = config.Instance.Rds.MasterHostname
		rdsUserName = config.Instance.Rds.Username
		rdsPassword = config.Instance.Rds.Password
		rdsPort     = config.Instance.Rds.Port
		rdsDBName   = config.Instance.Rds.DBName
	)

	switch mode {
	case MasterMode:
		hostName = config.Instance.Rds.MasterHostname
	case ReplicaMode:
		hostName = config.Instance.Rds.ReplicaHostname
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=UTC",
		rdsUserName,
		rdsPassword,
		hostName,
		rdsPort,
		rdsDBName,
	)
}
