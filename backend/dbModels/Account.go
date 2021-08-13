package dbModels

import "backend/utils"

type Account struct {
	Id           uint   `json:"id"`
	Nick         string `json:"nick"`
	Hashpassword string `json:"password,omitempty"`
}

func (acc *Account) CreateHash() error {
	if len(acc.Hashpassword) > 0 {
		hash, err := utils.CreateHash(acc.Hashpassword)
		if err != nil {
			return err
		}
		acc.Hashpassword = string(hash)
	}
	return nil
}
