package model_test

import (
	"math"
	"testing"

	"github.com/patricksferraz/accounting-services/service/time-record/domain/model"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestModel_TimeRecordStatus(t *testing.T) {

	status := model.Pending
	require.Equal(t, status.String(), model.Pending.String())
	status = model.Approved
	require.Equal(t, status.String(), model.Approved.String())

	otherStatus := model.TimeRecordStatus(faker.RandomInt(int(model.Approved)+1, math.MaxInt64))
	require.Equal(t, otherStatus.String(), "")
}
