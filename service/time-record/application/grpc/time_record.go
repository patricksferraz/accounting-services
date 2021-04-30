package grpc

import (
	"context"

	"github.com/patricksferraz/accounting-services/service/common/pb"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TimeRecordGrpcService struct {
	pb.UnimplementedTimeRecordServiceServer
	TimeRecordService *service.TimeRecordService
	AuthInterceptor   *AuthInterceptor
}

func (t *TimeRecordGrpcService) RegisterTimeRecord(ctx context.Context, in *pb.RegisterTimeRecordRequest) (*pb.TimeRecord, error) {
	timeRecord, err := t.TimeRecordService.Register(in.Time.AsTime(), in.Description, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.TimeRecord{}, err
	}

	return &pb.TimeRecord{
		Id:          timeRecord.ID,
		Time:        timestamppb.New(timeRecord.Time),
		Status:      pb.TimeRecord_Status(timeRecord.Status),
		Description: timeRecord.Description,
		RegularTime: timeRecord.RegularTime,
		EmployeeId:  timeRecord.EmployeeID,
		ApprovedBy:  timeRecord.ApprovedBy,
		CreatedAt:   timestamppb.New(timeRecord.CreatedAt),
		UpdatedAt:   timestamppb.New(timeRecord.UpdatedAt),
	}, nil
}

func (t *TimeRecordGrpcService) ApproveTimeRecord(ctc context.Context, in *pb.ApproveTimeRecordRequest) (*pb.ApproveTimeRecordResponse, error) {
	timeRecord, err := t.TimeRecordService.Approve(in.Id, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.ApproveTimeRecordResponse{
			Status: "not updated",
			Error:  err.Error(),
		}, err
	}

	return &pb.ApproveTimeRecordResponse{
		Id:     timeRecord.ID,
		Status: timeRecord.Status.String(),
	}, nil
}

func (t *TimeRecordGrpcService) FindTimeRecord(ctx context.Context, in *pb.FindTimeRecordRequest) (*pb.TimeRecord, error) {
	timeRecord, err := t.TimeRecordService.Find(in.Id)
	if err != nil {
		return &pb.TimeRecord{}, err
	}

	return &pb.TimeRecord{
		Id:          timeRecord.ID,
		Time:        timestamppb.New(timeRecord.Time),
		Status:      pb.TimeRecord_Status(timeRecord.Status),
		Description: timeRecord.Description,
		RegularTime: timeRecord.RegularTime,
		EmployeeId:  timeRecord.EmployeeID,
		ApprovedBy:  timeRecord.ApprovedBy,
		CreatedAt:   timestamppb.New(timeRecord.CreatedAt),
		UpdatedAt:   timestamppb.New(timeRecord.UpdatedAt),
	}, nil
}

func (t *TimeRecordGrpcService) ListTimeRecords(in *pb.ListTimeRecordsRequest, stream pb.TimeRecordService_ListTimeRecordsServer) error {
	timeRecords, err := t.TimeRecordService.FindAllByEmployeeID(in.EmployeeId, in.FromDate.AsTime(), in.ToDate.AsTime())
	if err != nil {
		return err
	}

	for _, timeRecord := range timeRecords {
		stream.Send(&pb.TimeRecord{
			Id:          timeRecord.ID,
			Time:        timestamppb.New(timeRecord.Time),
			Status:      pb.TimeRecord_Status(timeRecord.Status),
			Description: timeRecord.Description,
			RegularTime: timeRecord.RegularTime,
			EmployeeId:  timeRecord.EmployeeID,
			ApprovedBy:  timeRecord.ApprovedBy,
			CreatedAt:   timestamppb.New(timeRecord.CreatedAt),
			UpdatedAt:   timestamppb.New(timeRecord.UpdatedAt),
		})
	}

	return nil
}

func NewTimeRecordGrpcService(service *service.TimeRecordService, authInterceptor *AuthInterceptor) *TimeRecordGrpcService {
	return &TimeRecordGrpcService{
		TimeRecordService: service,
		AuthInterceptor:   authInterceptor,
	}
}
