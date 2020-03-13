package main

type UserService interface {
	GetUsers() ([]User, error)
}

type ourUserService struct {
	Persistence
}

func NewUserService(persistence Persistence) UserService {
	return ourUserService{persistence}
}

