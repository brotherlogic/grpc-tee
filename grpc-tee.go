package main

import (
	"log"
	"net"

	"github.com/mwitkow/grpc-proxy/proxy"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	port = ":50042"
)

func tee(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	return ctx, nil, grpc.Errorf(codes.Unimplemented, "We can't handle that")
}

func main() {
	var director proxy.StreamDirector
	director = tee
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to get port: %v", err)
	}
	server := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
	)
	proxy.RegisterService(server, director, "discovery.DiscoveryService", "Discover")
	server.Serve(lis)
}
