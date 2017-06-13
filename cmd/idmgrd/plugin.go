package main

import "google.golang.org/grpc"

type Plugin interface {
	Register(s *grpc.Server)
}
