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
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		rep := s.database.Account()
		result, err := rep.FindByNick(data.Nick)
		if err != nil {
			fmt.Print(err)
			utils.Respond(w, utils.Message(false, "Invalid data or db error"))
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
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		rep := s.database.Account()
		result, err := rep.Create(account)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid data or db error"))
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
			utils.Respond(w, utils.Message(false, "Invalid request"))
			return
		}

		rep := s.database.Account()
		result, err := rep.Login(account)
		if err != nil {
			utils.Respond(w, utils.Message(false, "Invalid data or db error"))
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
