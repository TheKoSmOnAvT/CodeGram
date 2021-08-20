package apiserver

import (
	"backend/dbModels"
	"backend/utils"
	"encoding/json"
	"net/http"
	"time"
)

func (s *APIServer) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := &dbModels.PostСreateModel{}
		err := json.NewDecoder(r.Body).Decode(post)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		post.Date = time.Now().Unix()
		post.AuthorId = r.Context().Value("user").(uint)

		repPost := s.database.Post()
		resultPost, err := repPost.Create(post)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid data or db error"))
			return
		}

		err = repPost.AddTechnologyListToPost(post)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid data or db error"))
			return
		}

		repAccount := s.database.Account()
		resultAccount, err := repAccount.GetUserById(post.AuthorId)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid data or db error"))
			return
		}

		resp := utils.Message(true, "success")
		resp["user"] = resultAccount
		resp["post"] = resultPost

		utils.Respond(w, resp)
	}
}

func (s *APIServer) DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := &dbModels.PostСreateModel{}
		err := json.NewDecoder(r.Body).Decode(post)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		post.AuthorId = r.Context().Value("user").(uint)

		repPost := s.database.Post()

		err = repPost.Delete(post)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		//PRAGMA foreign_keys = off
		err = repPost.DeleteTechnologyListToPost(post)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	}
}

func (s *APIServer) GetTechNology() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repPost := s.database.Post()
		res, err := repPost.TechnologyList()
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		resp := utils.Message(true, "success")
		resp["technology"] = res
		utils.Respond(w, resp)
	}
}

func (s *APIServer) LikePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repPost := s.database.Post()
		like := &dbModels.Likes{}
		err := json.NewDecoder(r.Body).Decode(like)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		err = repPost.LikePost(like)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	}
}

func (s *APIServer) UnlikePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repPost := s.database.Post()
		like := &dbModels.Likes{}
		err := json.NewDecoder(r.Body).Decode(like)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		err = repPost.UnlikePost(like)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	}
}

func (s *APIServer) GetMyFeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset := utils.ConvertNums(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))

		myId := r.Context().Value("user").(uint)
		repPost := s.database.Post()
		posts, err := repPost.GetMyFeed(myId, limit, offset)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		resp := utils.Message(true, "success")
		resp["posts"] = posts
		utils.Respond(w, resp)
	}
}

func (s *APIServer) GetPostsUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset := utils.ConvertNums(r.URL.Query().Get("limit"), r.URL.Query().Get("offset"))

		acc := &dbModels.Account{}
		err := json.NewDecoder(r.Body).Decode(acc)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		repPost := s.database.Post()
		posts, err := repPost.GetPostsUserById(acc.Id, limit, offset)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}
		resp := utils.Message(true, "success")
		resp["posts"] = posts
		utils.Respond(w, resp)
	}
}
