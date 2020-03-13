package main

import (
	"github.com/rs/zerolog"
	"golang.org/x/net/context"
	"noframework/api"
)

type handler struct {
	zerolog.Logger
	UserService
}

func NewHandler(logger zerolog.Logger, userService UserService) service.MyServiceServer {
	return handler{
		Logger: logger, UserService: userService,
	}
}

func (h handler) GetUsers(context.Context, *service.GetUsersRequest) (resp *service.GetUsersResponse, err error) {
	defer func() {
		h.Debug().Err(err).Msg("GetUsers")
	}()
	persistenceUsers, err := h.UserService.GetUsers()
	if err != nil {
		return nil, err
	}
	outUsers := make([]string, len(persistenceUsers))
	for i, user := range persistenceUsers {
		outUsers[i] = user.Id
	}
	return &service.GetUsersResponse{Users: outUsers}, nil
}
