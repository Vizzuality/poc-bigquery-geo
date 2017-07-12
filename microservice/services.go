package microservice

import "log"

// User ...
type User struct {
	Name string
}

func getUserService(name string) User {
	log.Println("[SERVICE] Getting user")
	return User{
		Name: name,
	}
}
