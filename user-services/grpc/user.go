package userGRPC

import (
	"context"
	"fmt"

	"github.com/j-ew-s/ms-curso-user-api/user"
	"google.golang.org/grpc"
)

func GetByUserId(userId string) (*user.User, error) {
	var cc *grpc.ClientConn

	cc, err := grpc.Dial(":5100", grpc.WithInsecure())

	if err != nil {
		fmt.Println(fmt.Sprintf("GetByUserId gRPC DIAL falhou: %v", err))
		return nil, err
	}

	defer cc.Close()
	u := user.NewUserServiceClient(cc)

	message := user.UserId{
		Id: userId,
	}

	response, err := u.GetUser(context.Background(), &message)

	if err != nil {
		fmt.Println(fmt.Sprintf("Erro na chamaga do GetUser : %v", err))
		return nil, err
	}

	return response, nil
}

func IsTokenValid(token string) (*user.TokenValidation, error) {
	var cc *grpc.ClientConn

	cc, err := grpc.Dial(":5001", grpc.WithInsecure())

	if err != nil {
		fmt.Println(fmt.Sprintf("IsTokenValid : gRPC DIAL falhou: %v", err))
		return nil, err
	}

	defer cc.Close()
	u := user.NewUserServiceClient(cc)

	message := user.Token{
		Token: token,
	}

	response, err := u.IsTokenValid(context.Background(), &message)

	if err != nil {
		fmt.Println(fmt.Sprintf("Erro na chamaga do IsTokenValid : %v", err))
		return nil, err
	}

	return response, nil
}
