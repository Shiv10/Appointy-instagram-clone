package routers

import (
	"fmt"
)

func GetUser() string{
	fmt.Println("Get a user")
	return "User sent."
}