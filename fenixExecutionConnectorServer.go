package main

import (
	"FenixExecutionsLoadGenerator/common_config"
	"FenixExecutionsLoadGenerator/gRPCServer"
	"github.com/sirupsen/logrus"
)

// Used for only process cleanup once
var cleanupProcessed = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		fenixExecutionsLoadGeneratorObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend GrpcServer Server
		fenixExecutionsLoadGeneratorObject.GrpcServer.StopGrpcServer()

	}
}

func fenixExecutionsLoadGeneratorMain() {

	// Set up BackendObject
	fenixExecutionsLoadGeneratorObject = &fenixExecutionsLoadGeneratorObjectStruct{
		logger:     nil,
		GrpcServer: &gRPCServer.FenixExecutionsLoadGeneratorGrpcObjectStruct{},
	}

	// Init logger
	//fenixExecutionsLoadGeneratorObject.InitLogger(loggerFileName)
	fenixExecutionsLoadGeneratorObject.logger = common_config.Logger

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Initiate  gRPC-server
	fenixExecutionsLoadGeneratorObject.GrpcServer.InitiategRPCObject(fenixExecutionsLoadGeneratorObject.logger)

	// Start Backend GrpcServer-server
	fenixExecutionsLoadGeneratorObject.GrpcServer.InitGrpcServer(fenixExecutionsLoadGeneratorObject.logger)

}
