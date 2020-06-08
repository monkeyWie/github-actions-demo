package main

import (
	"fmt"
	"github-action-demo/internal/service"
)

func main() {
	user := &service.User{Name: "hello", Age: 18}
	fmt.Println(user.GetName())
}
