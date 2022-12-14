package app

import (
	"github.com/spf13/pflag"
)

const (
	PrefixGRPCServer = "grpc-server."
)

const (
	SuffixAddress = "address"
	SuffixNetwork = "network"
)

func Flags() *pflag.FlagSet {
	f := new(pflag.FlagSet)

	// grpc server
	f.String(PrefixGRPCServer+SuffixAddress, ":50051", "gRPC server address to listen on")
	f.String(PrefixGRPCServer+SuffixNetwork, "tcp", "gRPC server network to listen on")

	return f
}
