package grpc

import (
	"context"
	"log"
	"time"

	"github.com/patricksferraz/accounting-services/client/domain/service"
	authmodel "github.com/patricksferraz/accounting-services/service/auth/domain/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authService *service.AuthService
	// authMethods map[string]bool
	jwt *authmodel.JWT
}

func (a *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Printf("--> unary interceptor: %s", method)

		return invoker(a.attachToken(ctx), method, req, reply, cc, opts...)
		// if a.authMethods[method] {
		// 	return invoker(a.attachToken(ctx), method, req, reply, cc, opts...)
		// }

		// return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (a *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		log.Printf("--> stream interceptor: %s", method)

		return streamer(a.attachToken(ctx), desc, cc, method, opts...)
		// if a.authMethods[method] {
		// 	return streamer(a.attachToken(ctx), desc, cc, method, opts...)
		// }

		// return streamer(ctx, desc, cc, method, opts...)
	}
}

func (a *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", a.jwt.AccessToken)
}

func (a *AuthInterceptor) login(username, password string) error {
	jwt, err := a.authService.Login(context.Background(), username, password)
	if err != nil {
		return err
	}
	a.jwt = jwt
	return nil
}

func (a *AuthInterceptor) scheduleRefreshToken() error {
	err := a.refreshToken()
	if err != nil {
		return err
	}

	go func() {
		wait := time.Duration(a.jwt.ExpiresIn)
		for {
			time.Sleep(wait)
			err := a.refreshToken()
			if err != nil {
				wait = time.Second
			} else {
				wait = time.Duration(a.jwt.ExpiresIn)
			}
		}
	}()

	return nil
}

func (a *AuthInterceptor) refreshToken() error {
	jwt, err := a.authService.RefreshToken(context.Background(), a.jwt.RefreshToken)
	if err != nil {
		return err
	}
	a.jwt = jwt
	return nil
}

func NewAuthInterceptor(authService *service.AuthService, username, password string) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		authService: authService,
		// authMethods: authMethods,
	}

	if err := interceptor.login(username, password); err != nil {
		return nil, err
	}
	if err := interceptor.scheduleRefreshToken(); err != nil {
		return nil, err
	}

	return interceptor, nil
}
