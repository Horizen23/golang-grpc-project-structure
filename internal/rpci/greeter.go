package rpci

import (
	"context"
	pb "golang-grpc-project-structure/grpc/gen"
	"golang-grpc-project-structure/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type RpcIGreeterServe interface {
	SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error)
}

type rpciGreeterServer struct {
	pb.UnimplementedGreeterServer
	greeterService services.GreeterService
}

func (r *rpciGreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	m, err := r.greeterService.SayHello(in.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: "Hello " + m}, nil
}

func NewGreeterServer(greeterService services.GreeterService, gsrv *grpc.Server) RpcIGreeterServe {
	s := &rpciGreeterServer{
		greeterService: greeterService,
	}
	pb.RegisterGreeterServer(gsrv, s)
	healthServer := health.NewServer()
	// Register the health server
	grpc_health_v1.RegisterHealthServer(gsrv, healthServer)
     
	// Set the health status
	healthServer.SetServingStatus("HotelContentService", grpc_health_v1.HealthCheckResponse_SERVING)
	return s
}
