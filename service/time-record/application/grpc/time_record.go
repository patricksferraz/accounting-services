package grpc

import (
	"context"

	"github.com/c4ut/accounting-services/service/common/pb"
	"github.com/c4ut/accounting-services/service/time-record/domain/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TimeRecordGrpcService struct {
	pb.UnimplementedTimeRecordServiceServer
	TimeRecordService *service.TimeRecordService
	AuthInterceptor   *AuthInterceptor
}

func (t *TimeRecordGrpcService) RegisterTimeRecord(ctx context.Context, in *pb.RegisterTimeRecordRequest) (*pb.TimeRecord, error) {
	timeRecord, err := t.TimeRecordService.Register(ctx, in.Time.AsTime(), in.Description, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.TimeRecord{}, err
	}

	return &pb.TimeRecord{
		Id:            timeRecord.ID,
		Time:          timestamppb.New(timeRecord.Time),
		Status:        pb.TimeRecord_Status(timeRecord.Status),
		Description:   timeRecord.Description,
		RefusedReason: timeRecord.RefusedReason,
		RegularTime:   timeRecord.RegularTime,
		EmployeeId:    timeRecord.EmployeeID,
		ApprovedBy:    timeRecord.ApprovedBy,
		RefusedBy:     timeRecord.RefusedBy,
		CreatedAt:     timestamppb.New(timeRecord.CreatedAt),
		UpdatedAt:     timestamppb.New(timeRecord.UpdatedAt),
	}, nil
}

func (t *TimeRecordGrpcService) ApproveTimeRecord(ctx context.Context, in *pb.ApproveTimeRecordRequest) (*pb.StatusResponse, error) {
	timeRecord, err := t.TimeRecordService.Approve(ctx, in.Id, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.StatusResponse{
			Status: "not updated",
			Error:  err.Error(),
		}, err
	}

	return &pb.StatusResponse{
		Status: "successfully " + timeRecord.Status.String(),
	}, nil
}

func (t *TimeRecordGrpcService) RefuseTimeRecord(ctx context.Context, in *pb.RefuseTimeRecordRequest) (*pb.StatusResponse, error) {
	timeRecord, err := t.TimeRecordService.Refuse(ctx, in.Id, in.RefusedReason, t.AuthInterceptor.EmployeeClaims.ID)
	if err != nil {
		return &pb.StatusResponse{
			Status: "not updated",
			Error:  err.Error(),
		}, err
	}

	return &pb.StatusResponse{
		Status: "successfully " + timeRecord.Status.String(),
	}, nil
}

func (t *TimeRecordGrpcService) FindTimeRecord(ctx context.Context, in *pb.FindTimeRecordRequest) (*pb.TimeRecord, error) {
	timeRecord, err := t.TimeRecordService.Find(ctx, in.Id)
	if err != nil {
		return &pb.TimeRecord{}, err
	}

	return &pb.TimeRecord{
		Id:            timeRecord.ID,
		Time:          timestamppb.New(timeRecord.Time),
		Status:        pb.TimeRecord_Status(timeRecord.Status),
		Description:   timeRecord.Description,
		RefusedReason: timeRecord.RefusedReason,
		RegularTime:   timeRecord.RegularTime,
		EmployeeId:    timeRecord.EmployeeID,
		ApprovedBy:    timeRecord.ApprovedBy,
		RefusedBy:     timeRecord.RefusedBy,
		CreatedAt:     timestamppb.New(timeRecord.CreatedAt),
		UpdatedAt:     timestamppb.New(timeRecord.UpdatedAt),
	}, nil
}

func (t *TimeRecordGrpcService) SearchTimeRecords(in *pb.SearchTimeRecordsRequest, stream pb.TimeRecordService_SearchTimeRecordsServer) error {
	timeRecords, err := t.TimeRecordService.FindAllByEmployeeID(stream.Context(), in.EmployeeId, in.FromDate.AsTime(), in.ToDate.AsTime())
	if err != nil {
		return err
	}

	for _, timeRecord := range timeRecords {
		stream.Send(&pb.TimeRecord{
			Id:            timeRecord.ID,
			Time:          timestamppb.New(timeRecord.Time),
			Status:        pb.TimeRecord_Status(timeRecord.Status),
			Description:   timeRecord.Description,
			RefusedReason: timeRecord.RefusedReason,
			RegularTime:   timeRecord.RegularTime,
			EmployeeId:    timeRecord.EmployeeID,
			ApprovedBy:    timeRecord.ApprovedBy,
			RefusedBy:     timeRecord.RefusedBy,
			CreatedAt:     timestamppb.New(timeRecord.CreatedAt),
			UpdatedAt:     timestamppb.New(timeRecord.UpdatedAt),
		})
	}

	return nil
}

func (t *TimeRecordGrpcService) ListTimeRecords(in *pb.ListTimeRecordsRequest, stream pb.TimeRecordService_ListTimeRecordsServer) error {
	timeRecords, err := t.TimeRecordService.FindAllByEmployeeID(stream.Context(), t.AuthInterceptor.EmployeeClaims.ID, in.FromDate.AsTime(), in.ToDate.AsTime())
	if err != nil {
		return err
	}

	for _, timeRecord := range timeRecords {
		stream.Send(&pb.TimeRecord{
			Id:            timeRecord.ID,
			Time:          timestamppb.New(timeRecord.Time),
			Status:        pb.TimeRecord_Status(timeRecord.Status),
			Description:   timeRecord.Description,
			RefusedReason: timeRecord.RefusedReason,
			RegularTime:   timeRecord.RegularTime,
			EmployeeId:    timeRecord.EmployeeID,
			ApprovedBy:    timeRecord.ApprovedBy,
			RefusedBy:     timeRecord.RefusedBy,
			CreatedAt:     timestamppb.New(timeRecord.CreatedAt),
			UpdatedAt:     timestamppb.New(timeRecord.UpdatedAt),
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
