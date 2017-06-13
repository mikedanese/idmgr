package main

import (
	"flag"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	api "containers.google.com/idmgr/pkg/api"
	_ "containers.google.com/idmgr/pkg/logger"
)

func main() {
	flag.Parse()
	glog.Infof("Starting")

	s, err := NewUnixServer(
		"/tmp/idmgr.sock",
		func(s *grpc.Server) {
			api.RegisterManagementServer(s, newManagementServer())
		},
		loggingInterceptor,
	)
	if err != nil {
		glog.Fatalf("couldn't start server: %v", err)
	}
	glog.Fatalf("Server exited unexpectedly: %v", s.Serve())
}
