package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/c4ut/accounting-services/service/time-record/domain/model"
	"github.com/c4ut/accounting-services/service/time-record/domain/service"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

type AuthMiddleware struct {
	AuthService    *service.AuthService
	EmployeeClaims *model.EmployeeClaims
}

func (a *AuthMiddleware) Require() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")
		if accessToken == "" {
			err := errors.New("authorization token is not provided")
			apm.CaptureError(ctx, err).Send()
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		claims, err := a.AuthService.Verify(ctx, accessToken)
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			ctx.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("access token is invalid: %v", err)})
			ctx.Abort()
			return
		}

		a.EmployeeClaims = claims

		// TODO: adds retricted permissions
		// for _, role := range claims.Roles {
		// 	if role == method {
		// 		return nil
		// 	}
		// }

		// return status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		AuthService: authService,
	}
}
