package rest

import (
	"net/http"

	"github.com/c4ut/accounting-services/service/time-record/domain/service"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

type TimeRecordRestService struct {
	TimeRecordService *service.TimeRecordService
	AuthMiddleware    *AuthMiddleware
}

// RegisterTimeRecord godoc
// @Security ApiKeyAuth
// @Summary register a new time record
// @ID registerTimeRecord
// @Tags Time Record
// @Description Router for registration a new time record
// @Accept json
// @Produce json
// @Param body body TimeRecordRequest true "JSON body for register a new time record"
// @Success 200 {object} TimeRecord
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /timeRecord [post]
func (t *TimeRecordRestService) RegisterTimeRecord(ctx *gin.Context) {
	var req TimeRecordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	timeRecord, err := t.TimeRecordService.Register(ctx, req.Time, req.Description, t.AuthMiddleware.EmployeeClaims.ID)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, timeRecord)
}

// ApproveTimeRecord godoc
// @Security ApiKeyAuth
// @Summary approve a pending time record
// @ID approveTimeRecord
// @Tags Time Record
// @Description Router for approve a pending time record
// @Accept json
// @Produce json
// @Param id path string true "Time Record ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /timeRecord/{id}/approve [post]
func (t *TimeRecordRestService) ApproveTimeRecord(ctx *gin.Context) {
	var req IDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	timeRecord, err := t.TimeRecordService.Approve(ctx, req.ID, t.AuthMiddleware.EmployeeClaims.ID)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		HTTPResponse{
			Code:    http.StatusOK,
			Message: "successfully " + timeRecord.Status.String(),
		},
	)
}

// RefuseTimeRecord godoc
// @Security ApiKeyAuth
// @Summary refuse a pending time record
// @ID refuseTimeRecord
// @Tags Time Record
// @Description Router for refuse a pending time record
// @Accept json
// @Produce json
// @Param id path string true "Time Record ID"
// @Param body body RefuseRequest true "JSON body for refuse a pending time record"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /timeRecord/{id}/refuse [post]
func (t *TimeRecordRestService) RefuseTimeRecord(ctx *gin.Context) {
	var uri IDRequest
	var body RefuseRequest

	if err := ctx.ShouldBindUri(&uri); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	timeRecord, err := t.TimeRecordService.Refuse(ctx, uri.ID, body.RefusedReason, t.AuthMiddleware.EmployeeClaims.ID)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		HTTPResponse{
			Code:    http.StatusOK,
			Message: "successfully " + timeRecord.Status.String(),
		},
	)
}

// FindTimeRecord godoc
// @Security ApiKeyAuth
// @Summary find a time record
// @ID findTimeRecord
// @Tags Time Record
// @Description Router for find a time record
// @Accept json
// @Produce json
// @Param id path string true "Time Record ID"
// @Success 200 {object} TimeRecord
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /timeRecord/{id} [get]
func (t *TimeRecordRestService) FindTimeRecord(ctx *gin.Context) {
	var req IDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	timeRecord, err := t.TimeRecordService.Find(ctx, req.ID)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, timeRecord)
}

// SearchTimeRecords godoc
// @Security ApiKeyAuth
// @Summary search time records by employee id
// @ID searchTimeRecords
// @Tags Time Record
// @Description Search for employee time records by `id`
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param body query TimeRecordsRequest true "JSON body for search time records"
// @Success 200 {array} TimeRecord
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /timeRecord/employee/{id} [get]
func (t *TimeRecordRestService) SearchTimeRecords(ctx *gin.Context) {
	var uri IDRequest
	var body TimeRecordsRequest

	if err := ctx.ShouldBindUri(&uri); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	if err := ctx.ShouldBindQuery(&body); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	timeRecords, err := t.TimeRecordService.FindAllByEmployeeID(ctx, uri.ID, body.FromDate, body.ToDate)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, timeRecords)
}

// ListTimeRecords godoc
// @Security ApiKeyAuth
// @Summary list the employee's time records
// @ID listTimeRecords
// @Tags Time Record
// @Description List the employee's time records
// @Accept json
// @Produce json
// @Param body query TimeRecordsRequest true "JSON body for list the employee's time records"
// @Success 200 {array} TimeRecord
// @Failure 400 {object} HTTPError
// @Failure 403 {object} HTTPError
// @Router /timeRecord [get]
func (t *TimeRecordRestService) ListTimeRecords(ctx *gin.Context) {
	var body TimeRecordsRequest

	if err := ctx.ShouldBindQuery(&body); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusBadRequest,
			HTTPError{
				Code:  http.StatusBadRequest,
				Error: err.Error(),
			},
		)
		return
	}

	timeRecords, err := t.TimeRecordService.FindAllByEmployeeID(ctx, t.AuthMiddleware.EmployeeClaims.ID, body.FromDate, body.ToDate)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(
			http.StatusForbidden,
			HTTPError{
				Code:  http.StatusForbidden,
				Error: err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, timeRecords)
}

func NewTimeRecordRestService(service *service.TimeRecordService, authMiddleware *AuthMiddleware) *TimeRecordRestService {
	return &TimeRecordRestService{
		TimeRecordService: service,
		AuthMiddleware:    authMiddleware,
	}
}
