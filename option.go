package neo

import (
	"github.com/bejens/neo/cfg"
	"google.golang.org/grpc/keepalive"
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

// Network is one of 'tcp' 'udp'
func Network(network string) Option {
	return &funcOption{func(o *option) {
		o.network = network
	}}
}

func Address(address string) Option {
	return &funcOption{func(o *option) {
		o.address = address
	}}
}

func fromConfig() []grpc.ServerOption {

	var options []grpc.ServerOption

	writeBufferSize, ok := cfg.Get[int]("server.grpc.write_buffer_size")
	if ok {
		options = append(options, grpc.WriteBufferSize(writeBufferSize))
	}

	readBufferSize, ok := cfg.Get[int]("server.grpc.read_buffer_size")
	if ok {
		options = append(options, grpc.ReadBufferSize(readBufferSize))
	}

	initialWindowSize, ok := cfg.Get[int32]("server.grpc.initial_window_size")
	if ok {
		options = append(options, grpc.InitialWindowSize(initialWindowSize))
	}

	initialConnWindowSize, ok := cfg.Get[int32]("server.grpc.initial_conn_window_size")
	if ok {
		options = append(options, grpc.InitialConnWindowSize(initialConnWindowSize))
	}

	keepaliveParamsMap, ok := cfg.Get[map[string]any]("server.grpc.keepalive_params")
	if ok {
		keepaliveParams := keepalive.ServerParameters{
			MaxConnectionIdle:     0,
			MaxConnectionAge:      0,
			MaxConnectionAgeGrace: 0,
			Time:                  0,
			Timeout:               0,
		}
		maxConnectionIdle, ok := keepaliveParamsMap["max_connection_idle"].(int64)
		if ok {
			keepaliveParams.MaxConnectionIdle = time.Duration(maxConnectionIdle) * time.Second
		}
		maxConnectionAge, ok := keepaliveParamsMap["max_connection_age"].(int64)
		if ok {
			keepaliveParams.MaxConnectionAge = time.Duration(maxConnectionAge) * time.Second
		}
		maxConnectionAgeGrace, ok := keepaliveParamsMap["max_connection_age_grace"].(int64)
		if ok {
			keepaliveParams.MaxConnectionAgeGrace = time.Duration(maxConnectionAgeGrace) * time.Second
		}
		keepaliveTime, ok := keepaliveParamsMap["Time"].(int64)
		if ok {
			keepaliveParams.Time = time.Duration(keepaliveTime) * time.Second
		}
		timeout, ok := keepaliveParamsMap["Timeout"].(int64)
		if ok {
			keepaliveParams.Timeout = time.Duration(timeout) * time.Second
		}
		options = append(options, grpc.KeepaliveParams(keepaliveParams))
	}

	keepaliveEnforcementPolicy, ok := cfg.Get[map[string]any]("server.grpc.keepalive_enforcement_policy")
	if ok {
		enforcementPolicy := keepalive.EnforcementPolicy{
			MinTime:             0,
			PermitWithoutStream: false,
		}

		minTime, ok := keepaliveEnforcementPolicy["min_time"].(int64)
		if ok {
			enforcementPolicy.MinTime = time.Duration(minTime) * time.Second
		}

		permitWithoutStream, ok := keepaliveEnforcementPolicy["permit_without_stream"].(bool)
		if ok {
			enforcementPolicy.PermitWithoutStream = permitWithoutStream
		}
		options = append(options, grpc.KeepaliveEnforcementPolicy(enforcementPolicy))
	}

	maxRecvMsgSize, ok := cfg.Get[int]("server.grpc.max_recv_msg_size")
	if ok {
		options = append(options, grpc.MaxRecvMsgSize(maxRecvMsgSize))
	}

	maxSendMsgSize, ok := cfg.Get[int]("server.grpc.max_send_msg_size")
	if ok {
		options = append(options, grpc.MaxSendMsgSize(maxSendMsgSize))
	}

	maxConcurrentStreams, ok := cfg.Get[uint32]("server.grpc.max_concurrent_streams")
	if ok {
		options = append(options, grpc.MaxConcurrentStreams(maxConcurrentStreams))
	}

	connectionTimeout, ok := cfg.Get[int64]("server.grpc.connection_timeout")
	if ok {
		options = append(options, grpc.ConnectionTimeout(time.Duration(connectionTimeout)*time.Second))
	}

	maxHeaderListSize, ok := cfg.Get[uint32]("server.grpc.max_header_list_size")
	if ok {
		options = append(options, grpc.MaxHeaderListSize(maxHeaderListSize))
	}

	headerTableSize, ok := cfg.Get[uint32]("server.grpc.header_table_size")
	if ok {
		options = append(options, grpc.HeaderTableSize(headerTableSize))
	}

	numStreamWorkers, ok := cfg.Get[uint32]("server.grpc.num_stream_workers")
	if ok {
		options = append(options, grpc.NumStreamWorkers(numStreamWorkers))
	}

	return options
}
