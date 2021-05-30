package repository

import (
	"context"
	"time"

	"github.com/c4ut/accounting-services/service/time-record/domain/model"
	"github.com/c4ut/accounting-services/service/time-record/infrastructure/db"
	"github.com/c4ut/accounting-services/service/time-record/infrastructure/db/collection"
	"go.elastic.co/apm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimeRecordRepository struct {
	M *db.Mongo
}

func (t *TimeRecordRepository) Register(ctx context.Context, timeRecord *model.TimeRecord) error {
	collection := t.M.Database.Collection(collection.TimeRecordCollection)
	_, err := collection.InsertOne(ctx, timeRecord)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return err
	}
	return nil
}

func (t *TimeRecordRepository) Save(ctx context.Context, timeRecord *model.TimeRecord) error {
	collection := t.M.Database.Collection(collection.TimeRecordCollection)
	_, err := collection.ReplaceOne(ctx, bson.M{"_id": timeRecord.ID}, timeRecord)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return err
	}
	return nil
}

func (t *TimeRecordRepository) Find(ctx context.Context, id string) (*model.TimeRecord, error) {
	var timeRecord *model.TimeRecord
	collection := t.M.Database.Collection(collection.TimeRecordCollection)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&timeRecord)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return timeRecord, err
	}
	return timeRecord, nil
}

func (t *TimeRecordRepository) FindAllByEmployeeID(ctx context.Context, employeeID string, fromDate, toDate time.Time) ([]*model.TimeRecord, error) {
	var timeRecords []*model.TimeRecord
	collection := t.M.Database.Collection(collection.TimeRecordCollection)

	findOpts := options.Find()
	findOpts.SetSort(bson.M{"time": -1})
	cur, err := collection.Find(
		ctx,
		bson.M{
			"employee_id": employeeID,
			"time": bson.M{
				"$gte": fromDate,
				"$lte": toDate,
			},
		},
		findOpts,
	)

	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	for cur.Next(ctx) {
		var timeRecord *model.TimeRecord
		err := cur.Decode(&timeRecord)
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			return nil, err
		}
		timeRecords = append(timeRecords, timeRecord)
	}

	if err := cur.Err(); err != nil {
		apm.CaptureError(ctx, err).Send()
		return nil, err
	}

	cur.Close(ctx)

	return timeRecords, nil
}

func NewTimeRecordRepository(database *db.Mongo) *TimeRecordRepository {
	return &TimeRecordRepository{
		M: database,
	}
}
