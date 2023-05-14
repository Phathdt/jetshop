package migrator

import (
	"database/sql"
	"flag"

	_ "github.com/pkg/pq"
	"github.com/pressly/goose/v3"
	sctx "jetshop/pkg/service-context"
)

const dialect = "postgres"

type migrator struct {
	id     string
	dsn    string
	logger sctx.Logger
}

func NewMigrator(id string) *migrator {
	return &migrator{id: id}
}

func (m *migrator) ID() string {
	return m.id
}

func (m *migrator) InitFlags() {
	flag.StringVar(&m.dsn, "db_dsn", "", "database connection-string")
}

func (m *migrator) Activate(sc sctx.ServiceContext) error {
	m.logger = sc.Logger(m.id)
	db, err := sql.Open("postgres", m.dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	if err = goose.SetDialect(dialect); err != nil {
		return err
	}

	if err = goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}

func (m *migrator) Stop() error {
	return nil
}
