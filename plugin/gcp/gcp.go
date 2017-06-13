package gcp

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	api "containers.google.com/idmgr/plugin/gcp/api"
)

func New() *metadataServer {
	return &metadataServer{}
}

func (mds *metadataServer) Register(s *grpc.Server) {
	api.RegisterMetadataServer(s, mds)
}

type metadataServer struct{}

func (s *metadataServer) GetToken(ctx context.Context, in *api.GetTokenRequest) (*api.GetTokenResponse, error) {
	return &api.GetTokenResponse{}, nil
}
