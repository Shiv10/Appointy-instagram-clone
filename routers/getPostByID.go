package routers

import (
	"fmt"
)

func GetPost() string{
	fmt.Println("Get a post")
	return "post sent."
}