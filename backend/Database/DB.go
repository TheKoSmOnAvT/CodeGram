package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	config            *Config
	context           *sql.DB
	accountRepository *AccountRepository
}

func New(config *Config) *DataBase {
	return &DataBase{
		config: config,
	}
}

func (DataBase *DataBase) Account() *AccountRepository {
	if DataBase.accountRepository != nil {
		return DataBase.accountRepository
	}
	DataBase.accountRepository = &AccountRepository{
		db: DataBase,
	}
	return DataBase.accountRepository
}

func (DataBase *DataBase) Open() error {
	context, err := sql.Open("sqlite3", DataBase.config.DataBaseURL)
	if err != nil {
		return err
	}
	if err := context.Ping(); err != nil {
		return err
	}
	DataBase.context = context
	return nil
}

func (DataBase *DataBase) Close() {
	DataBase.context.Close()
}
