package gRPCServer

import (
	"FenixExecutionsLoadGenerator/common_config"
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

// InitGrpcServer - Set up and start Backend gRPCServer-server
func (fenixExecutionsLoadGeneratorGrpcObject *FenixExecutionsLoadGeneratorGrpcObjectStruct) InitGrpcServer(logger *logrus.Logger) {

	var err error

	// Find first non allocated port from defined start port
	fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Info("Backend Server tries to start")

	fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
		"Id": "ca3593b1-466b-4536-be91-5e038de178f4",
		"common_config.ExecutionsLoadGeneratorPort: ": common_config.ExecutionsLoadGeneratorPort,
	}).Info("Start listening on:")
	lis, err = net.Listen("tcp", ":"+strconv.Itoa(common_config.ExecutionsLoadGeneratorPort))

	if err != nil {
		fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
			"Id":    "ad7815b3-63e8-4ab1-9d4a-987d9bd94c76",
			"err: ": err,
		}).Error("failed to listen:")
	} else {
		fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
			"Id": "ba070b9b-5d57-4c0a-ab4c-a76247a50fd3",
			"common_config.ExecutionsLoadGeneratorPort: ": common_config.ExecutionsLoadGeneratorPort,
		}).Info("Success in listening on port:")

	}

	// Create server and register the gRPC-service to the server
	fenixExecutionsLoadGeneratorGrpcServer = grpc.NewServer()
	fenixExecutionsLoadGeneratorGrpcApi.RegisterFenixExecutionsLoadGeneratorGrpcServicesServer(fenixExecutionsLoadGeneratorGrpcServer, &fenixExecutionsLoadGeneratorGrpcServicesServer{
		logger: logger})

	// Register Reflection on the same server to be able for calling agents to see the methods that are offered
	reflection.Register(fenixExecutionsLoadGeneratorGrpcServer)

	// Start server
	err = fenixExecutionsLoadGeneratorGrpcServer.Serve(lis)
	if err != nil {
		fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
			"Id":    "42abd1b8-2e01-4526-82b4-fb1d6af2b420",
			"err: ": err,
		}).Fatalln("Couldn't start gRPC server")
	}

}

// StopGrpcServer - Stop Backend gRPCServer-server
func (fenixExecutionsLoadGeneratorGrpcObject *FenixExecutionsLoadGeneratorGrpcObjectStruct) StopGrpcServer() {

	fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{}).Info("Gracefully stop for: fenixExecutionsLoadGeneratorGrpcServer")
	fenixExecutionsLoadGeneratorGrpcServer.GracefulStop()

	fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
		"common_config.ExecutionsLoadGeneratorPort: ": common_config.ExecutionsLoadGeneratorPort,
	}).Info("Close net.Listing")
	err := lis.Close()
	if err != nil {
		fenixExecutionsLoadGeneratorGrpcObject.logger.WithFields(logrus.Fields{
			"Id":    "6385920d-76c7-4139-8b4a-c5e629cf2301",
			"err: ": err,
			"common_config.ExecutionsLoadGeneratorPort": common_config.ExecutionsLoadGeneratorPort,
		}).Error("Couldn't stop listing on port")
	}

}
