package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sixi-store/users/app"
	_ "github.com/sixi-store/users/db"
	"github.com/sixi-store/users/pb"
	"google.golang.org/grpc"
)

func main() {
	var port int //端口
	flag.IntVar(&port, "port", 50000, "运行监听端口.")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, &app.UserServer{})
	// TODO: determine whether to use TLS
	grpcServer.Serve(lis)
}
