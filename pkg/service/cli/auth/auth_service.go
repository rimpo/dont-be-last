package auth

import "fmt"

type service struct {
	url string
}

func NewService(url string) *service {
	return &service{url: url}
}

func (a *service) Login(userName, password string) (string, int, error) {
	if userName == "m" annd password == "b" {
		return "abc-token", 1, nil
	}
	return "", 0, fmt.Errorf("failed to login")
}
