package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/patricksferraz/accounting-services/services/auth/application/grpc/pb"
	"github.com/patricksferraz/accounting-services/services/auth/domain/service"
	"github.com/patricksferraz/accounting-services/services/auth/infrastructure/db"
	"github.com/patricksferraz/accounting-services/services/auth/infrastructure/repository"
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
