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
		post := &dbModels.Post{}
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
		post := &dbModels.Post{}
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

		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	}
}
