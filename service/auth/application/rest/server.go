package rest

import (
	"fmt"
	"log"

	_ "github.com/c4ut/accounting-services/service/auth/application/rest/docs"
	_service "github.com/c4ut/accounting-services/service/auth/domain/service"
	"github.com/c4ut/accounting-services/service/auth/infrastructure/external"
	"github.com/c4ut/accounting-services/service/auth/infrastructure/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
)

// @title Auth Swagger API
// @version 1.0
// @description Swagger API for Golang Project Auth.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email comercial@coding4u.com.br

// @BasePath /api/v1
func StartRestServer(service *external.Keycloak, port int) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(apmgin.Middleware(r))

	authRepository := repository.NewAuthRepository(service)
	authService := _service.NewAuthService(authRepository)
	authRestService := NewAuthRestService(authService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("api/v1/auth")
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
