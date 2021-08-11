package database

import "backend/dbModels"

type AccountRepository struct {
	db *DataBase
}

func (c *AccountRepository) Create(acc *dbModels.Account) (*dbModels.Account, error) {
	if err := c.db.context.QueryRow("insert into account(nick, hashpassword) values ($1,$2) returing id", acc.Nick, acc.Hashpassword).Scan(&acc.Id); err != nil {
		return nil, err
	}
	return acc, nil
}

func (c *AccountRepository) FindByNick(word string) (*dbModels.Account, error) {
	acc := &dbModels.Account{}

	if err := c.db.context.QueryRow("select id, nick, hashpassword from account where nick like $1", word).Scan(&acc.Id, &acc.Nick, &acc.Hashpassword); err != nil {
		return nil, err
	}
	return acc, nil
}
