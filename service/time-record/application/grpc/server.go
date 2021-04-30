package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/patricksferraz/accounting-services/service/common/pb"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/service"
	"github.com/patricksferraz/accounting-services/service/time-record/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/mgo.v2"
)

func StartGrpcServer(database *mgo.Database, _service pb.AuthServiceClient, port int) {

	authService := service.NewAuthService(_service)
	interceptor := NewAuthInterceptor(authService)
	timeRecordRepository := repository.NewTimeRecordRepositoryDb(database)
	timeRecordService := service.NewTimeRecordService(timeRecordRepository)
	timeRecordGrpcService := NewTimeRecordGrpcService(timeRecordService, interceptor)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	reflection.Register(grpcServer)
	pb.RegisterTimeRecordServiceServer(grpcServer, timeRecordGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
