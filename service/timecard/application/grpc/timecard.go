package grpc

import (
	"context"
	"time"

	"github.com/patricksferraz/timecard-service/application/grpc/pb"
	"github.com/patricksferraz/timecard-service/domain/service"
)

type TimecardGrpcService struct {
	pb.UnimplementedTimecardServiceServer
	TimecardService service.TimecardService
}

func (t *TimecardGrpcService) RegisterTimecard(ctx context.Context, in *pb.TimecardRegister) (*pb.ResultInfo, error) {
	date, err := time.Parse(time.RFC3339, in.Date)
	if err != nil {
		return &pb.ResultInfo{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	timecard, err := t.TimecardService.RegisterTimecard(date, in.CompanyId, in.EmployeeId)
	if err != nil {
		return &pb.ResultInfo{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.ResultInfo{
		Id:     timecard.ID,
		Status: "created",
	}, nil
}

func (t *TimecardGrpcService) RegisterTimeRecord(ctx context.Context, in *pb.TimeRecordRegister) (*pb.ResultInfo, error) {
	_time, err := time.Parse(time.RFC3339, in.Time)
	if err != nil {
		return &pb.ResultInfo{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	timeRecord, err := t.TimecardService.RegisterTimeRecord(_time, in.TimecardId)
	if err != nil {
		return &pb.ResultInfo{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.ResultInfo{
		Id:     timeRecord.ID,
		Status: "created",
	}, nil
}

func (t *TimecardGrpcService) WaitConfirmationTimecard(ctc context.Context, in *pb.Timecard) (*pb.ResultInfo, error) {
	timecard, err := t.TimecardService.WaitConfirmationTimecard(in.Id)
	if err != nil {
		return &pb.ResultInfo{
			Status: "not updated",
			Error:  err.Error(),
		}, err
	}

	return &pb.ResultInfo{
		Id:     timecard.ID,
		Status: timecard.Status.String(),
	}, nil
}

func (t *TimecardGrpcService) ConfirmTimecard(ctc context.Context, in *pb.Timecard) (*pb.ResultInfo, error) {
	timecard, err := t.TimecardService.ConfirmTimecard(in.Id)
	if err != nil {
		return &pb.ResultInfo{
			Status: "not updated",
			Error:  err.Error(),
		}, err
	}

	return &pb.ResultInfo{
		Id:     timecard.ID,
		Status: timecard.Status.String(),
	}, nil
}

func (t *TimecardGrpcService) FindTimecard(ctx context.Context, in *pb.Timecard) (*pb.TimecardInfo, error) {
	timecard, err := t.TimecardService.FindTimecard(in.Id)
	if err != nil {
		return &pb.TimecardInfo{}, err
	}

	timeRecords, err := t.TimecardService.FindAllTimeRecordByTimecardID(timecard.ID)
	if err != nil {
		return &pb.TimecardInfo{}, err
	}

	var tr []*pb.TimeRecordInfo
	for _, timeRecord := range timeRecords {
		tr = append(tr, &pb.TimeRecordInfo{
			Id:         timeRecord.ID,
			Time:       timeRecord.Time.String(),
			TimecardId: timeRecord.TimecardID,
			CreatedAt:  timeRecord.CreatedAt.String(),
			UpdatedAt:  timeRecord.UpdatedAt.String(),
		})
	}

	return &pb.TimecardInfo{
		Id:          timecard.ID,
		Date:        timecard.Date.String(),
		CompanyId:   timecard.CompanyID,
		EmployeeId:  timecard.EmployeeID,
		TimeRecords: tr,
		CreatedAt:   timecard.CreatedAt.String(),
		UpdatedAt:   timecard.UpdatedAt.String(),
	}, nil
}

func (t *TimecardGrpcService) FindAllTimecardByCompanyID(in *pb.TimecardByCompanyID, stream pb.TimecardService_FindAllTimecardByCompanyIDServer) error {
	timecards, err := t.TimecardService.FindAllTimecardByCompanyID(in.CompanyId)
	if err != nil {
		return err
	}

	for _, timecard := range timecards {

		timeRecords, err := t.TimecardService.FindAllTimeRecordByTimecardID(timecard.ID)
		if err != nil {
			return err
		}

		var tr []*pb.TimeRecordInfo
		for _, timeRecord := range timeRecords {
			tr = append(tr, &pb.TimeRecordInfo{
				Id:         timeRecord.ID,
				Time:       timeRecord.Time.String(),
				TimecardId: timeRecord.TimecardID,
				CreatedAt:  timeRecord.CreatedAt.String(),
				UpdatedAt:  timeRecord.UpdatedAt.String(),
			})
		}

		stream.Send(&pb.TimecardInfo{
			Id:          timecard.ID,
			Date:        timecard.Date.String(),
			CompanyId:   timecard.CompanyID,
			EmployeeId:  timecard.EmployeeID,
			TimeRecords: tr,
			CreatedAt:   timecard.CreatedAt.String(),
			UpdatedAt:   timecard.UpdatedAt.String(),
		})
	}

	return nil
}

func (t *TimecardGrpcService) FindAllTimecardByEmployeeID(in *pb.TimecardByEmployeeID, stream pb.TimecardService_FindAllTimecardByEmployeeIDServer) error {
	timecards, err := t.TimecardService.FindAllTimecardByEmployeeID(in.EmployeeId)
	if err != nil {
		return err
	}

	for _, timecard := range timecards {

		timeRecords, err := t.TimecardService.FindAllTimeRecordByTimecardID(timecard.ID)
		if err != nil {
			return err
		}

		var tr []*pb.TimeRecordInfo
		for _, timeRecord := range timeRecords {
			tr = append(tr, &pb.TimeRecordInfo{
				Id:         timeRecord.ID,
				Time:       timeRecord.Time.String(),
				TimecardId: timeRecord.TimecardID,
				CreatedAt:  timeRecord.CreatedAt.String(),
				UpdatedAt:  timeRecord.UpdatedAt.String(),
			})
		}

		stream.Send(&pb.TimecardInfo{
			Id:          timecard.ID,
			Date:        timecard.Date.String(),
			CompanyId:   timecard.CompanyID,
			EmployeeId:  timecard.EmployeeID,
			TimeRecords: tr,
			CreatedAt:   timecard.CreatedAt.String(),
			UpdatedAt:   timecard.UpdatedAt.String(),
		})
	}

	return nil
}

func NewTimecardGrpcService(service service.TimecardService) *TimecardGrpcService {
	return &TimecardGrpcService{
		TimecardService: service,
	}
}
