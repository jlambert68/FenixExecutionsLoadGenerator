package common_config

import (
	fenixExecutionsLoadGeneratorGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionsLoadGenerator/fenixExecutionsLoadGeneratorGrpcApi/go_grpc_api"
)

// IsCallerUsingCorrectConnectorProtoFileVersion ********************************************************************************************************************
// Check if Caller  is using correct proto-file version
func IsCallerUsingCorrectConnectorProtoFileVersion(callingClientUuid string, usedProtoFileVersion fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum) (returnMessage *fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse) {

	var callerUseCorrectProtoFileVersion bool
	var protoFileExpected fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum
	var protoFileUsed fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum(GetHighestExecutionsLoadGeneratorProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		callerUseCorrectProtoFileVersion = true
	} else {
		callerUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if callerUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixExecutionsLoadGeneratorGrpcApi.ErrorCodesEnum
		var errorCode fenixExecutionsLoadGeneratorGrpcApi.ErrorCodesEnum

		errorCode = fenixExecutionsLoadGeneratorGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixExecutionsLoadGeneratorGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
			ProtoFileVersionUsedByFenixExecutionsLoadGenerator: protoFileExpected,
		}

		return returnMessage

	} else {
		return nil
	}

}

// GetHighestExecutionsLoadGeneratorProtoFileVersion
// Get the highest GetHighestExecutionsLoadGeneratorProtoFileVersion for Executions Load Generator
func GetHighestExecutionsLoadGeneratorProtoFileVersion() int32 {

	// Check if there already is a 'highestExecutionsLoadGeneratorProtoFileVersion' saved, if so use that one
	if highestExecutionsLoadGeneratorProtoFileVersion != -1 {
		return highestExecutionsLoadGeneratorProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixExecutionsLoadGeneratorGrpcApi.CurrentFenixExecutionsLoadGeneratorProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestExecutionsLoadGeneratorProtoFileVersion = maxValue

	return highestExecutionsLoadGeneratorProtoFileVersion
}
