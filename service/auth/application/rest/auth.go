package rest

import (
	"net/http"

	"github.com/c4ut/accounting-services/service/auth/domain/service"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

type AuthRestService struct {
	AuthService *service.AuthService
}

func (a *AuthRestService) Login(ctx *gin.Context) {
	var json Auth
	if err := ctx.ShouldBindJSON(&json); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := a.AuthService.Login(ctx, json.Username, json.Password)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

func (a *AuthRestService) RefreshToken(ctx *gin.Context) {
	var json RefreshToken
	if err := ctx.ShouldBindJSON(&json); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := a.AuthService.RefreshToken(ctx, json.RefreshToken)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

func (a *AuthRestService) FindEmployeeClaimsByToken(ctx *gin.Context) {
	var json AccessToken
	if err := ctx.ShouldBindJSON(&json); err != nil {
		apm.CaptureError(ctx, err).Send()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := a.AuthService.FindEmployeeClaimsByToken(ctx, json.AccessToken)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func NewAuthRestService(service *service.AuthService) *AuthRestService {
	return &AuthRestService{
		AuthService: service,
	}
}
