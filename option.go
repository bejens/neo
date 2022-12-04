package neo

import (
	"time"

	"google.golang.org/grpc"
)

type Option interface {
	apply(*option)
}

type funcOption struct {
	f func(*option)
}

func (fo *funcOption) apply(opt *option) {
	fo.f(opt)
}

type grpcFuncOption struct {
	f grpc.ServerOption
}

func (gfo *grpcFuncOption) apply(opt *option) {
	opt.grpcServerOpts = append(opt.grpcServerOpts, gfo.f)
}

func GrpcOption(opt grpc.ServerOption) Option {
	return &grpcFuncOption{
		f: opt,
	}
}

type option struct {
	network string
	address string

	grpcServerOpts []grpc.ServerOption
}

var defaultOptions = option{
	network: "tcp",
	address: ":3000",

	grpcServerOpts: defaultGrpcOptions,
}

var defaultGrpcOptions = []grpc.ServerOption{
	grpc.ConnectionTimeout(30 * time.Second),
}
