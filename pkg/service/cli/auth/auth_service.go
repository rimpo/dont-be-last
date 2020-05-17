package auth

import "fmt"

type service struct {
	url string
}

func NewService(url string) *service {
	return &service{url: url}
}

func (a *service) Login(userName, password string) (string, int, error) {
	fmt.Println("received login")
	return "abc", 1, nil
}
