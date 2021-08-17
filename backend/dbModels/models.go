package dbModels

import "github.com/dgrijalva/jwt-go"

type SearchUser struct {
	Nick string
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Technology struct {
	Id    uint
	Info  string
	Title string
}

type Post struct {
	Id       uint   `json:"id"`
	Code     string `json:"code"`
	Text     string `json:"text"`
	Date     int64  `json:"date"`
	AuthorId uint   `json:"author"`
}

type Likes struct {
	Id     uint
	Status bool
	User   Account
	Post   Post
}

type Sublist struct {
	Id     uint
	Sser   Account
	Sub_to Account
}

type Post_technology struct {
	Id         uint
	Post       Post
	Technology Technology
}
