package grpc

import (
	"context"
	"time"

	"github.com/patricksferraz/accounting-services/service/common/application/grpc/pb"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/service"
)

type TimeRecordGrpcService struct {
	pb.UnimplementedTimeRecordServiceServer
	TimeRecordService *service.TimeRecordService
	AuthInterceptor   *AuthInterceptor
}

func (t *TimeRecordGrpcService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.Response, error) {
	_time, err := time.Parse(time.RFC3339, in.Time)
	if err != nil {
		return &pb.Response{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	timeRecord, err := t.TimeRecordService.Register(_time, in.Description, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.Response{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.Response{
		Id:     timeRecord.ID,
		Status: "created",
	}, nil
}

func (t *TimeRecordGrpcService) Approve(ctc context.Context, in *pb.ApproveRequest) (*pb.Response, error) {
	timeRecord, err := t.TimeRecordService.Approve(in.Id, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.Response{
			Status: "not updated",
			Error:  err.Error(),
		}, err
	}

	return &pb.Response{
		Id:     timeRecord.ID,
		Status: timeRecord.Status.String(),
	}, nil
}

func (t *TimeRecordGrpcService) Find(ctx context.Context, in *pb.FindRequest) (*pb.TimeRecord, error) {
	timeRecord, err := t.TimeRecordService.Find(in.Id)
	if err != nil {
		return &pb.TimeRecord{}, err
	}

	return &pb.TimeRecord{
		Id:          timeRecord.ID,
		Time:        timeRecord.Time.Format(time.RFC3339),
		Status:      timeRecord.Status.String(),
		Description: timeRecord.Description,
		RegularTime: timeRecord.RegularTime,
		EmployeeId:  timeRecord.EmployeeID,
		ApprovedBy:  timeRecord.ApprovedBy,
		CreatedAt:   timeRecord.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   timeRecord.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (t *TimeRecordGrpcService) FindAllByEmployeeID(in *pb.FindAllByEmployeeIDRequest, stream pb.TimeRecordService_FindAllByEmployeeIDServer) error {
	timeRecords, err := t.TimeRecordService.FindAllByEmployeeID(in.EmployeeId)
	if err != nil {
		return err
	}

	for _, timeRecord := range timeRecords {
		stream.Send(&pb.TimeRecord{
			Id:          timeRecord.ID,
			Time:        timeRecord.Time.Format(time.RFC3339),
			Status:      timeRecord.Status.String(),
			Description: timeRecord.Description,
			RegularTime: timeRecord.RegularTime,
			EmployeeId:  timeRecord.EmployeeID,
			ApprovedBy:  timeRecord.ApprovedBy,
			CreatedAt:   timeRecord.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   timeRecord.UpdatedAt.Format(time.RFC3339),
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
