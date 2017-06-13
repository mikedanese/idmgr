package main

import (
	"path/filepath"
	"sync"

	api "containers.google.com/idmgr/pkg/api"
	"containers.google.com/idmgr/plugin/gcp"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func newManagementServer() *managementServer {
	return &managementServer{
		mdsStore: map[string]*UnixServer{},

		plugins: []Plugin{
			gcp.New(),
		},
	}
}

type managementServer struct {
	sync.Mutex
	mdsStore map[string]*UnixServer

	plugins []Plugin
}

func (ms *managementServer) CreateMetadata(ctx context.Context, in *api.CreateMetadataRequest) (*api.CreateMetadataResponse, error) {
	ms.Lock()
	defer ms.Unlock()
	if _, ok := ms.mdsStore[in.MountPath]; ok {
		return nil, grpc.Errorf(codes.AlreadyExists, "mds %q already exists", in.MountPath)
	}
	mds, err := NewUnixServer(
		filepath.Join(in.MountPath, "mds.sock"),
		func(s *grpc.Server) {
			for _, plugin := range ms.plugins {
				plugin.Register(s)
			}
		},
		loggingInterceptor,
		serviceAccountInterceptor(in.ServiceAccount),
	)
	if err != nil {
		return nil, err
	}
	ms.mdsStore[in.MountPath] = mds
	go mds.Serve()
	return &api.CreateMetadataResponse{}, nil
}

func (s *managementServer) DestroyMetadata(ctx context.Context, in *api.DestroyMetadataRequest) (*api.DestroyMetadataResponse, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.mdsStore[in.MountPath]; !ok {
		return nil, grpc.Errorf(codes.NotFound, "mds %q not found", in.MountPath)
	}
	s.mdsStore[in.MountPath].Stop()
	delete(s.mdsStore, in.MountPath)
	return &api.DestroyMetadataResponse{}, nil
}
