package database

import (
	"backend/dbModels"

	"golang.org/x/crypto/bcrypt"
)

type AccountRepository struct {
	db *DataBase
}

//registration user acc
func (c *AccountRepository) Create(acc *dbModels.Account) (*dbModels.Account, error) {
	err := acc.CreateHash()
	if err != nil {
		return nil, err
	}
	res, err := c.db.context.Exec("insert into account(nick, hashpassword) values ($1,$2);", acc.Nick, acc.Hashpassword)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	acc.Id = uint(id)
	return acc, nil
}

// find user in db by nickname
func (c *AccountRepository) FindByNick(word string) ([]*dbModels.Account, error) {
	acc := make([]*dbModels.Account, 0)
	word = "%" + word + "%"
	rows, err := c.db.context.Query("select id, nick from account where nick like $1", word)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		account := new(dbModels.Account)
		if err := rows.Scan(&account.Id, &account.Nick); err != nil {
			panic(err)
		}
		acc = append(acc, account)
	}

	return acc, nil
}

// login in account
func (c *AccountRepository) Login(acc *dbModels.Account) (*dbModels.Account, error) {
	searchedAcc := &dbModels.Account{}
	if err := c.db.context.QueryRow("select id, nick, hashpassword from account where nick like $1", &acc.Nick).Scan(&searchedAcc.Id, &searchedAcc.Nick, &searchedAcc.Hashpassword); err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(searchedAcc.Hashpassword), []byte(acc.Hashpassword))
	if err != nil {
		return nil, err
	}
	return searchedAcc, nil
}
