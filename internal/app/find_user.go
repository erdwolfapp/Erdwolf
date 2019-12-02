package app

import "fmt"

func (a *Application) FindUserById (id int) User {
	/* TODO: */fmt.Println("fixme: Implement FindUserById(int)")
	return User {
		Identifier: 0,
		Name: "unknown",
		AuthDomain: "none",
	}
}