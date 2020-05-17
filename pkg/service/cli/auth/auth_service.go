package auth

type service struct {
	url string
}

func NewService(url string) *service {
	return &service{url: url}
}

func (a *service) Login(userName, password string) (string, string, error) {
	return "", "", nil
}
