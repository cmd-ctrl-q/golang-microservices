package services

import (
	"https://github.com/cmd-ctrl-q/golang-microservices/mvc/domain"
)
func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}