package userGRPC

import (
	"context"
	"fmt"

	authGrpc "github.com/j-ew-s/ms-curso-auth-grpc/auth"
	"google.golang.org/grpc"
)

func IsTokenValid(token string) (*authGrpc.TokenValidation, error) {
	var cc *grpc.ClientConn
	connection := ":5500"

	cc, err := grpc.Dial(connection, grpc.WithInsecure())

	if err != nil {
		fmt.Println(fmt.Sprintf("IsTokenValid : gRPC DIAL falhou: %+v - PORT %s", err, connection))
		return nil, err
	}

	defer cc.Close()
	u := authGrpc.NewUserServiceClient(cc)

	message := authGrpc.Token{
		Token: token,
	}

	response, err := u.IsTokenValid(context.Background(), &message)

	if err != nil {
		fmt.Println(fmt.Sprintf("Erro na chamaga do IsTokenValid : %+v  -- Token : %+v", err, &message))
		return nil, err
	}

	return response, nil
}
