package dbModels

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	nick         string
	hashpassword uint
}

type Technology struct {
	gorm.Model
	info  string
	title string
}

type Post struct {
	gorm.Model
	code   string
	text   string
	date   uint64
	author Account
}

type Likes struct {
	gorm.Model
	status bool
	user   Account
	post   Post
}

type Sublist struct {
	gorm.Model
	user   Account
	sub_to Account
}

type Post_technology struct {
	gorm.Model
	post       Post
	technology Technology
}
