package grpcserverc

import (
	"flag"
	"fmt"
	"net"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	sctx "jetshop/service-context"
	"jetshop/service-context/component/grpcserverc/logging"
	"jetshop/service-context/component/grpcserverc/recovermiddleware"
)

type GrpcComponent interface {
	SetRegisterHdl(hdl func(*grpc.Server))
}

type grpcServer struct {
	id          string
	prefix      string
	port        int
	server      *grpc.Server
	logger      sctx.Logger
	registerHdl func(*grpc.Server)
}

func NewGrpcServer(id string) *grpcServer {
	return &grpcServer{id: id}
}

func (s *grpcServer) ID() string {
	return s.id
}

func (s *grpcServer) InitFlags() {
	flag.IntVar(&s.port, "grpc_port", 50051, "Port of gRPC service")
}

func (s *grpcServer) Recover() {
	if err := recover(); err != nil {
		s.logger.Error("recover error", err)
	}
}

func (s *grpcServer) Activate(sc sctx.ServiceContext) error {
	go func() {
		defer s.Recover()

		s.logger = sc.Logger(s.id)

		s.logger.Infoln("Setup gRPC service:", s.prefix)
		s.server = grpc.NewServer(
			grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
				otelgrpc.StreamServerInterceptor(),
				recovermiddleware.StreamServerInterceptor(),
				logging.StreamServerInterceptor(s.logger),
			)),
			grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
				otelgrpc.UnaryServerInterceptor(),
				recovermiddleware.UnaryServerInterceptor(),
				logging.UnaryServerInterceptor(s.logger),
			)),
		)

		if s.registerHdl != nil {
			s.logger.Infoln("registering services...")
			s.registerHdl(s.server)
		}

		address := fmt.Sprintf("0.0.0.0:%d", s.port)
		lis, err := net.Listen("tcp", address)

		if err != nil {
			s.logger.Errorln("Error %v", err)
		}

		_ = s.server.Serve(lis)
	}()

	return nil
}

func (s *grpcServer) Stop() error {
	s.server.Stop()

	return nil
}

func (s *grpcServer) SetRegisterHdl(hdl func(*grpc.Server)) {
	s.registerHdl = hdl
}
