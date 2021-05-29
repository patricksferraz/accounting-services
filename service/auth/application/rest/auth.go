package rest

import (
	http "net/http"

	"github.com/c4ut/accounting-services/service/auth/domain/service"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

type AuthRestService struct {
	AuthService *service.AuthService
}

// Login godoc
// @Summary log in
// @ID login
// @Tags Auth
// @Description System authentication
// @Accept json
// @Produce json
// @Param body body Auth true "JSON body for authentication"
// @Success 200 {object} JWT
// @Failure 401 {object} HTTPError
// @Router /login [post]
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

// RefreshToken godoc
// @Summary refresh token
// @ID refreshToken
// @Tags Auth
// @Description Refresh token route
// @Accept json
// @Produce json
// @Param body body RefreshToken true "JSON body for refresh token"
// @Success 200 {object} JWT
// @Failure 400 {object} HTTPError
// @Router /refreshToken [post]
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}

// FindEmployeeClaimsByToken godoc
// @Summary get employee claims
// @ID findEmployeeClaimsByToken
// @Tags Auth
// @Description Get Employee Claims by access token
// @Accept json
// @Produce json
// @Param body body AccessToken true "JSON body for get employee claims"
// @Success 200 {object} EmployeeClaims
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /employeeClaims [post]
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func NewAuthRestService(service *service.AuthService) *AuthRestService {
	return &AuthRestService{
		AuthService: service,
	}
}
