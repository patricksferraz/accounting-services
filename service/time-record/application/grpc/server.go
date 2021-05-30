package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/c4ut/accounting-services/service/common/logger"
	"github.com/c4ut/accounting-services/service/common/pb"
	"github.com/c4ut/accounting-services/service/time-record/domain/service"
	"github.com/c4ut/accounting-services/service/time-record/infrastructure/db"
	"github.com/c4ut/accounting-services/service/time-record/infrastructure/repository"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *db.Mongo, _service pb.AuthServiceClient, port int) {

	authService := service.NewAuthService(_service)
	interceptor := NewAuthInterceptor(authService)
	timeRecordRepository := repository.NewTimeRecordRepository(database)
	timeRecordService := service.NewTimeRecordService(timeRecordRepository)
	timeRecordGrpcService := NewTimeRecordGrpcService(timeRecordService, interceptor)

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
			interceptor.Unary(),
		),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	reflection.Register(grpcServer)
	pb.RegisterTimeRecordServiceServer(grpcServer, timeRecordGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Log.Fatal("cannot start grpc server", err)
	}
}
