package head

import (
	"context"
	"fart/proto"
	"fmt"
	"google.golang.org/grpc"
	"os/exec"
)

type serverAPI struct {
	proto.UnimplementedFartServer
}

func Register(gRPC *grpc.Server) {
	proto.RegisterFartServer(gRPC, &serverAPI{})
}

func (s serverAPI) DoCmd(ctx context.Context, req *proto.CmdRequest) (*proto.CmdResponse, error) {
	command := exec.Command("bash", "-c", req.GetCmd())

	output, err := command.Output()
	if err != nil {
		return nil, fmt.Errorf("непонятная ошибка %w", err)
	}
	response := &proto.CmdResponse{
		Text: string(output),
	}
	return response, nil
}
