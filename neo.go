package neo

import (
	"github.com/bejens/neo/logx"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"

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
			logx.Info("Server Graceful Stop")
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

func (neo *Neo) Stop() {
	neo.server.GracefulStop()
}

func (neo *Neo) GrpcServer() *grpc.Server {
	return neo.server
}

func New(options ...Option) *Neo {
	opts := defaultOptions
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &Neo{opt: opts}
}
