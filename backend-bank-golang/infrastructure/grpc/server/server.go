package server

import (
	"log"
	"net"

	"github.com/cleitonestefenon/imersao-fullcycle/infrastructure/grpc/pb"
	"github.com/cleitonestefenon/imersao-fullcycle/infrastructure/grpc/service"
	"github.com/cleitonestefenon/imersao-fullcycle/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("could not listen tpc port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	grpcServer.Serve(lis)
}
