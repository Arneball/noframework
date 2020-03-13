package main

type Persistence interface {
	GetUsers() ([]User, error)
}

func NewDummyPersistence() Persistence {
	return dummyPersistence{}
}

func (d dummyPersistence) GetUsers() ([]User, error) {
	return []User{{"1337"}}, nil
}

type dummyPersistence struct{}

