package backend

import (
	"context"
	"net/http"
	"os"
	"strings"

	utils "./utils"
	jwt "github.com/dgrijalva/jwt-go"
)

func Contains(array []string, element string) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(web http.ResponseWriter, request *http.Request) {
		notAuth := []string{"/api/user/registration", "/api/user/login"}
		if !Contains(notAuth, request.URL.Path) {
			next.ServeHTTP(web, request)
			return
		}
		response := make(map[string]interface{})
		tokenHeader := request.Header.Get("Authorization") //Получение токена

		if tokenHeader == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = utils.Message(false, "Missing auth token")
			web.WriteHeader(http.StatusForbidden)
			web.Header().Add("Content-Type", "application/json")
			utils.Respond(web, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //Токен обычно поставляется в формате `Bearer {token-body}`, мы проверяем, соответствует ли полученный токен этому требованию
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			web.WriteHeader(http.StatusForbidden)
			web.Header().Add("Content-Type", "application/json")
			utils.Respond(web, response)
			return
		}

		tokenPart := splitted[1] //Получаем вторую часть токена
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil || !token.Valid {
			response = utils.Message(false, "Token is not valid")
			web.WriteHeader(http.StatusForbidden)
			web.Header().Add("Content-Type", "application/json")
			utils.Respond(web, response)
			return
		}

		ctx := context.WithValue(request.Context(), "user", tk.UserId)
		request = request.WithContext(ctx)
		next.ServeHTTP(web, request) //middleware done
	})
}
