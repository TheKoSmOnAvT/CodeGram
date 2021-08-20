package apiserver

import (
	"backend/dbModels"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func (s *APIServer) Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &dbModels.SearchUser{}
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		rep := s.database.Account()
		result, err := rep.FindByNick(data.Nick)
		if err != nil {
			fmt.Print(err)
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		resp := utils.Message(true, "success")
		resp["data"] = result

		utils.Respond(w, resp)
	}
}

func (s *APIServer) Registration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account := &dbModels.Account{}
		err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		rep := s.database.Account()
		result, err := rep.Create(account)
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		//Create JWT token
		tk := &dbModels.Token{UserId: result.Id}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

		resp := utils.Message(true, "success")
		resp["data"] = &dbModels.Account{Nick: result.Nick, Id: result.Id, Token: tokenString}

		utils.Respond(w, resp)
	}
}

func (s *APIServer) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account := &dbModels.Account{}
		err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		rep := s.database.Account()
		result, err := rep.Login(account)
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		//Create JWT token
		tk := &dbModels.Token{UserId: result.Id}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

		resp := utils.Message(true, "success")
		resp["data"] = &dbModels.Account{Nick: result.Nick, Id: result.Id, Token: tokenString}

		utils.Respond(w, resp)
	}
}

func (s *APIServer) GetUserSubscribes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account := &dbModels.Account{}
		err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}

		rep := s.database.Account()
		subs, err := rep.GetSubscribes(account.Id)
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}
		resp := utils.Message(true, "success")
		resp["subs"] = subs

		utils.Respond(w, resp)
	}
}

func (s *APIServer) GetMySubscribes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		myId := r.Context().Value("user").(uint)
		rep := s.database.Account()
		subs, err := rep.GetSubscribes(myId)
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}
		resp := utils.Message(true, "success")
		resp["subs"] = subs

		utils.Respond(w, resp)
	}
}

func (s *APIServer) Subscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := &dbModels.Sublist{}
		err := json.NewDecoder(r.Body).Decode(sub)

		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}
		sub.User = r.Context().Value("user").(uint)

		rep := s.database.Account()
		err = rep.Subscribe(sub)
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	}
}

func (s *APIServer) Unsubscribe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sub := &dbModels.Sublist{}
		err := json.NewDecoder(r.Body).Decode(sub)

		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}
		sub.User = r.Context().Value("user").(uint)

		rep := s.database.Account()
		err = rep.Unsubscribe(sub)
		if err != nil {
			utils.Respond(w, utils.Message(false, string(err.Error())))
			return
		}
		resp := utils.Message(true, "success")
		utils.Respond(w, resp)
	}
}
