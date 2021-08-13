package apiserver

import (
	"backend/dbModels"
	"backend/utils"
	"encoding/json"
	"net/http"
)

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

		resp := utils.Message(true, "success")
		resp["data"] = &dbModels.Account{Nick: result.Nick, Id: result.Id}

		utils.Respond(w, resp)
	}
}
