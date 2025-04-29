package postgres

import (
	"backend/pkg/infra/storage/db"
	"backend/pkg/infra/storage/db/dbimpl"
	"backend/pkg/infra/storage/migrator"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type DB interface {
	db.DB
}

type Migrator interface {
	AddMigration(mg *migrator.Migrator)
}

type postgresDB struct {
	engine     *xorm.Engine
	migrations Migrator
	dialect    migrator.Dialect
	log        *zap.Logger
	db.DB
}

func New(migrations Migrator, connection string) (DB, error) {
	p := &postgresDB{
		log: zap.L().Named("postgres"),
	}

	engine, err := xorm.NewEngine("postgres", connection)
	if err != nil {
		return nil, err
	}

	p.engine = engine
	p.migrations = migrations
	p.dialect = migrator.NewDialect(p.engine)
	engine.SetTZDatabase(time.UTC)

	err = p.Migrate()
	if err != nil {
		p.log.Error("migration failed err: %v", zap.Any("errors", err))
		return nil, err
	}

	p.DB = dbimpl.NewSQL(sqlx.NewDb(p.engine.DB().DB, p.GetDialect().DriverName()))
	return p, nil
}

func (p *postgresDB) Migrate() error {
	migrator := migrator.NewMigrator(p.engine)
	p.migrations.AddMigration(migrator)
	return migrator.Start()
}

func (p *postgresDB) GetDialect() migrator.Dialect {
	return p.dialect
}
