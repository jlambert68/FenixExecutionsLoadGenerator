package messagesToGuiExecutionServer

import (
	"FenixExecutionsLoadGenerator/common_config"
	"FenixExecutionsLoadGenerator/gcp"
	"context"
	fenixGuiExecutionGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// SendAreYouAliveToFenixGuiExecutionServer - Ask Execution Connector to check if Worker is up and running
func (toGuiExecutionObject *MessagesToGuiExecutionObjectStruct) ExecuteTestCase(
	testCaseExecutionRequest *fenixExecutionsLoadGeneratorGrpcApi.TestCaseExecutionRequest) {

	common_config.Logger.WithFields(logrus.Fields{
		"id": "bc373c0e-3ec8-49b7-926f-7ddec1f1e89e",
	}).Debug("Incoming 'ExecuteTestCase'")

	common_config.Logger.WithFields(logrus.Fields{
		"id": "66de077a-b258-46d5-a6b6-ed131d1c393b",
	}).Debug("Outgoing 'ExecuteTestCase'")

	var ctx context.Context
	var returnMessageAckNack bool
	var returnMessageString string

	ctx = context.Background()

	// Set up connection to Server
	ctx, err := toGuiExecutionObject.SetConnectionToFenixGuiExecutionServer(ctx)
	if err != nil {
		common_config.Logger.WithFields(logrus.Fields{
			"id":  "385c6a17-82a7-4616-bcaa-1aabcd4e29fe",
			"err": err,
		}).Error("Couldn't set up connection to GuiExecutionServer")

		return
	}

	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		toGuiExecutionObject.Logger.WithFields(logrus.Fields{
			"ID": "c5ba19bd-75ff-4366-818d-745d4d7f1a52",
		}).Debug("Running Defer Cancel function")
		cancel()
	}()

	// Only add access token when run on GCP
	if common_config.ExecutionLocationForGuiExecutionServer == common_config.GCP && common_config.GCPAuthentication == true {

		// Add Access token
		ctx, returnMessageAckNack, returnMessageString = gcp.Gcp.GenerateGCPAccessToken(ctx)
		if returnMessageAckNack == false {
			common_config.Logger.WithFields(logrus.Fields{
				"id":                  "0245556a-6fae-4471-be80-c844712bf26d",
				"err":                 err,
				"returnMessageString": returnMessageString,
			}).Error("Couldn't generate access token")

			return
		}

	}

	// Extract incoming parameters
	var (
		numberOfParallellGroups                 int32
		numberOfMilliSecondsBetweenItemsInGroup int32
		numberOfItemsPerGroup                   int32
	)

	numberOfParallellGroups = testCaseExecutionRequest.NumberOfParallellGroups
	numberOfMilliSecondsBetweenItemsInGroup = testCaseExecutionRequest.NumberOfSecondsBetweenItemsInGroup
	numberOfItemsPerGroup = testCaseExecutionRequest.NumberOfItemsPerGroup

	// When TestCaseUuid is empty then use predefined UUID
	var testCaseUuidToUse string
	if testCaseExecutionRequest.TestCaseUuid == "" {
		testCaseUuidToUse = "0e85ba4c-d55a-49f8-80aa-dc106035a759"
	} else {
		testCaseUuidToUse = testCaseExecutionRequest.TestCaseUuid
	}

	// Create Message to be Sent to GuiExecutionServer
	var initiateSingleTestCaseExecutionRequestMessage *fenixGuiExecutionGrpcApi.InitiateSingleTestCaseExecutionRequestMessage
	initiateSingleTestCaseExecutionRequestMessage = &fenixGuiExecutionGrpcApi.InitiateSingleTestCaseExecutionRequestMessage{
		UserAndApplicationRunTimeIdentification: &fenixGuiExecutionGrpcApi.UserAndApplicationRunTimeIdentificationMessage{
			ApplicationRunTimeUuid: "70f7c67c-f1b2-44b2-a583-97e109174c9e", // Just some UUID-value
			UserId:                 "LoadGenerator",
			ProtoFileVersionUsedByClient: fenixGuiExecutionGrpcApi.CurrentFenixExecutionGuiProtoFileVersionEnum(
				GetHighestGuiExecutionServerProtoFileVersion()),
		},
		TestCaseUuid:    testCaseUuidToUse,
		TestDataSetUuid: "c02ba879-3571-46d2-a99a-63a91b2235f9", // Just some UUID-value
	}

	// Create parallell groups of execution
	var groupCounter int32
	var waitGroup sync.WaitGroup
	for groupCounter = 0; groupCounter < numberOfParallellGroups; groupCounter++ {

		// Add for group to wait for in the end
		waitGroup.Add(1)

		// Start up each group
		go func(waitGroupReference *sync.WaitGroup, groupNumber int32) {

			// Variables need in gRPC-call
			var tempFenixGuiExecutionGrpcClient fenixGuiExecutionGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClientClient
			tempFenixGuiExecutionGrpcClient = fenixGuiExecutionGrpcApi.NewFenixExecutionServerGuiGrpcServicesForGuiClientClient(remoteFenixGuiExecutionServerConnection)

			// Loop all items for each group
			var itemCounter int32
			for itemCounter = 0; itemCounter < numberOfItemsPerGroup; itemCounter++ {

				// Do gRPC-call
				var initiateSingleTestCaseExecutionResponseMessage *fenixGuiExecutionGrpcApi.InitiateSingleTestCaseExecutionResponseMessage
				initiateSingleTestCaseExecutionResponseMessage, err = tempFenixGuiExecutionGrpcClient.InitiateTestCaseExecution(ctx, initiateSingleTestCaseExecutionRequestMessage)

				// Shouldn't happen
				if err != nil {
					common_config.Logger.WithFields(logrus.Fields{
						"ID":    "07c1dbcf-2fbb-4cc8-8178-59d4ee17e1b8",
						"error": err,
						"initiateSingleTestCaseExecutionRequestMessage": initiateSingleTestCaseExecutionRequestMessage,
					}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'InitiateTestCaseExecution'")

					return

				} else if initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.AckNack == false {
					// FenixTestDataSyncServer couldn't handle gPRC call
					common_config.Logger.WithFields(logrus.Fields{
						"ID": "3a02696a-5c94-4ef0-9846-4bd0d40daeaf",
						"initiateSingleTestCaseExecutionRequestMessage": initiateSingleTestCaseExecutionRequestMessage,
						"Message from Fenix Execution Server":           initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.Comments,
					}).Error("Problem to do gRPC-call to FenixGuiExecutionServer for 'InitiateTestCaseExecution'")

					return
				}

				common_config.Logger.WithFields(logrus.Fields{
					"ID":                    "775da478-c193-4b36-b399-f87b247cff7e",
					"TestCaseUuid":          testCaseUuidToUse,
					"TestCaseExecutionUuid": initiateSingleTestCaseExecutionResponseMessage.TestCasesInExecutionQueue.TestCaseExecutionUuid,
					"AckNack":               initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.AckNack,
					"Comments":              initiateSingleTestCaseExecutionResponseMessage.AckNackResponse.Comments,
				}).Info("Response frm GuiExecutionServer after initiating Execution for TestCase'")

				// Sleep before sending next TestCase for execution in this Group
				time.Sleep(time.Millisecond * time.Duration(numberOfMilliSecondsBetweenItemsInGroup))

			}

			// Group finished
			common_config.Logger.WithFields(logrus.Fields{
				"ID":           "b1943931-e2e3-42a1-a1be-17ed8af324fb",
				"TestCaseUuid": testCaseUuidToUse,
				"Group number": groupNumber,
			}).Info("Group is finished")

			waitGroupReference.Done()

		}(&waitGroup, groupCounter)
	}

	// Wait for all parallell groups to finish
	waitGroup.Wait()

	// All groups are finished
	common_config.Logger.WithFields(logrus.Fields{
		"ID": "0d54923e-7c3d-4adc-a6d9-debb25f22268",
	}).Info("All groups are finished")

	return

}
