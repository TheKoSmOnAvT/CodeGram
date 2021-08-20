package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(web http.ResponseWriter, data map[string]interface{}) {
	web.Header().Add("Content-Type", "application/json")
	json.NewEncoder(web).Encode(data)
}

func CreateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ConvertNums(limitStr string, offsetStr string) (uint, uint) {
	offset, _ := strconv.ParseUint(offsetStr, 10, 64)
	limit, _ := strconv.ParseUint(limitStr, 10, 64)

	if offset <= 0 {
		offset = 1
	}
	if limit <= 0 {
		limit = 1
	}
	offset = (offset - 1) * limit

	return uint(limit), uint(offset)
}
