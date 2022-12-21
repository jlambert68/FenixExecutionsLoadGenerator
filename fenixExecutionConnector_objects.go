package main

import (
	"FenixExecutionsLoadGenerator/gRPCServer"
	"github.com/sirupsen/logrus"
)

type fenixExecutionsLoadGeneratorObjectStruct struct {
	logger     *logrus.Logger
	GrpcServer *gRPCServer.FenixExecutionsLoadGeneratorGrpcObjectStruct
}

// Variable holding everything together
var fenixExecutionsLoadGeneratorObject *fenixExecutionsLoadGeneratorObjectStruct
