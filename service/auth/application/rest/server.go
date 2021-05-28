package rest

import (
	"fmt"
	"log"

	_service "github.com/c4ut/accounting-services/service/auth/domain/service"
	"github.com/c4ut/accounting-services/service/auth/infrastructure/external"
	"github.com/c4ut/accounting-services/service/auth/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func StartRestServer(service *external.Keycloak, port int) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(apmgin.Middleware(r))

	authRepository := repository.NewAuthRepository(service)
	authService := _service.NewAuthService(authRepository)
	authRestService := NewAuthRestService(authService)

	v1 := r.Group("api/v1")
	{
		v1.POST("/login", authRestService.Login)
		v1.POST("/refreshToken", authRestService.RefreshToken)
		v1.POST("/employeeClaims", authRestService.FindEmployeeClaimsByToken)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Run(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
