package grpc_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/patricksferraz/accounting-services/service/common/application/grpc/pb"
	"github.com/patricksferraz/accounting-services/service/time-record/application/grpc"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"github.com/patricksferraz/accounting-services/service/time-record/domain/service"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	gogrpc "google.golang.org/grpc"
	"syreclabs.com/go/faker"
)

type mockTimeRecordService_FindAllByEmployeeIDServer struct {
	gogrpc.ServerStream
	Results []*pb.TimeRecord
}

func (_m *mockTimeRecordService_FindAllByEmployeeIDServer) Send(timeRecord *pb.TimeRecord) error {
	_m.Results = append(_m.Results, timeRecord)
	return nil
}

type repository struct{}

func (r *repository) Register(timeRecord *model.TimeRecord) error {
	// NOTE: Force error
	if timeRecord.Description == "error" {
		return errors.New("")
	}
	return nil
}

func (r *repository) Save(timeRecord *model.TimeRecord) error {
	// NOTE: Force error
	if timeRecord.ID == "c03c4cd4-5211-4209-ac68-17e441152b1d" {
		return errors.New("")
	}
	return nil
}

func (r *repository) Find(id string) (*model.TimeRecord, error) {
	timeRecord := model.TimeRecord{
		Time:        time.Now().AddDate(0, 0, -1),
		Status:      model.Pending,
		Description: faker.Lorem().Sentence(10),
		RegularTime: false,
		EmployeeID:  "67fe1eea-25a4-4f23-bf67-64f9a085311d",
	}
	timeRecord.ID = id
	// NOTE: Force error
	if id == "c4a80742-5294-4f1e-8ea9-5126c9389d6f" {
		return nil, errors.New("")
	}
	return &timeRecord, nil
}

func (r *repository) FindAllByEmployeeID(employeeID string) ([]*model.TimeRecord, error) {
	var timeRecords []*model.TimeRecord
	timeRecords = append(
		timeRecords,
		&model.TimeRecord{
			Time:        time.Now().AddDate(0, 0, -1),
			Status:      model.Pending,
			Description: faker.Lorem().Sentence(10),
			RegularTime: false,
			EmployeeID:  uuid.NewV4().String(),
		},
	)
	// NOTE: Force error
	if employeeID == "" {
		return nil, errors.New("")
	}
	return timeRecords, nil
}

func TestGrpc_Register(t *testing.T) {

	interceptor := &grpc.AuthInterceptor{
		EmployeeClaims: &model.EmployeeClaims{
			ID: uuid.NewV4().String(),
		},
	}
	timeRecordService := service.NewTimeRecordService(new(repository))
	timeRecordGrpcService := grpc.NewTimeRecordGrpcService(timeRecordService, interceptor)

	ctx := new(context.Context)
	registerRequest := &pb.RegisterRequest{
		Time:        time.Now().Format(time.RFC3339),
		Description: faker.Lorem().Sentence(10),
	}

	_, err := timeRecordGrpcService.Register(*ctx, registerRequest)
	require.Nil(t, err)

	registerRequest.Time = time.Now().String()
	_, err = timeRecordGrpcService.Register(*ctx, registerRequest)
	require.NotNil(t, err)

	registerRequest.Time = time.Now().Format(time.RFC3339)
	registerRequest.Description = "error"
	_, err = timeRecordGrpcService.Register(*ctx, registerRequest)
	require.NotNil(t, err)
}

func TestGrpc_Approve(t *testing.T) {

	interceptor := &grpc.AuthInterceptor{
		EmployeeClaims: &model.EmployeeClaims{
			ID: uuid.NewV4().String(),
		},
	}
	timeRecordService := service.NewTimeRecordService(new(repository))
	timeRecordGrpcService := grpc.NewTimeRecordGrpcService(timeRecordService, interceptor)

	ctx := new(context.Context)
	approveRequest := &pb.ApproveRequest{
		Id: uuid.NewV4().String(),
	}

	_, err := timeRecordGrpcService.Approve(*ctx, approveRequest)
	require.Nil(t, err)

	approveRequest.Id = "c4a80742-5294-4f1e-8ea9-5126c9389d6f"
	_, err = timeRecordGrpcService.Approve(*ctx, approveRequest)
	require.NotNil(t, err)
}

func TestGrpc_Find(t *testing.T) {

	interceptor := &grpc.AuthInterceptor{}
	timeRecordService := service.NewTimeRecordService(new(repository))
	timeRecordGrpcService := grpc.NewTimeRecordGrpcService(timeRecordService, interceptor)

	ctx := new(context.Context)
	findRequest := &pb.FindRequest{
		Id: uuid.NewV4().String(),
	}

	_, err := timeRecordGrpcService.Find(*ctx, findRequest)
	require.Nil(t, err)

	findRequest.Id = "c4a80742-5294-4f1e-8ea9-5126c9389d6f"
	_, err = timeRecordGrpcService.Find(*ctx, findRequest)
	require.NotNil(t, err)
}

func TestGrpc_FindAllByEmployeeID(t *testing.T) {

	interceptor := &grpc.AuthInterceptor{}
	timeRecordService := service.NewTimeRecordService(new(repository))
	timeRecordGrpcService := grpc.NewTimeRecordGrpcService(timeRecordService, interceptor)

	findAllByEmployeeIDRequest := &pb.FindAllByEmployeeIDRequest{
		EmployeeId: uuid.NewV4().String(),
	}

	mock := &mockTimeRecordService_FindAllByEmployeeIDServer{}
	err := timeRecordGrpcService.FindAllByEmployeeID(findAllByEmployeeIDRequest, mock)
	require.Equal(t, 1, len(mock.Results))
	require.Nil(t, err)

	findAllByEmployeeIDRequest.EmployeeId = ""
	err = timeRecordGrpcService.FindAllByEmployeeID(findAllByEmployeeIDRequest, mock)
	require.NotNil(t, err)
}
