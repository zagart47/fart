package main

import (
	"context"
	"fart/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strings"
	"time"
)

func main() {
	var ADDRESS = ":11111"
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(conn)
	fart := proto.NewFartClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	cmd := strings.Join(os.Args[1:], " ")
	response, err := fart.DoCmd(ctx, &proto.CmdRequest{Cmd: cmd})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(response.GetText())
	if err != nil {
		fmt.Println(err.Error())
	}

}
