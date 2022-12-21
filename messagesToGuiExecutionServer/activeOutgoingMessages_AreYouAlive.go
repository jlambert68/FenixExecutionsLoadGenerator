package messagesToGuiExecutionServer

import (
	"FenixExecutionsLoadGenerator/common_config"
	"FenixExecutionsLoadGenerator/gcp"
	"context"
	fenixGuiExecutionGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"time"
)

// SendAreYouAliveToFenixGuiExecutionServer - Ask Execution Connector to check if Worker is up and running
func (toGuiExecutionObject *MessagesToGuiExecutionObjectStruct) SendAreYouAliveToFenixGuiExecutionServer() (bool, string) {

	common_config.Logger.WithFields(logrus.Fields{
		"id": "5792072c-20a9-490b-a7cf-8c4f80979552",
	}).Debug("Incoming 'SendAreYouAliveToFenixGuiExecutionServer'")

	common_config.Logger.WithFields(logrus.Fields{
		"id": "353930b1-5c6f-4826-955c-19f543e2ab85",
	}).Debug("Outgoing 'SendAreYouAliveToFenixGuiExecutionServer'")

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string

	ctx = context.Background()

	// Set up connection to Server
	ctx, err := toGuiExecutionObject.SetConnectionToFenixGuiExecutionServer(ctx)
	if err != nil {
		return false, err.Error()
	}

	// Create the message with all test data to be sent to Fenix
	emptyParameter := &fenixGuiExecutionGrpcApi.EmptyParameter{

		ProtoFileVersionUsedByClient: fenixGuiExecutionGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(common_config.GetHighestExecutionsLoadGeneratorProtoFileVersion()),
	}

	// Do gRPC-call
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		toGuiExecutionObject.Logger.WithFields(logrus.Fields{
			"ID": "c5ba19bd-75ff-4366-818d-745d4d7f1a52",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForFenixGuiExecutionServer == common_config.GCP && common_config.GCPAuthentication == true {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			return false, returnMessageString
		}

	}

	// Do the gRPC-call
	//md2 := MetadataFromHeaders(headers)
	//myctx := metadata.NewOutgoingContext(ctx, md2)

	returnMessage, err := fenixGuiExecutionGrpcClient.AreYouAlive(ctx, emptyParameter)

	// Shouldn't happen
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"ID":    "818aaf0b-4112-4be4-97b9-21cc084c7b8b",
			"error": err,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendAreYouAliveToFenixGuiExecutionServer'")

		return false, err.Error()

	} else if returnMessage.AckNack == false {
		// FenixTestDataSyncServer couldn't handle gPRC call
		common_config.Logger.WithFields(logrus.Fields{
			"ID":                                  "2ecbc800-2fb6-4e88-858d-a421b61c5529",
			"Message from Fenix Execution Server": returnMessage.Comments,
		}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'SendAreYouAliveToFenixGuiExecutionServer'")

		return false, err.Error()
	}

	return returnMessage.AckNack, returnMessage.Comments

}
