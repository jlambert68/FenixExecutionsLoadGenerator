package gRPCServer

import (
	"FenixExecutionsLoadGenerator/common_config"
	"FenixExecutionsLoadGenerator/messagesToGuiExecutionServer"
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// GuiExecutionServerExecuteTestCase - *********************************************************************
// Receive request for generate load towards ExecutionServer
func (s *fenixExecutionsLoadGeneratorGrpcServicesServer) GuiExecutionServerExecuteTestCase(
	ctx context.Context, testCaseExecutionRequest *fenixExecutionsLoadGeneratorGrpcApi.TestCaseExecutionRequest) (
	*fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse, error) {

	s.logger.WithFields(logrus.Fields{
		"id": "8d460e6a-bee4-44ab-b50d-0bf4630b4199",
	}).Debug("Incoming 'gRPCServer - GuiExecutionServerExecuteTestCase'")

	s.logger.WithFields(logrus.Fields{
		"id": "c5a7ceac-c186-4859-a252-a6e1e312b9c2",
	}).Debug("Outgoing 'gRPCServer - GuiExecutionServerExecuteTestCase'")

	// Current user
	userID := "gRPC-api doesn't support UserId"

	// Check if Client is using correct proto files version
	returnMessage := common_config.IsCallerUsingCorrectConnectorProtoFileVersion(
		userID, fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum(
			testCaseExecutionRequest.ProtoFileVersionUsedByCaller))
	if returnMessage != nil {

		// Exiting
		return returnMessage, nil
	}

	// Set up instance to use for execution gPRC
	var fenixGuiExecutionObject *messagesToGuiExecutionServer.MessagesToGuiExecutionObjectStruct
	fenixGuiExecutionObject = &messagesToGuiExecutionServer.MessagesToGuiExecutionObjectStruct{
		Logger: s.logger,
	}

	go fenixGuiExecutionObject.ExecuteTestCase(testCaseExecutionRequest)

	// Create Error Codes
	var errorCodes []fenixExecutionsLoadGeneratorGrpcApi.ErrorCodesEnum

	ackNackResponseMessage := &fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse{
		AckNack:    true,
		Comments:   "",
		ErrorCodes: errorCodes,
		ProtoFileVersionUsedByFenixExecutionsLoadGenerator: fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum(
			common_config.GetHighestExecutionsLoadGeneratorProtoFileVersion()),
	}

	return ackNackResponseMessage, nil

}
