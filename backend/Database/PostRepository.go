package database

import (
	"backend/dbModels"
)

type PostRepository struct {
	db *DataBase
}

func (c *PostRepository) Create(post *dbModels.PostСreateModel) (*dbModels.PostСreateModel, error) {
	res, err := c.db.context.Exec("insert into post(author, code, text, date) values ($1,$2);", post.AuthorId, post.Code, post.Text, post.Date)
	if err != nil {
		return nil, err
	}
	id, _ := res.LastInsertId()
	post.Id = uint(id)
	return post, nil
}

func (c *PostRepository) AddTechnologyListToPost(post *dbModels.PostСreateModel) error {
	for _, techId := range post.Technologys {
		_, err := c.db.context.Exec("insert into post_technology(post, technology) values ($1,$2);", post.Id, techId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *PostRepository) DeleteTechnologyListToPost(post *dbModels.PostСreateModel) error {
	_, err := c.db.context.Exec("delete from post_technology where post = $1;", post.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *PostRepository) Delete(post *dbModels.PostСreateModel) error {
	_, err := c.db.context.Exec("delete from post where author = $1 and id = $2;", post.AuthorId, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *PostRepository) TechnologyList() ([]*dbModels.Technology, error) {
	techs := make([]*dbModels.Technology, 0)

	rows, err := c.db.context.Query("select id, title, info from technology")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		technology := new(dbModels.Technology)
		if err := rows.Scan(&technology.Id, &technology.Title, &technology.Info); err != nil {
			return nil, err
		}
		techs = append(techs, technology)
	}

	return techs, nil
}

func (c *PostRepository) LikePost(likeModel *dbModels.Likes) error {
	_, err := c.db.context.Exec("insert into likes(post, user) values ($1,$2);", likeModel.Post, likeModel.User)
	if err != nil {
		return err
	}
	return nil
}

func (c *PostRepository) UnlikePost(likeModel *dbModels.Likes) error {
	_, err := c.db.context.Exec("delete from likes where post = $1 and user = $2;", likeModel.Post, likeModel.User)
	if err != nil {
		return err
	}
	return nil
}
