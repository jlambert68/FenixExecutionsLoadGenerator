package common_config

// ***********************************************************************************************************
// The following variables receives their values from environment variables

// ExecutionLocationForExecutionsLoadGenerator - Where is the Worker running
var ExecutionLocationForExecutionsLoadGenerator ExecutionLocationTypeType

// ExecutionLocationForGuiExecutionServer  - Where is Fenix Execution Server running
var ExecutionLocationForGuiExecutionServer ExecutionLocationTypeType

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
	GuiExecutionServerAddress            string
	GuiExecutionServerPort               int
	FenixGuiExecutionServerAddressToDial string
	ExecutionsLoadGeneratorPort          int
	GCPAuthentication                    bool
	UseServiceAccount                    bool
	ApplicationShouldRunInTray           bool
)
