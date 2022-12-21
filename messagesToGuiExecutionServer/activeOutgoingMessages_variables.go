package messagesToGuiExecutionServer

import (
	"FenixExecutionsLoadGenerator/gcp"
	fenixGuiExecutionGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionServerGuiGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type MessagesToGuiExecutionObjectStruct struct {
	Logger *logrus.Logger
	//GcpAccessToken *oauth2.Token
	Gcp gcp.GcpObjectStruct
	//CommandChannelReference *connectorEngine.ExecutionEngineChannelType
}

// Variables used for contacting Fenix Execution Worker Server
var (
	remoteFenixGuiExecutionServerConnection *grpc.ClientConn
	FenixGuiExecutionAddressToDial          string
	fenixGuiExecutionGrpcClient             fenixGuiExecutionGrpcApi.FenixExecutionServerGuiGrpcServicesForGuiClientClient
)

var highestGuiExecutionsServerProtoFileVersion int32 = -1
