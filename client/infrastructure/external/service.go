package external

import (
	"google.golang.org/grpc"
)

func ConnectService(serviceAddr string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {

	conn, err := grpc.Dial(serviceAddr, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
