package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	config            *Config
	context           *sql.DB
	accountRepository *AccountRepository
	postRepository    *PostRepository
}

func New(config *Config) *DataBase {
	return &DataBase{
		config: config,
	}
}

//Account Entity
func (DataBase *DataBase) Account() *AccountRepository {
	if DataBase.accountRepository != nil {
		return DataBase.accountRepository
	}
	DataBase.accountRepository = &AccountRepository{
		db: DataBase,
	}
	return DataBase.accountRepository
}

//Post Entity
func (DataBase *DataBase) Post() *PostRepository {
	if DataBase.postRepository != nil {
		return DataBase.postRepository
	}
	DataBase.postRepository = &PostRepository{
		db: DataBase,
	}
	return DataBase.postRepository
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
