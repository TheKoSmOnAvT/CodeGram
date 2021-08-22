package dbModels

import "github.com/dgrijalva/jwt-go"

type SearchUser struct {
	Nick string
}

type Feed struct {
	UserId      uint          `json:"userId"`
	UserNick    string        `json:"userNick"`
	PostId      uint          `json:"postId"`
	Code        string        `json:"code"`
	Text        string        `json:"text"`
	Date        int64         `json:"date"`
	LikeStatus  bool          `json:"likeStatus"`
	Technologys []*Technology `json:"technologys,omitempty"`
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Technology struct {
	Id    uint   `json:"id"`
	Info  string `json:"info"`
	Title string `json:"Title"`
}

type Post struct {
	Id       uint   `json:"id"`
	Code     string `json:"code"`
	Text     string `json:"text"`
	Date     int64  `json:"date"`
	AuthorId uint   `json:"author,omitempty"`
}

type PostСreateModel struct {
	Id          uint   `json:"id"`
	Code        string `json:"code"`
	Text        string `json:"text"`
	Date        int64  `json:"date"`
	AuthorId    uint   `json:"author"`
	Technologys []uint `json:"technologys"`
}

type Likes struct {
	Id   uint `json:"id"`
	User uint `json:"user"`
	Post uint `json:"post"`
}

type Sublist struct {
	Id     uint `json:"id"`
	User   uint `json:"user"`
	Sub_to uint `json:"subTo"`
}

type Post_technology struct {
	Id         uint
	Post       Post
	Technology Technology
}
