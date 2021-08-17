package database

import (
	"backend/dbModels"
)

type PostRepository struct {
	db *DataBase
}

func (c *PostRepository) Create(post *dbModels.Post) (*dbModels.Post, error) {
	res, err := c.db.context.Exec("insert into post(author, code, text, date) values ($1,$2);", post.AuthorId, post.Code, post.Text, post.Date)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	post.Id = uint(id)
	return post, nil
}

func (c *PostRepository) Delete(post *dbModels.Post) error {
	_, err := c.db.context.Exec("delete from post where author = $1 and id = $2;", post.AuthorId, post.Id)
	if err != nil {
		return err
	}
	return nil
}
