package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/patricksferraz/accounting-services/service/auth/domain/service"
	"github.com/patricksferraz/accounting-services/service/auth/infrastructure/db"
	"github.com/patricksferraz/accounting-services/service/auth/infrastructure/repository"
	"github.com/patricksferraz/accounting-services/service/common/application/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *db.Keycloak, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	authRepository := repository.AuthRepositoryDb{Db: database}
	authService := service.AuthService{
		AuthRepository: &authRepository,
	}
	authGrpcService := NewAuthGrpcService(authService)
	pb.RegisterAuthServiceServer(grpcServer, authGrpcService)

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
