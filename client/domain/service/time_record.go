package service

import (
	"context"
	"io"
	"time"

	"github.com/patricksferraz/accounting-services/client/domain/model"
	"github.com/patricksferraz/accounting-services/service/common/pb"
	trmodel "github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TimeRecordService struct {
	Service pb.TimeRecordServiceClient
}

func (t *TimeRecordService) Register(ctx context.Context, _time time.Time, description string) (*trmodel.TimeRecord, error) {
	req := &pb.RegisterTimeRecordRequest{
		Time:        timestamppb.New(_time),
		Description: description,
	}

	res, err := t.Service.RegisterTimeRecord(ctx, req)
	if err != nil {
		return nil, err
	}

	tr := &trmodel.TimeRecord{
		Time:        res.Time.AsTime(),
		Status:      trmodel.TimeRecordStatus(res.Status),
		Description: res.Description,
		RegularTime: res.RegularTime,
		EmployeeID:  res.EmployeeId,
		ApprovedBy:  res.ApprovedBy,
	}
	tr.ID = res.Id
	tr.CreatedAt = res.CreatedAt.AsTime()
	tr.UpdatedAt = res.UpdatedAt.AsTime()

	return tr, nil
}

func (t *TimeRecordService) Approve(ctx context.Context, id string) (*model.Response, error) {
	req := &pb.ApproveTimeRecordRequest{
		Id: id,
	}

	res, err := t.Service.ApproveTimeRecord(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.Response{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func (t *TimeRecordService) Refuse(ctx context.Context, id, refusedReason string) (*model.Response, error) {
	req := &pb.RefuseTimeRecordRequest{
		Id:            id,
		RefusedReason: refusedReason,
	}

	res, err := t.Service.RefuseTimeRecord(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.Response{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}

func (t *TimeRecordService) Find(ctx context.Context, id string) (*trmodel.TimeRecord, error) {
	req := &pb.FindTimeRecordRequest{
		Id: id,
	}

	res, err := t.Service.FindTimeRecord(ctx, req)
	if err != nil {
		return nil, err
	}

	tr := &trmodel.TimeRecord{
		Time:        res.Time.AsTime(),
		Status:      trmodel.TimeRecordStatus(res.Status),
		Description: res.Description,
		RegularTime: res.RegularTime,
		EmployeeID:  res.EmployeeId,
		ApprovedBy:  res.ApprovedBy,
	}
	tr.ID = res.Id
	tr.CreatedAt = res.CreatedAt.AsTime()
	tr.UpdatedAt = res.UpdatedAt.AsTime()

	return tr, nil
}

func (t *TimeRecordService) SearchTimeRecords(ctx context.Context, employeeID string, fromDate, toDate time.Time) ([]*trmodel.TimeRecord, error) {
	var timeRecords []*trmodel.TimeRecord

	req := &pb.SearchTimeRecordsRequest{
		EmployeeId: employeeID,
		FromDate:   timestamppb.New(fromDate),
		ToDate:     timestamppb.New(toDate),
	}

	stream, err := t.Service.SearchTimeRecords(ctx, req)
	if err != nil {
		return nil, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return timeRecords, nil
		}
		if err != nil {
			return nil, err
		}

		tr := &trmodel.TimeRecord{
			Time:        res.Time.AsTime(),
			Status:      trmodel.TimeRecordStatus(res.Status),
			Description: res.Description,
			RegularTime: res.RegularTime,
			EmployeeID:  res.EmployeeId,
			ApprovedBy:  res.ApprovedBy,
		}
		tr.ID = res.Id
		tr.CreatedAt = res.CreatedAt.AsTime()
		tr.UpdatedAt = res.UpdatedAt.AsTime()
		timeRecords = append(timeRecords, tr)
	}
}

func (t *TimeRecordService) ListTimeRecords(ctx context.Context, fromDate, toDate time.Time) ([]*trmodel.TimeRecord, error) {
	var timeRecords []*trmodel.TimeRecord

	req := &pb.ListTimeRecordsRequest{
		FromDate: timestamppb.New(fromDate),
		ToDate:   timestamppb.New(toDate),
	}

	stream, err := t.Service.ListTimeRecords(ctx, req)
	if err != nil {
		return nil, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return timeRecords, nil
		}
		if err != nil {
			return nil, err
		}

		tr := &trmodel.TimeRecord{
			Time:        res.Time.AsTime(),
			Status:      trmodel.TimeRecordStatus(res.Status),
			Description: res.Description,
			RegularTime: res.RegularTime,
			EmployeeID:  res.EmployeeId,
			ApprovedBy:  res.ApprovedBy,
		}
		tr.ID = res.Id
		tr.CreatedAt = res.CreatedAt.AsTime()
		tr.UpdatedAt = res.UpdatedAt.AsTime()
		timeRecords = append(timeRecords, tr)
	}
}

func NewTimeRecordService(service pb.TimeRecordServiceClient) *TimeRecordService {
	return &TimeRecordService{
		Service: service,
	}
}
