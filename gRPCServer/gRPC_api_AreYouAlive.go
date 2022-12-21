package gRPCServer

import (
	"FenixExecutionsLoadGenerator/common_config"
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// AreYouAlive - *********************************************************************
// Anyone can check if Execution Connector server is alive with this service
func (s *fenixExecutionsLoadGeneratorGrpcServicesServer) AreYouAlive(
	ctx context.Context, emptyParameter *fenixExecutionsLoadGeneratorGrpcApi.EmptyParameter) (
	*fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse, error) {

	s.logger.WithFields(logrus.Fields{
		"id": "1ff67695-9a8b-4821-811d-0ab8d33c4d8b",
	}).Debug("Incoming 'gRPCServer - AreYouAlive'")

	s.logger.WithFields(logrus.Fields{
		"id": "9c7f0c3d-7e9f-4c91-934e-8d7a22926d84",
	}).Debug("Outgoing 'gRPCServer - AreYouAlive'")

	ackNackResponse := &fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse{
		AckNack:    true,
		Comments:   "I'am alive.",
		ErrorCodes: nil,
		ProtoFileVersionUsedByFenixExecutionsLoadGenerator: fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum(
			common_config.GetHighestExecutionsLoadGeneratorProtoFileVersion()),
	}

	return ackNackResponse, nil
}
