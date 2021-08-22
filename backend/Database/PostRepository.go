package database

import (
	"backend/dbModels"
)

type PostRepository struct {
	db *DataBase
}

func (c *PostRepository) Create(post *dbModels.PostСreateModel) (*dbModels.PostСreateModel, error) {
	res, err := c.db.context.Exec("insert into post(author, code, text, date) values ($1,$2,$3,$4);", post.AuthorId, post.Code, post.Text, post.Date)
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

func (c *PostRepository) CheckToDelete(post *dbModels.PostСreateModel) error {
	postCheck := &dbModels.Post{}
	if err := c.db.context.QueryRow("select Id from post where author = $1 and id = $2;", post.AuthorId, post.Id).Scan(&postCheck.Id); err != nil {
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

func (c *PostRepository) GetMyFeed(myId uint, limit uint, offset uint) ([]*dbModels.Feed, error) {
	feeds := make([]*dbModels.Feed, 0)
	postsRows, err := c.db.context.Query("select id, code, text, date, author from post where author in (select sub_to from sublist where user = $1)  order by date desc limit $2 offset $3", myId, limit, offset)
	if err != nil {
		return nil, err
	}

	for postsRows.Next() {
		post := new(dbModels.Post)
		if err := postsRows.Scan(&post.Id, &post.Code, &post.Text, &post.Date, &post.AuthorId); err != nil {
			return nil, err
		}

		techs := make([]*dbModels.Technology, 0)
		techRows, err := c.db.context.Query("select id, title, info from technology where id in(select technology from post_technology where post = $1)", post.Id)
		if err != nil {
			return nil, err
		}
		for techRows.Next() {
			tech := new(dbModels.Technology)
			if err := techRows.Scan(&tech.Id, &tech.Title, &tech.Info); err != nil {
				return nil, err
			}

			techs = append(techs, tech)
		}

		acc := &dbModels.Account{}
		accountResult, err := c.db.context.Query("select id, nick from account where id = $1", post.AuthorId)
		if err != nil {
			return nil, err
		}
		for accountResult.Next() {
			if err := accountResult.Scan(&acc.Id, &acc.Nick); err != nil {
				return nil, err
			}
		}

		likeResult, err := c.db.context.Query("select count(*) from likes where post = $1 and user = $2", post.Id, myId)
		if err != nil {
			return nil, err
		}

		likeStatus := false
		count := 0
		for likeResult.Next() {
			if err := likeResult.Scan(&count); err != nil {
				return nil, err
			}
		}
		if count > 0 {
			likeStatus = true
		}

		feed := &dbModels.Feed{UserId: post.AuthorId, Technologys: techs, PostId: post.Id, Code: post.Code, Text: post.Text, Date: post.Date, UserNick: acc.Nick, LikeStatus: likeStatus}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}

func (c *PostRepository) GetPostsUserById(myId uint, userId uint, limit uint, offset uint) ([]*dbModels.Feed, error) {
	feeds := make([]*dbModels.Feed, 0)
	postsRows, err := c.db.context.Query("select id, code, text, date from post where author = $1 order by date desc limit $2 offset $3", userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for postsRows.Next() {
		post := new(dbModels.Post)
		if err := postsRows.Scan(&post.Id, &post.Code, &post.Text, &post.Date); err != nil {
			return nil, err
		}

		techs := make([]*dbModels.Technology, 0)
		techRows, err := c.db.context.Query("select id, title, info from technology where id in(select technology from post_technology where post = $1)", post.Id)
		if err != nil {
			return nil, err
		}
		for techRows.Next() {
			tech := new(dbModels.Technology)
			if err := techRows.Scan(&tech.Id, &tech.Title, &tech.Info); err != nil {
				return nil, err
			}
			techs = append(techs, tech)
		}

		acc := &dbModels.Account{}
		if err := c.db.context.QueryRow("select id, nick from account where id = $1", userId).Scan(&acc.Id, &acc.Nick); err != nil {
			return nil, err
		}

		likeResult, err := c.db.context.Query("select count(*) from likes where post = $1 and user = $2", post.Id, myId)
		if err != nil {
			return nil, err
		}

		likeStatus := false
		count := 0
		for likeResult.Next() {
			if err := likeResult.Scan(&count); err != nil {
				return nil, err
			}
		}
		if count > 0 {
			likeStatus = true
		}

		feed := &dbModels.Feed{UserId: userId, Technologys: techs, PostId: post.Id, Code: post.Code, Text: post.Text, Date: post.Date, UserNick: acc.Nick, LikeStatus: likeStatus}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}
