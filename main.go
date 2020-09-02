package main

import (
	"fmt"
	"go/demo/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Matt",
		LastName:  "Sage",
	}
	fmt.Println(u)
}
