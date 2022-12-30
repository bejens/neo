package neo

import (
	"github.com/bejens/neo/cfg"
	"net"
	"os"
	"os/signal"

	"github.com/bejens/neo/logx"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Neo struct {
	opt option

	server *grpc.Server
}

func (neo *Neo) Run() error {

	//Graceful Stop
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, os.Kill)
	go func() {
		select {
		case <-sign:
			neo.Stop()
		}
	}()

	neo.server = grpc.NewServer(neo.opt.grpcServerOpts...)
	lis, err := net.Listen(neo.opt.network, neo.opt.address)
	if err != nil {
		return err
	}

	logx.Info("Server Starting...",
		zap.String("address", neo.opt.address))
	return neo.server.Serve(lis)
}

func (neo *Neo) Register(sd *grpc.ServiceDesc, ss interface{}) {
	neo.server.RegisterService(sd, ss)
}

func (neo *Neo) GrpcServer() *grpc.Server {
	return neo.server
}

func (neo *Neo) Stop() {
	logx.Info("Server Graceful Stop")
	neo.server.GracefulStop()
}

func New(options ...Option) (*Neo, error) {

	if err := cfg.InitCfg(); err != nil {
		return nil, err
	}
	grpcOpts := fromConfig()

	opts := defaultOptions
	for _, opt := range options {
		opt.apply(&opts)
	}
	opts.grpcServerOpts = append(opts.grpcServerOpts, grpcOpts...)

	return &Neo{opt: opts}, nil
}
