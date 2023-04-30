package neo

import (
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/bejens/neo/cfg"

	"github.com/bejens/neo/logx"

	"google.golang.org/grpc"
)

type Neo struct {
	opt option

	serviceInfos sync.Map

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

	neo.serviceInfos.Range(func(key, value any) bool {
		sd := key.(*grpc.ServiceDesc)
		logx.Info("Register Service Success11", logx.String("service_name", sd.ServiceName))
		neo.server.RegisterService(sd, value)
		return true
	})

	lis, err := net.Listen(neo.opt.network, neo.opt.address)
	if err != nil {
		return err
	}

	logx.Info("Server Starting...",
		logx.String("address", neo.opt.address))
	return neo.server.Serve(lis)
}

func (neo *Neo) Register(sd *grpc.ServiceDesc, ss interface{}) {
	neo.serviceInfos.Store(sd, ss)
}

func (neo *Neo) GrpcServer() *grpc.Server {
	return neo.server
}

func (neo *Neo) Stop() {
	logx.Info("Server Graceful Stop")
	if err := logx.Sync(); err != nil {
		logx.Warn("Close Logger Fail", logx.Err(err))
	}
	neo.server.GracefulStop()
}

func New(options ...Option) (*Neo, error) {

	if err := cfg.InitCfg(); err != nil {
		logx.Error("Init Config Error", logx.Err(err))
		return nil, err
	}

	opts := defaultOptions
	grpcOpts, neoOpts := fromConfig()

	for _, opt := range options {
		opt.apply(&opts)
	}
	for _, opt := range grpcOpts {
		GrpcOption(opt).apply(&opts)
	}
	for _, opt := range neoOpts {
		opt.apply(&opts)
	}

	return &Neo{opt: opts}, nil
}
