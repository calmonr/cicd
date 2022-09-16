package app

import "github.com/spf13/viper"

type Config struct {
	GRPCServer struct {
		Address, Network string
	}
}

func (c Config) Fill() Config {
	// grpc server
	c.GRPCServer.Address = viper.GetString(PrefixGRPCServer + SuffixAddress)
	c.GRPCServer.Network = viper.GetString(PrefixGRPCServer + SuffixNetwork)

	return c
}
