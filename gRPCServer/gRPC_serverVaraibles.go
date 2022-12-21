package gRPCServer

import (
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type FenixExecutionsLoadGeneratorGrpcObjectStruct struct {
	logger *logrus.Logger
	//ExecutionsLoadGeneratorGrpcObject *FenixExecutionsLoadGeneratorGrpcObjectStruct
	//CommandChannelReference *connectorEngine.ExecutionEngineChannelType
}

// gRPCServer variables
var (
	fenixExecutionsLoadGeneratorGrpcServer *grpc.Server
	//registerFenixExecutionsLoadGeneratorGrpcServicesServer       *grpc.Server
	//registerFenixExecutionsLoadGeneratorWorkerGrpcServicesServer *grpc.Server
	lis net.Listener
)

// gRPCServer Server type
type fenixExecutionsLoadGeneratorGrpcServicesServer struct {
	logger *logrus.Logger
	fenixExecutionsLoadGeneratorGrpcApi.UnimplementedFenixExecutionsLoadGeneratorGrpcServicesServer
}
