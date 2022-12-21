package gRPCServer

import (
	"FenixExecutionsLoadGenerator/common_config"
	"FenixExecutionsLoadGenerator/messagesToGuiExecutionServer"
	"fmt"
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// AreYouAlive - *********************************************************************
// Anyone can check if Fenix Execution Worker server is alive with this service, should be used to check serves for Connector
func (s *fenixExecutionsLoadGeneratorGrpcServicesServer) GuiExecutionServerAreYouAlive(
	ctx context.Context, emptyParameter *fenixExecutionsLoadGeneratorGrpcApi.EmptyParameter) (
	*fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse, error) {

	s.logger.WithFields(logrus.Fields{
		"id": "dabd04a3-5357-4904-aceb-3493fa7396b6",
	}).Debug("Incoming 'gRPCServer - GuiExecutionServerAreYouAlive'")

	s.logger.WithFields(logrus.Fields{
		"id": "b9003ecf-b686-429b-b603-261f78e9c787",
	}).Debug("Outgoing 'gRPCServer - GuiExecutionServerAreYouAlive'")

	// Current user
	userID := "gRPC-api doesn't support UserId"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(userID, fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum(emptyParameter.ProtoFileVersionUsedByCaller))
	if returnMessage != nil {

		// Exiting
		return returnMessage, nil
	}

	// Set up instance to use for execution gPRC
	var fenixGuiExecutionObject *messagesToGuiExecutionServer.MessagesToGuiExecutionObjectStruct
	fenixGuiExecutionObject = &messagesToGuiExecutionServer.MessagesToGuiExecutionObjectStruct{
		Logger: s.logger,
		//GcpAccessToken: nil,
	}

	response, responseMessage := fenixGuiExecutionObject.SendAreYouAliveToFenixGuiExecutionServer()

	// Create Error Codes
	var errorCodes []fenixExecutionsLoadGeneratorGrpcApi.ErrorCodesEnum

	ackNackResponseMessage := &fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse{
		AckNack:    response,
		Comments:   fmt.Sprintf("The response from GuiExecutionServer is '%s'", responseMessage),
		ErrorCodes: errorCodes,
		ProtoFileVersionUsedByFenixExecutionsLoadGenerator: fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum(
			common_config.GetHighestExecutionsLoadGeneratorProtoFileVersion()),
	}

	return ackNackResponseMessage, nil

}
