package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/annbelievable/go-microservice-gokit/endpoint"
	"github.com/annbelievable/go-microservice-gokit/proto"
	"github.com/annbelievable/go-microservice-gokit/service"
	"github.com/annbelievable/go-microservice-gokit/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewContactService(logger)
	endpoint := endpoint.MakeContactEndpoints(service)
	grpcServer := transport.NewGrpcServer(endpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":8000")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		reflection.Register(baseServer)
		proto.RegisterContactServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started.")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
