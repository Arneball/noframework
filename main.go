//go:generate protoc --go_out=plugins=grpc:api --grpc-gateway_out=logtostderr=true:api service.proto

package main

import (
	"bufio"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"gopkg.in/errgo.v2/fmt/errors"
	"net"
	service "noframework/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := grpc.NewServer()
	logger := zerolog.New(bufio.NewWriter(os.Stdout)).With().CallerWithSkipFrameCount(2).Logger()

	persistence := NewDummyPersistence()
	userService := NewUserService(persistence)
	handler := NewHandler(logger,userService)

	service.RegisterMyServiceServer(server, handler)
	listen, err := net.Listen("tcp", ":8080")

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <- interrupt
		logger.Err(errors.Newf("Signal received %s", sig))
	}()
	if err != nil {
		panic(err)
	}
	panic(server.Serve(listen))
}
