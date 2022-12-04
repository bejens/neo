package neo

import (
	"google.golang.org/grpc"
	"net"
)

type Neo struct {
	opt option

	server *grpc.Server
}

func (neo *Neo) Run() error {

	server := grpc.NewServer(neo.opt.grpcServerOpts...)
	lis, err := net.Listen(neo.opt.network, neo.opt.address)
	if err != nil {
		return err
	}

	return server.Serve(lis)
}

func (neo *Neo) GrpcServer() *grpc.Server {
	return neo.server
}

func New(options ...Option) (*Neo, error) {
	opts := defaultOptions
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &Neo{opt: opts}, nil
}
