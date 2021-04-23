package utils_test

import (
	"testing"
	"time"

	"github.com/patricksferraz/accounting-services/service/time-record/utils"
	"github.com/stretchr/testify/require"
)

func TestUtils_DateEqual(t *testing.T) {

	today := time.Now()

	result := utils.DateEqual(today, today.Add(time.Second))
	require.True(t, result)
	result = utils.DateEqual(today, today.AddDate(0, 0, 1))
	require.False(t, result)
}
