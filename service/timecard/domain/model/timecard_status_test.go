package model_test

import (
	"math"
	"testing"

	"github.com/patricksferraz/timecard-service/domain/model"
	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestModel_TimecardStatus(t *testing.T) {

	status := model.Open
	require.Equal(t, status.String(), model.Open.String())
	status = model.WaitingConfirmation
	require.Equal(t, status.String(), model.WaitingConfirmation.String())
	status = model.Confirmed
	require.Equal(t, status.String(), model.Confirmed.String())

	otherStatus := model.TimecardStatus(faker.RandomInt(int(model.Confirmed)+1, math.MaxInt64))
	require.Equal(t, otherStatus.String(), "")
}
