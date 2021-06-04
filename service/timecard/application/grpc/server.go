package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/patricksferraz/timecard-service/application/grpc/pb"
	"github.com/patricksferraz/timecard-service/domain/service"
	"github.com/patricksferraz/timecard-service/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/mgo.v2"
)

func StartGrpcServer(database *mgo.Database, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	timecardRepository := repository.TimecardRepositoryDb{Db: database}
	timeRecordRepository := repository.TimeRecordRepositoryDb{Db: database}
	timecardService := service.TimecardService{
		TimecardRepository:   &timecardRepository,
		TimeRecordRepository: &timeRecordRepository,
	}
	timecardGrpcService := NewTimecardGrpcService(timecardService)
	pb.RegisterTimecardServiceServer(grpcServer, timecardGrpcService)

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
