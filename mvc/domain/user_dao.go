package domain

import (
	"errors"
	"fmt"
)

// temp database
var (
	users = map[int64]*User{
		// "123": &User{Id:1, FirstName:"Bruce", LastName:"Wayne", Email:"batman@email.com"},
		123: {Id: 1, FirstName: "Bruce", LastName: "Wayne", Email: "batman@email.com"},
	}
)

func GetUser(userID int64) (*User, error) {

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("user %v was not found", userID))

	// Another way
	// get user from users where id is userID
	// user := users[userID]
	// if user == nil {
	// 	//return User{}, errors.New(fmt.Sprintf("user %v was not found", userID))
	// 	return nil, errors.New(fmt.Sprintf("user %v was not found", userID))
	// }
}
