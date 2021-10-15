package shhttp

import (
	shards "github.com/pg-sharding/spqr/router/protos"
	"google.golang.org/grpc"
)

func Register(server *grpc.Server) {
	shards.RegisterQueryServiceServer(server, &Routerserver{})

	shards.RegisterShardServiceServer(server, &Spqrserver{})
}
