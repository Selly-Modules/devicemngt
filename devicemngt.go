package devicemngt

import (
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Selly-Modules/postgresql"
	"github.com/jmoiron/sqlx"
)

// PostgreSQLConfig ...
type PostgreSQLConfig struct {
	Host, User, Password, DBName, Port, SSLMode string
}

// Config ...
type Config struct {
	// PostgreSQL config, for save documents
	PostgreSQL PostgreSQLConfig
}

// Service ...
type Service struct {
	Config
	DB      *sqlx.DB
	Builder squirrel.StatementBuilderType
}

var s *Service

// NewInstance ...
func NewInstance(config Config) error {
	if config.PostgreSQL.Host == "" {
		return errors.New("please provide all necessary information: source, postgresql")
	}

	// Connect PG
	err := postgresql.Connect(
		config.PostgreSQL.Host,
		config.PostgreSQL.User,
		config.PostgreSQL.Password,
		config.PostgreSQL.DBName,
		config.PostgreSQL.Port,
		config.PostgreSQL.SSLMode,
	)
	if err != nil {
		fmt.Println("Cannot init module DEVICE MANAGEMENT", err)
		return err
	}

	s = &Service{
		Config:  config,
		DB:      postgresql.GetSqlxInstance(),
		Builder: postgresql.GetStmBuilder(),
	}

	// Create schema
	schemaContent := fmt.Sprintf(`
		%s
		
  `,
		DeviceManagementSchema,
	)
	if _, err = s.DB.MustExec(schemaContent).RowsAffected(); err != nil {
		panic(err)
	}

	// TODO: Index db

	return nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}
