package gRPCServer

import (
	"github.com/sirupsen/logrus"
)

// InitiategRPCObject - Initiate local logger object
func (fenixExecutionsLoadGeneratorGrpcObject *FenixExecutionsLoadGeneratorGrpcObjectStruct) InitiategRPCObject(logger *logrus.Logger) {

	fenixExecutionsLoadGeneratorGrpcObject.logger = logger
	//fenixExecutionsLoadGeneratorGrpcObject.CommandChannelReference = commandChannelReference

}

/*
// InitiateLocalObject - Initiate local 'ExecutionsLoadGeneratorGrpcObject'
func (fenixExecutionsLoadGeneratorGrpcObject *FenixExecutionsLoadGeneratorGrpcObjectStruct) InitiateLocalObject(inFenixExecutionsLoadGeneratorGrpcObject *FenixExecutionsLoadGeneratorGrpcObjectStruct) {

	fenixExecutionsLoadGeneratorGrpcObject.ExecutionsLoadGeneratorGrpcObject = inFenixExecutionsLoadGeneratorGrpcObject
}


*/
