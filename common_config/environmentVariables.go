package common_config

// ***********************************************************************************************************
// The following variables receives their values from environment variables

// ExecutionLocationForConnector - Where is the Worker running
var ExecutionLocationForConnector ExecutionLocationTypeType

// ExecutionLocationForFenixGuiExecutionServer  - Where is Fenix Execution Server running
var ExecutionLocationForFenixGuiExecutionServer ExecutionLocationTypeType

// ExecutionLocationTypeType - Definitions for where client and Fenix Server is running
type ExecutionLocationTypeType int

// Constants used for where stuff is running
const (
	LocalhostNoDocker ExecutionLocationTypeType = iota
	LocalhostDocker
	GCP
)

// Address to Fenix Execution Worker & Execution Connector, will have their values from Environment variables at startup
var (
	FenixGuiExecutionAddress       string
	FenixGuiExecutionPort          int
	FenixGuiExecutionAddressToDial string
	ExecutionsLoadGeneratorPort    int
	GCPAuthentication              bool
	CAEngineAddress                string
	CAEngineAddressPath            string
	UseInternalWebServerForTest    bool
	UseServiceAccount              bool
	ApplicationShouldRunInTray     bool
	TurnOffCallToWorker            bool
)
